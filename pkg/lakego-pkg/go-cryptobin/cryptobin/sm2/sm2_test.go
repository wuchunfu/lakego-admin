package sm2

import (
    "testing"
    "crypto/rand"
    "encoding/base64"

    cryptobin_test "github.com/deatil/go-cryptobin/tool/test"
)

// signedData = S4vhrJoHXn98ByNw73CSOCqguYeuc4LrhsIHqkv/xA8Waw7YOLsfQzOKzxAjF0vyPKKSEQpq4zEgj9Mb/VL1pQ==
func Test_SM2_SignBytes(t *testing.T) {
    assertError := cryptobin_test.AssertErrorT(t)
    assertNotEmpty := cryptobin_test.AssertNotEmptyT(t)

    uid := "N002462434000000"

    sm2key := "NBtl7WnuUtA2v5FaebEkU0/Jj1IodLGT6lQqwkzmd2E="
    sm2keyBytes, err2 := base64.StdEncoding.DecodeString(sm2key)

    assertError(err2, "sm2keyDecode")

    data := "123123"

    signedData := NewSM2().
        FromString(data).
        FromPrivateKeyBytes(sm2keyBytes).
        SignBytes([]byte(uid)).
        ToBase64String()

    assertNotEmpty(signedData, "sm2-SignBytes")
}

func Test_SM2_VerifyBytes(t *testing.T) {
    assertError := cryptobin_test.AssertErrorT(t)
    assertBool := cryptobin_test.AssertBoolT(t)

    uid := "N002462434000000"

    sm2key := "NBtl7WnuUtA2v5FaebEkU0/Jj1IodLGT6lQqwkzmd2E="
    sm2keyBytes, err2 := base64.StdEncoding.DecodeString(sm2key)

    assertError(err2, "sm2keyDecode")

    data := "123123"
    signedData := "S4vhrJoHXn98ByNw73CSOCqguYeuc4LrhsIHqkv/xA8Waw7YOLsfQzOKzxAjF0vyPKKSEQpq4zEgj9Mb/VL1pQ=="

    verify := NewSM2().
        FromBase64String(signedData).
        FromPrivateKeyBytes(sm2keyBytes).
        MakePublicKey().
        VerifyBytes([]byte(data), []byte(uid))

    assertError(verify.Error(), "sm2VerifyError")
    assertBool(verify.ToVerify(), "sm2-VerifyBytes")
}

func Test_SM2_Encrypt_C1C2C3(t *testing.T) {
    assertEqual := cryptobin_test.AssertEqualT(t)
    assertError := cryptobin_test.AssertErrorT(t)
    assertNotEmpty := cryptobin_test.AssertNotEmptyT(t)

    sm2key := "NBtl7WnuUtA2v5FaebEkU0/Jj1IodLGT6lQqwkzmd2E="
    sm2keyBytes, err2 := base64.StdEncoding.DecodeString(sm2key)

    assertError(err2, "Encrypt_C1C2C3-sm2keyDecode")

    data := "test-pass"

    sm2 := NewSM2()

    en := sm2.
        FromString(data).
        FromPrivateKeyBytes(sm2keyBytes).
        MakePublicKey().
        SetMode("C1C2C3"). // C1C3C2 | C1C2C3
        Encrypt()
    enData := en.ToBase64String()

    assertError(en.Error(), "Encrypt_C1C2C3-Encrypt")
    assertNotEmpty(enData, "Encrypt_C1C2C3-Encrypt")

    de := sm2.
        FromBase64String(enData).
        FromPrivateKeyBytes(sm2keyBytes).
        SetMode("C1C2C3"). // C1C3C2 | C1C2C3
        Decrypt()
    deData := de.ToString()

    assertError(de.Error(), "Encrypt_C1C2C3-Decrypt")
    assertNotEmpty(deData, "Encrypt_C1C2C3-Decrypt")

    assertEqual(data, deData, "Encrypt_C1C2C3-Dedata")
}

func Test_SM2_Encrypt_C1C3C2(t *testing.T) {
    assertEqual := cryptobin_test.AssertEqualT(t)
    assertError := cryptobin_test.AssertErrorT(t)
    assertNotEmpty := cryptobin_test.AssertNotEmptyT(t)

    sm2key := "NBtl7WnuUtA2v5FaebEkU0/Jj1IodLGT6lQqwkzmd2E="
    sm2keyBytes, err2 := base64.StdEncoding.DecodeString(sm2key)

    assertError(err2, "Encrypt_C1C3C2-sm2keyDecode")

    data := "test-pass"

    sm2 := NewSM2()

    en := sm2.
        FromString(data).
        FromPrivateKeyBytes(sm2keyBytes).
        MakePublicKey().
        SetMode("C1C3C2"). // C1C3C2 | C1C2C3
        Encrypt()
    enData := en.ToBase64String()

    assertError(en.Error(), "Encrypt_C1C3C2-Encrypt")
    assertNotEmpty(enData, "Encrypt_C1C3C2-Encrypt")

    de := sm2.
        FromBase64String(enData).
        FromPrivateKeyBytes(sm2keyBytes).
        SetMode("C1C3C2"). // C1C3C2 | C1C2C3
        Decrypt()
    deData := de.ToString()

    assertError(de.Error(), "Encrypt_C1C3C2-Decrypt")
    assertNotEmpty(deData, "Encrypt_C1C3C2-Decrypt")

    assertEqual(data, deData, "Encrypt_C1C3C2-Dedata")
}

func Test_SM2_EncryptASN1_C1C2C3(t *testing.T) {
    assertEqual := cryptobin_test.AssertEqualT(t)
    assertError := cryptobin_test.AssertErrorT(t)
    assertNotEmpty := cryptobin_test.AssertNotEmptyT(t)

    sm2key := "NBtl7WnuUtA2v5FaebEkU0/Jj1IodLGT6lQqwkzmd2E="
    sm2keyBytes, err2 := base64.StdEncoding.DecodeString(sm2key)

    assertError(err2, "DecryptASN1_C1C2C3-sm2keyDecode")

    data := "test-pass"

    sm2 := NewSM2()

    en := sm2.
        FromString(data).
        FromPrivateKeyBytes(sm2keyBytes).
        MakePublicKey().
        SetMode("C1C2C3"). // C1C3C2 | C1C2C3
        EncryptASN1()
    enData := en.ToBase64String()

    assertError(en.Error(), "DecryptASN1_C1C2C3-Encrypt")
    assertNotEmpty(enData, "DecryptASN1_C1C2C3-Encrypt")

    de := sm2.
        FromBase64String(enData).
        FromPrivateKeyBytes(sm2keyBytes).
        SetMode("C1C2C3"). // C1C3C2 | C1C2C3
        DecryptASN1()
    deData := de.ToString()

    assertError(de.Error(), "DecryptASN1_C1C2C3-Decrypt")
    assertNotEmpty(deData, "DecryptASN1_C1C2C3-Decrypt")

    assertEqual(data, deData, "DecryptASN1_C1C2C3-Dedata")
}

func Test_SM2_EncryptASN1_C1C3C2(t *testing.T) {
    assertEqual := cryptobin_test.AssertEqualT(t)
    assertError := cryptobin_test.AssertErrorT(t)
    assertNotEmpty := cryptobin_test.AssertNotEmptyT(t)

    sm2key := "NBtl7WnuUtA2v5FaebEkU0/Jj1IodLGT6lQqwkzmd2E="
    sm2keyBytes, err2 := base64.StdEncoding.DecodeString(sm2key)

    assertError(err2, "DecryptASN1_C1C3C2-sm2keyDecode")

    data := "test-pass"

    sm2 := NewSM2()

    en := sm2.
        FromString(data).
        FromPrivateKeyBytes(sm2keyBytes).
        MakePublicKey().
        SetMode("C1C3C2"). // C1C3C2 | C1C2C3
        EncryptASN1()
    enData := en.ToBase64String()

    assertError(en.Error(), "DecryptASN1_C1C3C2-Encrypt")
    assertNotEmpty(enData, "DecryptASN1_C1C3C2-Encrypt")

    de := sm2.
        FromBase64String(enData).
        FromPrivateKeyBytes(sm2keyBytes).
        SetMode("C1C3C2"). // C1C3C2 | C1C2C3
        DecryptASN1()
    deData := de.ToString()

    assertError(de.Error(), "DecryptASN1_C1C3C2-Decrypt")
    assertNotEmpty(deData, "DecryptASN1_C1C3C2-Decrypt")

    assertEqual(data, deData, "DecryptASN1_C1C3C2-Dedata")
}

var (
    prikeyPKCS1 = `
-----BEGIN SM2 PRIVATE KEY-----
MHcCAQEEIAVunzkO+VYC1MFl3TfjjEHkc21eRBz+qRxbgEA6BP/FoAoGCCqBHM9V
AYItoUQDQgAEAnfcXztAc2zQ+uHuRlXuMohDdsncWxQFjrpxv5Ae3/PgH9vewt4A
oEvRqcwOBWtAXNDP6E74e5ocagfMUbq4hQ==
-----END SM2 PRIVATE KEY-----

    `
    pubkeyPKCS1 = `
-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEAnfcXztAc2zQ+uHuRlXuMohDdsnc
WxQFjrpxv5Ae3/PgH9vewt4AoEvRqcwOBWtAXNDP6E74e5ocagfMUbq4hQ==
-----END PUBLIC KEY-----

    `
)

func Test_PKCS1Sign(t *testing.T) {
    assertBool := cryptobin_test.AssertBoolT(t)
    assertError := cryptobin_test.AssertErrorT(t)
    assertNotEmpty := cryptobin_test.AssertNotEmptyT(t)

    data := "test-pass"

    // 签名
    objSign := New().
        FromString(data).
        FromPKCS1PrivateKey([]byte(prikeyPKCS1)).
        Sign()
    signed := objSign.ToBase64String()

    assertError(objSign.Error(), "PKCS1Sign-Sign")
    assertNotEmpty(signed, "PKCS1Sign-Sign")

    // 验证
    objVerify := New().
        FromBase64String(signed).
        FromPublicKey([]byte(pubkeyPKCS1)).
        Verify([]byte(data))

    assertError(objVerify.Error(), "PKCS1Sign-Verify")
    assertBool(objVerify.ToVerify(), "PKCS1Sign-Verify")
}

func Test_PKCS1Encrypt(t *testing.T) {
    assertEqual := cryptobin_test.AssertEqualT(t)
    assertError := cryptobin_test.AssertErrorT(t)
    assertNotEmpty := cryptobin_test.AssertNotEmptyT(t)

    data := "123tesfd!df"

    objEn := New().
        FromString(data).
        FromPublicKey([]byte(pubkeyPKCS1)).
        Encrypt()
    enData := objEn.ToBase64String()

    assertError(objEn.Error(), "PKCS1Encrypt-Encrypt")
    assertNotEmpty(enData, "PKCS1Encrypt-Encrypt")

    objDe := New().
        FromBase64String(enData).
        FromPKCS1PrivateKey([]byte(prikeyPKCS1)).
        Decrypt()
    deData := objDe.ToString()

    assertError(objDe.Error(), "PKCS1Encrypt-Decrypt")
    assertNotEmpty(deData, "PKCS1Encrypt-Decrypt")

    assertEqual(data, deData, "PKCS1Encrypt-Dedata")
}

var (
    testSM2PublicKeyX  = "a4b75c4c8c44d11687bdd93c0883e630c895234beb685910efbe27009ad911fa"
    testSM2PublicKeyY  = "d521f5e8249de7a405f254a9888cbb8e651fd60c50bd22bd182a4bc7d1261c94"
    testSM2PrivateKeyD = "0f495b5445eb59ddecf0626f5ca0041c550584f0189e89d95f8d4c52499ff838"
)

func Test_CreateKey(t *testing.T) {
    assertError := cryptobin_test.AssertErrorT(t)
    assertNotEmpty := cryptobin_test.AssertNotEmptyT(t)

    pub := New().
        FromPublicKeyXYString(testSM2PublicKeyX, testSM2PublicKeyY).
        CreatePublicKey()
    pubData := pub.ToKeyString()

    assertError(pub.Error(), "CreateKey-pub")
    assertNotEmpty(pubData, "CreateKey-pub")

    // ======

    pri := New().
        FromPrivateKeyString(testSM2PrivateKeyD).
        CreatePrivateKey()
    priData := pri.ToKeyString()

    assertError(pri.Error(), "CreateKey-pri")
    assertNotEmpty(priData, "CreateKey-pri")
}

var testPEMCiphers = []string{
    "DESCBC",
    "DESEDE3CBC",
    "AES128CBC",
    "AES192CBC",
    "AES256CBC",

    "DESCFB",
    "DESEDE3CFB",
    "AES128CFB",
    "AES192CFB",
    "AES256CFB",

    "DESOFB",
    "DESEDE3OFB",
    "AES128OFB",
    "AES192OFB",
    "AES256OFB",

    "DESCTR",
    "DESEDE3CTR",
    "AES128CTR",
    "AES192CTR",
    "AES256CTR",
}

func Test_CreatePKCS1PrivateKeyWithPassword(t *testing.T) {
    for _, cipher := range testPEMCiphers{
        test_CreatePKCS1PrivateKeyWithPassword(t, cipher)
    }
}

func test_CreatePKCS1PrivateKeyWithPassword(t *testing.T, cipher string) {
    assertEqual := cryptobin_test.AssertEqualT(t)
    assertError := cryptobin_test.AssertErrorT(t)
    assertNotEmpty := cryptobin_test.AssertNotEmptyT(t)

    t.Run(cipher, func(t *testing.T) {
        pass := make([]byte, 12)
        _, err := rand.Read(pass)
        if err != nil {
            t.Fatal(err)
        }

        gen := New().GenerateKey()

        prikey := gen.GetPrivateKey()

        pri := gen.
            CreatePKCS1PrivateKeyWithPassword(string(pass), cipher).
            ToKeyString()

        assertError(gen.Error(), "Test_CreatePKCS1PrivateKeyWithPassword")
        assertNotEmpty(pri, "Test_CreatePKCS1PrivateKeyWithPassword-pri")

        newPrikey := New().
            FromPKCS1PrivateKeyWithPassword([]byte(pri), string(pass)).
            GetPrivateKey()

        assertNotEmpty(newPrikey, "Test_CreatePKCS1PrivateKeyWithPassword-newPrikey")

        assertEqual(newPrikey, prikey, "Test_CreatePKCS1PrivateKeyWithPassword")
    })
}
