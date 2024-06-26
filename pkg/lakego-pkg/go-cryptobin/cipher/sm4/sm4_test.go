package sm4

import (
    "bytes"
    "reflect"
    "testing"
    "math/rand"
)

func Test_Cipher(t *testing.T) {
    random := rand.New(rand.NewSource(99))
    max := 100

    var encrypted [16]byte
    var decrypted [16]byte

    for i := 0; i < max; i++ {
        key := make([]byte, 16)
        random.Read(key)
        value := make([]byte, 16)
        random.Read(value)

        cipher1, err := NewCipher(key)
        if err != nil {
            t.Fatal(err.Error())
        }

        cipher1.Encrypt(encrypted[:], value)

        cipher2, err := NewCipher(key)
        if err != nil {
            t.Fatal(err.Error())
        }

        cipher2.Decrypt(decrypted[:], encrypted[:])

        if !bytes.Equal(decrypted[:], value[:]) {
            t.Errorf("encryption/decryption failed: % 02x != % 02x\n", decrypted, value)
        }
    }
}

func Test_Check(t *testing.T) {
    src := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}
    expected := []byte{0x68, 0x1e, 0xdf, 0x34, 0xd2, 0x06, 0x96, 0x5e, 0x86, 0xb3, 0xe9, 0x4f, 0x53, 0x6e, 0x42, 0x46}

    c, err := NewCipher(src)
    if err != nil {
        t.Fatal(err)
    }

    dst := make([]byte, 16)
    c.Encrypt(dst, src)
    if !reflect.DeepEqual(dst, expected) {
        t.Errorf("expected=%x, result=%x\n", expected, dst)
    }

    c.Decrypt(dst, expected)
    if !reflect.DeepEqual(dst, src) {
        t.Errorf("expected=%x, result=%x\n", src, dst)
    }
}

func Test_Check2(t *testing.T) {
    src := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}
    expected := []byte{0x59, 0x52, 0x98, 0xc7, 0xc6, 0xfd, 0x27, 0x1f, 0x04, 0x02, 0xf8, 0x04, 0xc3, 0x3d, 0x3f, 0x66}

    c, err := NewCipher(src)
    if err != nil {
        t.Fatal(err)
    }

    dst := make([]byte, 16)
    copy(dst, src)

    n := 1000000
    if testing.Short() {
        n = 1000
        expected = []byte{215, 53, 233, 28, 197, 104, 156, 243, 18, 188, 193, 239, 183, 64, 232, 19}
    }

    for i := 0; i < n; i++ {
        c.Encrypt(dst, dst)
    }

    if !reflect.DeepEqual(dst, expected) {
        t.Errorf("expected=%x, result=%x\n", expected, dst)
    }
}
