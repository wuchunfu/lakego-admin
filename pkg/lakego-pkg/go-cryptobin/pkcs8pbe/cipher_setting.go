package pkcs8pbe

import (
    "crypto/des"
    "crypto/md5"
    "crypto/sha1"
    "crypto/cipher"
    "encoding/asn1"

    cryptobin_md2 "github.com/deatil/go-cryptobin/hash/md2"
    cryptobin_rc2 "github.com/deatil/go-cryptobin/cipher/rc2"
    cryptobin_des "github.com/deatil/go-cryptobin/cipher/des"
)

var (
    // pcks12 模式
    oidPbeWithSHA1And3DES    = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 12, 1, 3}
    oidPbeWithSHA1And2DES    = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 12, 1, 4}
    oidPbeWithSHA1AndRC2_128 = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 12, 1, 5}
    oidPbeWithSHA1AndRC2_40  = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 12, 1, 6}
    oidPbeWithSHA1AndRC4_128 = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 12, 1, 1}
    oidPbeWithSHA1AndRC4_40  = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 12, 1, 2}

    // pkcs5-v1.5 模式
    oidPbeWithMD2AndDES      = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 5, 1}
    oidPbeWithMD2AndRC2_64   = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 5, 4}
    oidPbeWithMD5AndDES      = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 5, 3}
    oidPbeWithMD5AndRC2_64   = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 5, 6}
    oidPbeWithSHA1AndDES     = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 5, 10}
    oidPbeWithSHA1AndRC2_64  = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 5, 11}
)

var (
    newRC2Cipher = func(key []byte) (cipher.Block, error) {
        return cryptobin_rc2.NewCipher(key, len(key)*8)
    }
)

// pkcs12 模式
var PEMCipherSHA1And3DES = CipherBlockCBC{
    cipherFunc:     des.NewTripleDESCipher,
    hashFunc:       sha1.New,
    derivedKeyFunc: derivedKeyWithPbkdf,
    saltSize:       des.BlockSize,
    keySize:        24,
    blockSize:      des.BlockSize,
    iterationCount: 2048,
    oid:            oidPbeWithSHA1And3DES,
}
var PEMCipherSHA1And2DES = CipherBlockCBC{
    cipherFunc:     cryptobin_des.NewTwoDESCipher,
    hashFunc:       sha1.New,
    derivedKeyFunc: derivedKeyWithPbkdf,
    saltSize:       cryptobin_des.BlockSize,
    keySize:        16,
    blockSize:      cryptobin_des.BlockSize,
    iterationCount: 2048,
    oid:            oidPbeWithSHA1And2DES,
}
var PEMCipherSHA1AndRC2_128 = CipherBlockCBC{
    cipherFunc:     newRC2Cipher,
    hashFunc:       sha1.New,
    derivedKeyFunc: derivedKeyWithPbkdf,
    saltSize:       20,
    keySize:        16,
    blockSize:      cryptobin_rc2.BlockSize,
    iterationCount: 2048,
    oid:            oidPbeWithSHA1AndRC2_128,
}
var PEMCipherSHA1AndRC2_40 = CipherBlockCBC{
    cipherFunc:     newRC2Cipher,
    hashFunc:       sha1.New,
    derivedKeyFunc: derivedKeyWithPbkdf,
    saltSize:       20,
    keySize:        5,
    blockSize:      cryptobin_rc2.BlockSize,
    iterationCount: 2048,
    oid:            oidPbeWithSHA1AndRC2_40,
}
var PEMCipherSHA1AndRC4_128 = CipherRC4{
    hashFunc:       sha1.New,
    derivedKeyFunc: derivedKeyWithPbkdf,
    saltSize:       20,
    keySize:        16,
    blockSize:      16,
    iterationCount: 2048,
    oid:            oidPbeWithSHA1AndRC4_128,
}
var PEMCipherSHA1AndRC4_40 = CipherRC4{
    hashFunc:       sha1.New,
    derivedKeyFunc: derivedKeyWithPbkdf,
    saltSize:       20,
    keySize:        5,
    blockSize:      5,
    iterationCount: 2048,
    oid:            oidPbeWithSHA1AndRC4_40,
}

// pkcs5-v1.5 模式
var PEMCipherMD2AndDES = CipherBlockCBC{
    cipherFunc:     des.NewCipher,
    hashFunc:       cryptobin_md2.New,
    derivedKeyFunc: derivedKey,
    saltSize:       des.BlockSize,
    keySize:        8,
    blockSize:      des.BlockSize,
    iterationCount: 2048,
    oid:            oidPbeWithMD2AndDES,
}
var PEMCipherMD2AndRC2_64 = CipherBlockCBC{
    cipherFunc:     newRC2Cipher,
    hashFunc:       cryptobin_md2.New,
    derivedKeyFunc: derivedKey,
    saltSize:       20,
    keySize:        8,
    blockSize:      cryptobin_rc2.BlockSize,
    iterationCount: 2048,
    oid:            oidPbeWithMD2AndRC2_64,
}
var PEMCipherMD5AndDES = CipherBlockCBC{
    cipherFunc:     des.NewCipher,
    hashFunc:       md5.New,
    derivedKeyFunc: derivedKey,
    saltSize:       des.BlockSize,
    keySize:        8,
    blockSize:      des.BlockSize,
    iterationCount: 2048,
    oid:            oidPbeWithMD5AndDES,
}
var PEMCipherMD5AndRC2_64 = CipherBlockCBC{
    cipherFunc:     newRC2Cipher,
    hashFunc:       md5.New,
    derivedKeyFunc: derivedKey,
    saltSize:       20,
    keySize:        8,
    blockSize:      cryptobin_rc2.BlockSize,
    iterationCount: 2048,
    oid:            oidPbeWithMD5AndRC2_64,
}
var PEMCipherSHA1AndDES = CipherBlockCBC{
    cipherFunc:     des.NewCipher,
    hashFunc:       sha1.New,
    derivedKeyFunc: derivedKey,
    saltSize:       des.BlockSize,
    keySize:        8,
    blockSize:      des.BlockSize,
    iterationCount: 2048,
    oid:            oidPbeWithSHA1AndDES,
}
var PEMCipherSHA1AndRC2_64 = CipherBlockCBC{
    cipherFunc:     newRC2Cipher,
    hashFunc:       sha1.New,
    derivedKeyFunc: derivedKey,
    saltSize:       20,
    keySize:        8,
    blockSize:      cryptobin_rc2.BlockSize,
    iterationCount: 2048,
    oid:            oidPbeWithSHA1AndRC2_64,
}

func init() {
    // pkcs12 模式
    AddCipher(oidPbeWithSHA1And3DES, func() PEMCipher {
        return PEMCipherSHA1And3DES
    })
    AddCipher(oidPbeWithSHA1And2DES, func() PEMCipher {
        return PEMCipherSHA1And2DES
    })
    AddCipher(oidPbeWithSHA1AndRC2_128, func() PEMCipher {
        return PEMCipherSHA1AndRC2_128
    })
    AddCipher(oidPbeWithSHA1AndRC2_40, func() PEMCipher {
        return PEMCipherSHA1AndRC2_40
    })
    AddCipher(oidPbeWithSHA1AndRC4_128, func() PEMCipher {
        return PEMCipherSHA1AndRC4_128
    })
    AddCipher(oidPbeWithSHA1AndRC4_40, func() PEMCipher {
        return PEMCipherSHA1AndRC4_40
    })

    // pkcs5-v1.5 模式
    AddCipher(oidPbeWithMD2AndDES, func() PEMCipher {
        return PEMCipherMD2AndDES
    })
    AddCipher(oidPbeWithMD2AndRC2_64, func() PEMCipher {
        return PEMCipherMD2AndRC2_64
    })
    AddCipher(oidPbeWithMD5AndDES, func() PEMCipher {
        return PEMCipherMD5AndDES
    })
    AddCipher(oidPbeWithMD5AndRC2_64, func() PEMCipher {
        return PEMCipherMD5AndRC2_64
    })
    AddCipher(oidPbeWithSHA1AndDES, func() PEMCipher {
        return PEMCipherSHA1AndDES
    })
    AddCipher(oidPbeWithSHA1AndRC2_64, func() PEMCipher {
        return PEMCipherSHA1AndRC2_64
    })
}
