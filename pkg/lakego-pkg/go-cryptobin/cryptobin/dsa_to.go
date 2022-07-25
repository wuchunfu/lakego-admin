package cryptobin

// 私钥/公钥
func (this DSA) ToKeyBytes() []byte {
    return this.keyData
}

// 私钥/公钥
func (this DSA) ToKeyString() string {
    return string(this.keyData)
}

// ==========

// 输出字节
func (this DSA) ToBytes() []byte {
    return this.paredData
}

// 输出字符
func (this DSA) ToString() string {
    return string(this.paredData)
}

// 输出Base64
func (this DSA) ToBase64String() string {
    return NewEncoding().Base64Encode(this.paredData)
}

// 输出Hex
func (this DSA) ToHexString() string {
    return NewEncoding().HexEncode(this.paredData)
}

// ==========

// 验证结果
func (this DSA) ToVeryed() bool {
    return this.veryed
}
