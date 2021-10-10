package controller

import (
    "github.com/gin-gonic/gin"

    "github.com/deatil/lakego-admin/lakego/facade/auth"
    "github.com/deatil/lakego-admin/lakego/facade/config"
    "github.com/deatil/lakego-admin/lakego/facade/captcha"
    "github.com/deatil/lakego-admin/lakego/facade/cache"
    "github.com/deatil/lakego-admin/lakego/support/hash"

    "github.com/deatil/lakego-admin/admin/model"
    "github.com/deatil/lakego-admin/admin/support/jwt"
    "github.com/deatil/lakego-admin/admin/support/http/code"
    authPassword "github.com/deatil/lakego-admin/lakego/auth/password"
    passportValidate "github.com/deatil/lakego-admin/admin/validate/passport"
)

type Passport struct {
    Base
}

/**
 * 验证码
 */
func (control *Passport) Captcha(ctx *gin.Context) {
    c := captcha.New()
    id, b64s, err := c.Generate()
    if err != nil {
        control.Error(ctx, "error", code.StatusError)
    }

    key := config.New("auth").GetString("Passport.HeaderCaptchaKey")

    control.SetHeader(ctx, key, id)
    control.SuccessWithData(ctx, "获取成功", gin.H{
        "captcha": b64s,
    })
}

/**
 * 登陆
 */
func (control *Passport) Login(ctx *gin.Context) {
    // 接收数据
    post := make(map[string]interface{})
    ctx.BindJSON(&post)

    validateErr := passportValidate.Login(post)
    if validateErr != "" {
        control.Error(ctx, validateErr, code.LoginError)
        return
    }

    name := post["name"].(string)
    password := post["password"].(string)
    captchaCode := post["captcha"].(string)

    // 验证码检测
    key := config.New("auth").GetString("Passport.HeaderCaptchaKey")
    captchaId := ctx.GetHeader(key)

    ok := captcha.New().Verify(captchaId, captchaCode, false)
    if !ok {
        control.Error(ctx, "验证码错误", code.LoginError)
        return
    }

    // 用户信息
    admin := map[string]interface{}{}
    err := model.NewAdmin().
        Where(&model.Admin{Name: name}).
        First(&admin).
        Error
    if err != nil {
        control.Error(ctx, "账号或者密码错误", code.LoginError)
        return
    }

    // 验证密码
    checkStatus := authPassword.CheckPassword(admin["password"].(string), password, admin["password_salt"].(string))
    if !checkStatus {
        control.Error(ctx, "账号或者密码错误", code.LoginError)
        return
    }

    // 生成 token
    aud := jwt.GetJwtAud(ctx)
    jwter := auth.NewWithAud(aud)

    // token 数据
    tokenData := map[string]string{
        "id": admin["id"].(string),
    }

    // 授权 token
    accessToken, err := jwter.MakeAccessToken(tokenData)
    if err != nil {
        control.Error(ctx, "授权token生成失败", code.LoginError)
        return
    }

    // 刷新 token
    refreshToken, err := jwter.MakeRefreshToken(tokenData)
    if err != nil {
        control.Error(ctx, "刷新token生成失败", code.LoginError)
        return
    }

    // 授权 token 过期时间
    expiresIn := jwter.GetAccessExpiresIn()

    // 数据输出
    control.SuccessWithData(ctx, "获取成功", gin.H{
        "access_token": accessToken,
        "expires_in": expiresIn,
        "refresh_token": refreshToken,
    })
}

/**
 * 刷新 token
 */
func (control *Passport) RefreshToken(ctx *gin.Context) {
    // 接收数据
    post := make(map[string]interface{})
    ctx.BindJSON(&post)

    var refreshToken interface{}
    var ok bool

    if refreshToken, ok = post["refresh_token"]; !ok {
        control.Error(ctx, "refreshToken不能为空", code.JwtRefreshTokenFail)
        return
    }

    c := cache.New()
    refreshTokenPutTime, _ := c.Get(hash.MD5(refreshToken.(string)))
    refreshTokenPutTime = refreshTokenPutTime.(string)
    if refreshTokenPutTime != "" {
        control.Error(ctx, "refreshToken已失效", code.JwtRefreshTokenFail)
        return
    }

    // jwt
    aud := jwt.GetJwtAud(ctx)
    jwter := auth.NewWithAud(aud)

    // 拿取数据
    adminId := jwter.GetRefreshTokenData(refreshToken.(string), "id")
    if adminId == "" {
        control.Error(ctx, "刷新Token失败", code.JwtRefreshTokenFail)
        return
    }

    // token 数据
    tokenData := map[string]string{
        "id": adminId,
    }

    // 授权 token
    accessToken, err := jwter.MakeAccessToken(tokenData)
    if err != nil {
        control.Error(ctx, "生成 access_token 失败", code.JwtRefreshTokenFail)
        return
    }

    // 授权 token 过期时间
    expiresIn := jwter.GetAccessExpiresIn()

    // 数据输出
    control.SuccessWithData(ctx, "获取成功", gin.H{
        "access_token": accessToken,
        "expires_in": expiresIn,
    })
}

/**
 * 退出
 */
func (control *Passport) Logout(ctx *gin.Context) {
    // 接收数据
    post := make(map[string]interface{})
    ctx.BindJSON(&post)

    var refreshToken interface{}
    var ok bool

    if refreshToken, ok = post["refresh_token"]; !ok {
        control.Error(ctx, "refreshToken 不能为空", code.JwtRefreshTokenFail)
        return
    }

    c := cache.New()
    refreshTokenPutString, _ := c.Get(hash.MD5(refreshToken.(string)))
    refreshTokenPutString = refreshTokenPutString.(string)
    if refreshTokenPutString != "" {
        control.Error(ctx, "refreshToken 已失效", code.JwtRefreshTokenFail)
        return
    }

    // jwt
    aud := jwt.GetJwtAud(ctx)
    jwter := auth.NewWithAud(aud)

    // 拿取数据
    claims, claimsErr := jwter.GetRefreshTokenClaims(refreshToken.(string))
    if claimsErr != nil {
        control.Error(ctx, "refreshToken 已失效", code.JwtRefreshTokenFail)
        return
    }

    // 当前账号ID
    adminId := jwter.GetDataFromTokenClaims(claims, "id")

    // 过期时间
    exp := jwter.GetFromTokenClaims(claims, "exp")
    iat := jwter.GetFromTokenClaims(claims, "iat")
    refreshTokenExpiresIn := exp.(float64) - iat.(float64)

    nowAdminId, _ := ctx.Get("admin_id")
    if adminId != nowAdminId.(string) {
        control.Error(ctx, "退出失败", code.JwtRefreshTokenFail)
        return
    }

    // 当前 accessToken
    accessToken, _ := ctx.Get("access_token")

    // 加入黑名单
    c.Put(hash.MD5(accessToken.(string)), "no", int64(refreshTokenExpiresIn))
    c.Put(hash.MD5(refreshToken.(string)), "no", int64(refreshTokenExpiresIn))

    // 数据输出
    control.Success(ctx, "退出成功")
}