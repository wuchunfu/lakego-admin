package x509

import (
    "fmt"
    "crypto"
    "errors"
    "unicode"

    "golang.org/x/crypto/cryptobyte"
    cryptobyte_asn1 "golang.org/x/crypto/cryptobyte/asn1"
)

func forEachSAN(der cryptobyte.String, callback func(tag int, data []byte) error) error {
    if !der.ReadASN1(&der, cryptobyte_asn1.SEQUENCE) {
        return errors.New("x509: invalid subject alternative names")
    }

    for !der.Empty() {
        var san cryptobyte.String
        var tag cryptobyte_asn1.Tag
        if !der.ReadAnyASN1(&san, &tag) {
            return errors.New("x509: invalid subject alternative name")
        }

        if err := callback(int(tag^0x80), san); err != nil {
            return err
        }
    }

    return nil
}

func isIA5String(s string) error {
    for _, r := range s {
        // Per RFC5280 "IA5String is limited to the set of ASCII characters"
        if r > unicode.MaxASCII {
            return fmt.Errorf("x509: %q cannot be encoded as an IA5String", s)
        }
    }

    return nil
}

// isValidIPMask reports whether mask consists of zero or more 1 bits, followed by zero bits.
func isValidIPMask(mask []byte) bool {
    seenZero := false

    for _, b := range mask {
        if seenZero {
            if b != 0 {
                return false
            }

            continue
        }

        switch b {
            case 0x00, 0x80, 0xc0, 0xe0, 0xf0, 0xf8, 0xfc, 0xfe:
                seenZero = true
            case 0xff:
            default:
                return false
        }
    }

    return true
}

func isRSASignHash(h crypto.Hash) bool {
    switch h {
        case crypto.MD5,
            crypto.SHA1,
            crypto.SHA224,
            crypto.SHA256,
            crypto.SHA384,
            crypto.SHA512,
            crypto.MD5SHA1,
            crypto.RIPEMD160:
            return true
        default:
            return false
    }
}
