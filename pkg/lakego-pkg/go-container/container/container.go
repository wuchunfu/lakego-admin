package container

import (
    "fmt"
    "errors"
    "reflect"
)

// Container Bind data
type ContainerBind struct {
    // bind func
    Concrete any

    // shared
    Shared bool
}

/**
 * Container
 *
 * @create 2024-8-26
 * @author deatil
 */
type Container struct {
    // bind
    bind           map[string]ContainerBind

    // instances
    instances      map[string]any

    // invokeCallback
    invokeCallback map[string][]any
}

// NewContainer creates and returns a new *Container.
func NewContainer() *Container {
    container := &Container{
        bind:           make(map[string]ContainerBind),
        instances:      make(map[string]any),
        invokeCallback: make(map[string][]any),
    }

    container.Singleton(container, func() *Container {
        return container
    })

    return container
}

// New creates and returns a new *Container.
func New() *Container {
    return NewContainer()
}

// Provide
func (this *Container) Provide(fn any) error {
    fnType := reflect.TypeOf(fn)

    numOut := fnType.NumOut()
    if numOut < 1 {
        return errors.New("go-container: Provide set fail")
    }

    outs := this.Call(fn, nil)

    var errPtr *error
    errType := reflect.TypeOf(errPtr).Elem()

    for i := 0; i < numOut; i++ {
        outValue := reflect.ValueOf(outs[i])
        if ifImplements(outValue.Type(), errType) && !outValue.IsNil() {
            return outs[i].(error)
        }

        if isStruct(outs[i]) {
            abstract := getTypeName(fnType.Out(i))

            this.Singleton(abstract, outs[i])
        }
    }

    return nil
}

// Invoke for get dep struct
func (this *Container) Invoke(fn any) []any {
    return this.Call(fn, nil)
}

// Get object from bind
func (this *Container) Get(abstracts any) any {
    abstract := formatName(abstracts)

    if this.Has(abstract) {
        return this.Make(abstract, nil)
    }

    panic(fmt.Sprintf("go-container: bind not exists: %s", abstract))
}

// Bind object
func (this *Container) Bind(abstract any, concrete any) *Container {
    return this.BindWithShared(abstract, concrete, false)
}

// Singleton object
func (this *Container) Singleton(abstract any, concrete any) *Container {
    return this.BindWithShared(abstract, concrete, true)
}

// Bind object With Shared
func (this *Container) BindWithShared(abstracts any, concrete any, shared bool) *Container {
    if isStruct(concrete) {
        abstract := formatName(abstracts)

        this.Instance(abstract, concrete)
    } else {
        abstract := this.GetAlias(abstracts)

        this.bind[abstract] = ContainerBind{
            Concrete: concrete,
            Shared:   shared,
        }
    }

    return this
}

// Instance
func (this *Container) Instance(abstracts any, instance any) *Container {
    abstract := this.GetAlias(abstracts)

    this.bind[abstract] = ContainerBind{
        Shared: true,
    }

    this.instances[abstract] = instance

    return this
}

// Resolving
func (this *Container) Resolving(abstracts string, callback any) {
    abstract := formatName(abstracts)

    if abstract == "*" {
        if _, ok := this.invokeCallback["*"]; !ok {
            this.invokeCallback["*"] = make([]any, 0)
        }

        this.invokeCallback["*"] = append(this.invokeCallback["*"], callback)
    }

    abstract = this.GetAlias(abstract)

    if _, ok := this.invokeCallback[abstract]; !ok {
        this.invokeCallback[abstract] = make([]any, 0)
    }

    this.invokeCallback[abstract] = append(this.invokeCallback[abstract], callback)
}

// Get Alias
func (this *Container) GetAlias(abstracts any) string {
    abstract := formatName(abstracts)

    if _, ok := this.bind[abstract]; ok {
        bind := this.bind[abstract]

        if concrete, ok := bind.Concrete.(string); ok {
            return this.GetAlias(concrete)
        }
    }

    return abstract
}

// Make
func (this *Container) Make(abstracts any, vars []any) any {
    abstract := this.GetAlias(abstracts)

    bind := this.bind[abstract]

    if instance, ok := this.instances[abstract]; ok && bind.Shared {
        return instance
    }

    objects := this.Call(bind.Concrete, vars)

    var object any
    if len(objects) > 0 {
        object = objects[0]
    }

    if bind.Shared {
        this.instances[abstract] = object
    }

    this.invokeAfter(abstract, object)

    return object
}

// Bound
func (this *Container) Bound(abstracts any) bool {
    abstract := formatName(abstracts)

    if _, ok := this.bind[abstract]; ok {
        return true
    }

    if _, ok := this.instances[abstract]; ok {
        return true
    }

    return false
}

// Has
func (this *Container) Has(abstracts any) bool {
    return this.Bound(abstracts)
}

// Exists
func (this *Container) Exists(abstracts any) bool {
    abstract := this.GetAlias(abstracts)

    if _, ok := this.instances[abstract]; ok {
        return true
    }

    return false
}

// IsShared
func (this *Container) IsShared(abstracts any) bool {
    abstract := formatName(abstracts)

    if _, ok := this.instances[abstract]; ok {
        return true
    }

    if bind, ok := this.bind[abstract]; ok && bind.Shared {
        return true
    }

    return false
}

// Delete
func (this *Container) Delete(abstracts any) {
    abstract := this.GetAlias(abstracts)

    if _, ok := this.instances[abstract]; ok {
        delete(this.instances, abstract)
    }
}

// Delete
func (this *Container) invokeAfter(abstract string, object any) {
    if _, ok := this.invokeCallback["*"]; ok {
        for _, fn := range this.invokeCallback["*"] {
            this.Call(fn, []any{object, this})
        }
    }

    if _, ok := this.invokeCallback[abstract]; ok {
        for _, fn := range this.invokeCallback[abstract] {
            this.Call(fn, []any{object, this})
        }
    }
}

// Call
func (this *Container) Call(fn any, args []any) []any {
    switch in := fn.(type) {
        case []any:
            if len(in) > 1 {
                method := ""
                if methodName, ok := in[1].(string); ok {
                    method = methodName
                }

                return this.call(in[0], method, args)
            } else if len(in) == 1 {
                return this.call(in[0], "", args)
            }

            panic("go-container: call slice func error")
        default:
            return this.call(in, "", args)
    }
}

func (this *Container) call(in any, method string, params []any) []any {
    if eventMethod, ok := in.(reflect.Value); ok {
        if eventMethod.Kind() == reflect.Func {
            return this.CallFunc(eventMethod, params)
        }

        return this.CallStructMethod(eventMethod, method, params)
    } else if isFunc(in) {
        return this.CallFunc(in, params)
    } else {
        return this.CallStructMethod(in, method, params)
    }
}

// Call Func
func (this *Container) CallFunc(fn any, args []any) []any {
    var val reflect.Value
    if fnVal, ok := fn.(reflect.Value); ok {
        val = fnVal
    } else {
        val = reflect.ValueOf(fn)
    }

    if val.Kind() != reflect.Func {
        panic("go-container: func type error")
    }

    return this.baseCall(val, args)
}

// Call struct method
func (this *Container) CallStructMethod(class any, method string, args []any) []any {
    var val reflect.Value
    if fnVal, ok := class.(reflect.Value); ok {
        val = fnVal
    } else {
        val = reflect.ValueOf(class)
    }

    if val.Kind() != reflect.Pointer && val.Kind() != reflect.Struct {
        panic("go-container: struct type error")
    }

    newMethod := val.MethodByName(method)
    return this.baseCall(newMethod, args)
}

// base Call Func
func (this *Container) baseCall(fn reflect.Value, args []any) (output []any) {
    if fn.Kind() != reflect.Func {
        panic("go-container: call func type error")
    }

    if !fn.IsValid() {
        panic("go-container: call func valid error")
    }

    fnType := fn.Type()

    // 参数
    params := this.bindParams(fnType, args)

    res := fn.Call(params)

    // 返回数据
    output = make([]any, 0)
    for _, v := range res {
        output = append(output, v.Interface())
    }

    return
}

// bind params
func (this *Container) bindParams(fnType reflect.Type, args []any) []reflect.Value {
    numIn := fnType.NumIn()

    newArgs := make([]any, 0)
    if len(args) == numIn {
        newArgs = args
    } else {
        j := 0
        for i := 0; i < numIn; i++ {
            fnTypeIn := fnType.In(i)
            fnTypeKind := fnTypeIn.Kind()
            fnTypeName := getTypeName(fnTypeIn)

            if fnTypeKind == reflect.Pointer || fnTypeKind == reflect.Struct {
                if j < len(args) {
                    argsType := getStructName(args[j])

                    // 传入参数和函数参数类型一致时
                    if fnTypeName == argsType {
                        newArgs = append(newArgs, args[j])

                        j++
                    } else {
                        if this.Has(fnTypeName) {
                            newArgs = append(newArgs, this.Make(fnTypeName, nil))
                        }
                    }
                } else {
                    if this.Has(fnTypeName) {
                        newArgs = append(newArgs, this.Make(fnTypeName, nil))
                    }
                }

            } else if fnTypeKind == reflect.Interface {
                // when has set interface as key
                if this.Has(fnTypeName) {
                    newArgs = append(newArgs, this.Make(fnTypeName, nil))
                } else {
                    if name, ok := this.getImplementsBind(fnTypeIn); ok {
                        newArgs = append(newArgs, this.Make(name, nil))
                    }
                }
            } else {
                newArgs = append(newArgs, args[j])

                j++
            }
        }
    }

    if len(newArgs) != numIn {
        err := fmt.Sprintf("go-container: func params error (args %d, func args %d)", len(args), numIn)
        panic(err)
    }

    // 参数
    params := make([]reflect.Value, 0)
    for i := 0; i < numIn; i++ {
        dataValue := this.convertTo(fnType.In(i), newArgs[i])
        params = append(params, dataValue)
    }

    return params
}

// src convert type to new typ
func (this *Container) convertTo(typ reflect.Type, src any) reflect.Value {
    dataKey := getTypeName(typ)

    fieldType := reflect.TypeOf(src)
    if !fieldType.ConvertibleTo(typ) {
        return reflect.New(typ).Elem()
    }

    fieldValue := reflect.ValueOf(src)

    if dataKey != getTypeName(fieldType) {
        fieldValue = fieldValue.Convert(typ)
    }

    return fieldValue
}

func (this *Container) getImplementsBind(typ reflect.Type) (string, bool) {
    for name2, instance := range this.instances {
        value := reflect.TypeOf(instance)
        if value == nil {
            continue
        }

        valueName := getTypeName(value)

        if ifImplements(value, typ) && valueName == name2 {
            return name2, true
        }
    }

    for name, bind := range this.bind {
        concreteType := reflect.TypeOf(bind.Concrete)
        if concreteType == nil {
            continue
        }

        if concreteType.NumOut() > 0 {
            value := concreteType.Out(0)
            valueName := getTypeName(value)

            if ifImplements(value, typ) && valueName == name {
                return name, true
            }
        }
    }

    return "", false
}

