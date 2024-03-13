package xoodoo_test

import (
    "fmt"

    "github.com/deatil/go-cryptobin/cipher/xoodoo/xoodoo"
)

func Example() {
    newXoodoo, _ := xoodoo.NewXoodoo(xoodoo.MaxRounds, [xoodoo.StateSizeBytes]byte{})
    fmt.Printf("Starting State:%#v\n", newXoodoo.Bytes())
    newXoodoo.Permutation()
    fmt.Printf("Permuted State:%#v\n", newXoodoo.Bytes())
    // Output: Starting State:[]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
    // Permuted State:[]byte{0x8d, 0xd8, 0xd5, 0x89, 0xbf, 0xfc, 0x63, 0xa9, 0x19, 0x2d, 0x23, 0x1b, 0x14, 0xa0, 0xa5, 0xff, 0x6, 0x81, 0xb1, 0x36, 0xfe, 0xc1, 0xc7, 0xaf, 0xbe, 0x7c, 0xe5, 0xae, 0xbd, 0x40, 0x75, 0xa7, 0x70, 0xe8, 0x86, 0x2e, 0xc9, 0xb7, 0xf5, 0xfe, 0xf2, 0xad, 0x4f, 0x8b, 0x62, 0x40, 0x4f, 0x5e}
}
