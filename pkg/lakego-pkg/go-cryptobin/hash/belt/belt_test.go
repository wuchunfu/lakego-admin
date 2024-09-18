package belt

import (
    "bytes"
    "testing"
    "encoding/hex"
)

func fromHex(s string) []byte {
    h, _ := hex.DecodeString(s)
    return h
}

func Test_Hash(t *testing.T) {
    msg := []byte("test-datatest-datatest-datatest-datatest-datatest-datatest-datatest-datatest-datatest-datatest-datatest-datatest-datatest-datatest-datatest-datatest-datatest-datatest-datatest-datatest-datatest-datatest-datatest-datatest-datatest-datatest-datatest-datatest-datatest-data")

    h := New()
    h.Write(msg)
    dst := h.Sum(nil)

    if len(dst) == 0 {
        t.Error("Hash make error")
    }
}

type testData struct {
    msg []byte
    md []byte
}

func Test_Hash_Check(t *testing.T) {
   tests := []testData{
        {
            []byte{
                0xB1, 0x94, 0xBA, 0xC8, 0x0A, 0x08, 0xF5, 0x3B,
                0x36, 0x6D, 0x00, 0x8E, 0x58,
            },
            []byte{
                0xAB, 0xEF, 0x97, 0x25, 0xD4, 0xC5, 0xA8, 0x35,
                0x97, 0xA3, 0x67, 0xD1, 0x44, 0x94, 0xCC, 0x25,
                0x42, 0xF2, 0x0F, 0x65, 0x9D, 0xDF, 0xEC, 0xC9,
                0x61, 0xA3, 0xEC, 0x55, 0x0C, 0xBA, 0x8C, 0x75,
            },
        },
        {
            []byte{
                0xB1, 0x94, 0xBA, 0xC8, 0x0A, 0x08, 0xF5, 0x3B,
                0x36, 0x6D, 0x00, 0x8E, 0x58, 0x4A, 0x5D, 0xE4,
                0x85, 0x04, 0xFA, 0x9D, 0x1B, 0xB6, 0xC7, 0xAC,
                0x25, 0x2E, 0x72, 0xC2, 0x02, 0xFD, 0xCE, 0x0D,
            },
            []byte{
                0x74, 0x9E, 0x4C, 0x36, 0x53, 0xAE, 0xCE, 0x5E,
                0x48, 0xDB, 0x47, 0x61, 0x22, 0x77, 0x42, 0xEB,
                0x6D, 0xBE, 0x13, 0xF4, 0xA8, 0x0F, 0x7B, 0xEF,
                0xF1, 0xA9, 0xCF, 0x8D, 0x10, 0xEE, 0x77, 0x86,
            },
        },
        {
            []byte{
                0xB1, 0x94, 0xBA, 0xC8, 0x0A, 0x08, 0xF5, 0x3B,
                0x36, 0x6D, 0x00, 0x8E, 0x58, 0x4A, 0x5D, 0xE4,
                0x85, 0x04, 0xFA, 0x9D, 0x1B, 0xB6, 0xC7, 0xAC,
                0x25, 0x2E, 0x72, 0xC2, 0x02, 0xFD, 0xCE, 0x0D,
                0x5B, 0xE3, 0xD6, 0x12, 0x17, 0xB9, 0x61, 0x81,
                0xFE, 0x67, 0x86, 0xAD, 0x71, 0x6B, 0x89, 0x0B,
            },
            []byte{
                0x9D, 0x02, 0xEE, 0x44, 0x6F, 0xB6, 0xA2, 0x9F,
                0xE5, 0xC9, 0x82, 0xD4, 0xB1, 0x3A, 0xF9, 0xD3,
                0xE9, 0x08, 0x61, 0xBC, 0x4C, 0xEF, 0x27, 0xCF,
                0x30, 0x6B, 0xFB, 0x0B, 0x17, 0x4A, 0x15, 0x4A,
            },
        },

        {
            fromHex("b194bac80a08f53b366d008e584a5de48504fa9d1bb6c7ac252e72c202fdce0d5be3d61217b96181fe6786ad716b890b5cb0c0ff33c356b835c405aed8e07f99e12bdc1ae28257ec703fccf095ee8df1c1ab76389fe678caf7c6f860d5bb9c4f"),
            fromHex("c2fcd359337235d240e6498969ea3f5c73c8967ea4923d8476a62944573b7e87"),
        },

    }

    h := New()

    for i, test := range tests {
        h.Reset()
        h.Write(test.msg)
        sum := h.Sum(nil)

        if !bytes.Equal(sum, test.md) {
            t.Errorf("[%d] New fail, got %x, want %x", i, sum, test.md)
        }

        // =====

        sum2 := Sum(test.msg)

        if !bytes.Equal(sum2[:], test.md) {
            t.Errorf("[%d] Sum fail, got %x, want %x", i, sum2, test.md)
        }
    }
}

func Test_Cipher(t *testing.T) {
    k := []byte("12345678ujikolkj")

    var ks [BELT_KEY_SCHED_LEN]byte
    belt_init(k, &ks)

    data := []byte("1111111122222222")

    var in, out, dst [BELT_BLOCK_LEN]byte
    copy(in[:], data)

    belt_encrypt(in, &out, ks)
    belt_decrypt(out, &dst, ks)

    if !bytes.Equal(dst[:], data) {
        t.Errorf("decrypt fail, got %x, want %x", dst, data)
    }
}
