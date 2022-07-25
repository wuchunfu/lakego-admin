package cryptobin

import (
    "crypto/dsa"
)

// 设置 PrivateKey
func (this DSA) WithPrivateKey(data *dsa.PrivateKey) DSA {
    this.privateKey = data

    return this
}

// 设置 PublicKey
func (this DSA) WithPublicKey(data *dsa.PublicKey) DSA {
    this.publicKey = data

    return this
}

// 设置 data
func (this DSA) WithData(data []byte) DSA {
    this.data = data

    return this
}

// 设置 paredData
func (this DSA) WithParedData(data []byte) DSA {
    this.paredData = data

    return this
}

// 设置 hash 类型
// 可用参数可查看 Hash 结构体数据
func (this DSA) WithSignHash(data string) DSA {
    this.signHash = data

    return this
}

// 设置 veryed
func (this DSA) WithVeryed(data bool) DSA {
    this.veryed = data

    return this
}

// 设置错误
func (this DSA) WithError(err error) DSA {
    this.Error = err

    return this
}
