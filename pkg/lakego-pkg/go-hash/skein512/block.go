package skein512

const keyScheduleParity = 0x1bd11bdaa9fc1a22

// block calculates one UBI block from src with key k and tweak t, and puts
// the result into dst. The source and destination arrays can be the same.
func block(k *[8]uint64, t *[2]uint64, dst, src *[8]uint64) {
    x0, x1, x2, x3, x4, x5, x6, x7 := src[0], src[1], src[2], src[3], src[4], src[5], src[6], src[7]
    k8 := keyScheduleParity ^ k[0] ^ k[1] ^ k[2] ^ k[3] ^ k[4] ^ k[5] ^ k[6] ^ k[7]
    t2 := t[0] ^ t[1]

    x0 += k[0]
    x1 += k[1]
    x2 += k[2]
    x3 += k[3]
    x4 += k[4]
    x5 += k[5] + t[0]
    x6 += k[6] + t[1]
    x7 += k[7]

    x0 += x1
    x1 = (x1<<46 | x1>>(64-46)) ^ x0
    x2 += x3
    x3 = (x3<<36 | x3>>(64-36)) ^ x2
    x4 += x5
    x5 = (x5<<19 | x5>>(64-19)) ^ x4
    x6 += x7
    x7 = (x7<<37 | x7>>(64-37)) ^ x6

    x2 += x1
    x1 = (x1<<33 | x1>>(64-33)) ^ x2
    x4 += x7
    x7 = (x7<<27 | x7>>(64-27)) ^ x4
    x6 += x5
    x5 = (x5<<14 | x5>>(64-14)) ^ x6
    x0 += x3
    x3 = (x3<<42 | x3>>(64-42)) ^ x0

    x4 += x1
    x1 = (x1<<17 | x1>>(64-17)) ^ x4
    x6 += x3
    x3 = (x3<<49 | x3>>(64-49)) ^ x6
    x0 += x5
    x5 = (x5<<36 | x5>>(64-36)) ^ x0
    x2 += x7
    x7 = (x7<<39 | x7>>(64-39)) ^ x2

    x6 += x1
    x1 = (x1<<44 | x1>>(64-44)) ^ x6
    x0 += x7
    x7 = (x7<<9 | x7>>(64-9)) ^ x0
    x2 += x5
    x5 = (x5<<54 | x5>>(64-54)) ^ x2
    x4 += x3
    x3 = (x3<<56 | x3>>(64-56)) ^ x4

    x0 += k[1]
    x1 += k[2]
    x2 += k[3]
    x3 += k[4]
    x4 += k[5]
    x5 += k[6] + t[1]
    x6 += k[7] + t2
    x7 += k8 + 1

    x0 += x1
    x1 = (x1<<39 | x1>>(64-39)) ^ x0
    x2 += x3
    x3 = (x3<<30 | x3>>(64-30)) ^ x2
    x4 += x5
    x5 = (x5<<34 | x5>>(64-34)) ^ x4
    x6 += x7
    x7 = (x7<<24 | x7>>(64-24)) ^ x6

    x2 += x1
    x1 = (x1<<13 | x1>>(64-13)) ^ x2
    x4 += x7
    x7 = (x7<<50 | x7>>(64-50)) ^ x4
    x6 += x5
    x5 = (x5<<10 | x5>>(64-10)) ^ x6
    x0 += x3
    x3 = (x3<<17 | x3>>(64-17)) ^ x0

    x4 += x1
    x1 = (x1<<25 | x1>>(64-25)) ^ x4
    x6 += x3
    x3 = (x3<<29 | x3>>(64-29)) ^ x6
    x0 += x5
    x5 = (x5<<39 | x5>>(64-39)) ^ x0
    x2 += x7
    x7 = (x7<<43 | x7>>(64-43)) ^ x2

    x6 += x1
    x1 = (x1<<8 | x1>>(64-8)) ^ x6
    x0 += x7
    x7 = (x7<<35 | x7>>(64-35)) ^ x0
    x2 += x5
    x5 = (x5<<56 | x5>>(64-56)) ^ x2
    x4 += x3
    x3 = (x3<<22 | x3>>(64-22)) ^ x4

    x0 += k[2]
    x1 += k[3]
    x2 += k[4]
    x3 += k[5]
    x4 += k[6]
    x5 += k[7] + t2
    x6 += k8 + t[0]
    x7 += k[0] + 2

    x0 += x1
    x1 = (x1<<46 | x1>>(64-46)) ^ x0
    x2 += x3
    x3 = (x3<<36 | x3>>(64-36)) ^ x2
    x4 += x5
    x5 = (x5<<19 | x5>>(64-19)) ^ x4
    x6 += x7
    x7 = (x7<<37 | x7>>(64-37)) ^ x6

    x2 += x1
    x1 = (x1<<33 | x1>>(64-33)) ^ x2
    x4 += x7
    x7 = (x7<<27 | x7>>(64-27)) ^ x4
    x6 += x5
    x5 = (x5<<14 | x5>>(64-14)) ^ x6
    x0 += x3
    x3 = (x3<<42 | x3>>(64-42)) ^ x0

    x4 += x1
    x1 = (x1<<17 | x1>>(64-17)) ^ x4
    x6 += x3
    x3 = (x3<<49 | x3>>(64-49)) ^ x6
    x0 += x5
    x5 = (x5<<36 | x5>>(64-36)) ^ x0
    x2 += x7
    x7 = (x7<<39 | x7>>(64-39)) ^ x2

    x6 += x1
    x1 = (x1<<44 | x1>>(64-44)) ^ x6
    x0 += x7
    x7 = (x7<<9 | x7>>(64-9)) ^ x0
    x2 += x5
    x5 = (x5<<54 | x5>>(64-54)) ^ x2
    x4 += x3
    x3 = (x3<<56 | x3>>(64-56)) ^ x4

    x0 += k[3]
    x1 += k[4]
    x2 += k[5]
    x3 += k[6]
    x4 += k[7]
    x5 += k8 + t[0]
    x6 += k[0] + t[1]
    x7 += k[1] + 3

    x0 += x1
    x1 = (x1<<39 | x1>>(64-39)) ^ x0
    x2 += x3
    x3 = (x3<<30 | x3>>(64-30)) ^ x2
    x4 += x5
    x5 = (x5<<34 | x5>>(64-34)) ^ x4
    x6 += x7
    x7 = (x7<<24 | x7>>(64-24)) ^ x6

    x2 += x1
    x1 = (x1<<13 | x1>>(64-13)) ^ x2
    x4 += x7
    x7 = (x7<<50 | x7>>(64-50)) ^ x4
    x6 += x5
    x5 = (x5<<10 | x5>>(64-10)) ^ x6
    x0 += x3
    x3 = (x3<<17 | x3>>(64-17)) ^ x0

    x4 += x1
    x1 = (x1<<25 | x1>>(64-25)) ^ x4
    x6 += x3
    x3 = (x3<<29 | x3>>(64-29)) ^ x6
    x0 += x5
    x5 = (x5<<39 | x5>>(64-39)) ^ x0
    x2 += x7
    x7 = (x7<<43 | x7>>(64-43)) ^ x2

    x6 += x1
    x1 = (x1<<8 | x1>>(64-8)) ^ x6
    x0 += x7
    x7 = (x7<<35 | x7>>(64-35)) ^ x0
    x2 += x5
    x5 = (x5<<56 | x5>>(64-56)) ^ x2
    x4 += x3
    x3 = (x3<<22 | x3>>(64-22)) ^ x4

    x0 += k[4]
    x1 += k[5]
    x2 += k[6]
    x3 += k[7]
    x4 += k8
    x5 += k[0] + t[1]
    x6 += k[1] + t2
    x7 += k[2] + 4

    x0 += x1
    x1 = (x1<<46 | x1>>(64-46)) ^ x0
    x2 += x3
    x3 = (x3<<36 | x3>>(64-36)) ^ x2
    x4 += x5
    x5 = (x5<<19 | x5>>(64-19)) ^ x4
    x6 += x7
    x7 = (x7<<37 | x7>>(64-37)) ^ x6

    x2 += x1
    x1 = (x1<<33 | x1>>(64-33)) ^ x2
    x4 += x7
    x7 = (x7<<27 | x7>>(64-27)) ^ x4
    x6 += x5
    x5 = (x5<<14 | x5>>(64-14)) ^ x6
    x0 += x3
    x3 = (x3<<42 | x3>>(64-42)) ^ x0

    x4 += x1
    x1 = (x1<<17 | x1>>(64-17)) ^ x4
    x6 += x3
    x3 = (x3<<49 | x3>>(64-49)) ^ x6
    x0 += x5
    x5 = (x5<<36 | x5>>(64-36)) ^ x0
    x2 += x7
    x7 = (x7<<39 | x7>>(64-39)) ^ x2

    x6 += x1
    x1 = (x1<<44 | x1>>(64-44)) ^ x6
    x0 += x7
    x7 = (x7<<9 | x7>>(64-9)) ^ x0
    x2 += x5
    x5 = (x5<<54 | x5>>(64-54)) ^ x2
    x4 += x3
    x3 = (x3<<56 | x3>>(64-56)) ^ x4

    x0 += k[5]
    x1 += k[6]
    x2 += k[7]
    x3 += k8
    x4 += k[0]
    x5 += k[1] + t2
    x6 += k[2] + t[0]
    x7 += k[3] + 5

    x0 += x1
    x1 = (x1<<39 | x1>>(64-39)) ^ x0
    x2 += x3
    x3 = (x3<<30 | x3>>(64-30)) ^ x2
    x4 += x5
    x5 = (x5<<34 | x5>>(64-34)) ^ x4
    x6 += x7
    x7 = (x7<<24 | x7>>(64-24)) ^ x6

    x2 += x1
    x1 = (x1<<13 | x1>>(64-13)) ^ x2
    x4 += x7
    x7 = (x7<<50 | x7>>(64-50)) ^ x4
    x6 += x5
    x5 = (x5<<10 | x5>>(64-10)) ^ x6
    x0 += x3
    x3 = (x3<<17 | x3>>(64-17)) ^ x0

    x4 += x1
    x1 = (x1<<25 | x1>>(64-25)) ^ x4
    x6 += x3
    x3 = (x3<<29 | x3>>(64-29)) ^ x6
    x0 += x5
    x5 = (x5<<39 | x5>>(64-39)) ^ x0
    x2 += x7
    x7 = (x7<<43 | x7>>(64-43)) ^ x2

    x6 += x1
    x1 = (x1<<8 | x1>>(64-8)) ^ x6
    x0 += x7
    x7 = (x7<<35 | x7>>(64-35)) ^ x0
    x2 += x5
    x5 = (x5<<56 | x5>>(64-56)) ^ x2
    x4 += x3
    x3 = (x3<<22 | x3>>(64-22)) ^ x4

    x0 += k[6]
    x1 += k[7]
    x2 += k8
    x3 += k[0]
    x4 += k[1]
    x5 += k[2] + t[0]
    x6 += k[3] + t[1]
    x7 += k[4] + 6

    x0 += x1
    x1 = (x1<<46 | x1>>(64-46)) ^ x0
    x2 += x3
    x3 = (x3<<36 | x3>>(64-36)) ^ x2
    x4 += x5
    x5 = (x5<<19 | x5>>(64-19)) ^ x4
    x6 += x7
    x7 = (x7<<37 | x7>>(64-37)) ^ x6

    x2 += x1
    x1 = (x1<<33 | x1>>(64-33)) ^ x2
    x4 += x7
    x7 = (x7<<27 | x7>>(64-27)) ^ x4
    x6 += x5
    x5 = (x5<<14 | x5>>(64-14)) ^ x6
    x0 += x3
    x3 = (x3<<42 | x3>>(64-42)) ^ x0

    x4 += x1
    x1 = (x1<<17 | x1>>(64-17)) ^ x4
    x6 += x3
    x3 = (x3<<49 | x3>>(64-49)) ^ x6
    x0 += x5
    x5 = (x5<<36 | x5>>(64-36)) ^ x0
    x2 += x7
    x7 = (x7<<39 | x7>>(64-39)) ^ x2

    x6 += x1
    x1 = (x1<<44 | x1>>(64-44)) ^ x6
    x0 += x7
    x7 = (x7<<9 | x7>>(64-9)) ^ x0
    x2 += x5
    x5 = (x5<<54 | x5>>(64-54)) ^ x2
    x4 += x3
    x3 = (x3<<56 | x3>>(64-56)) ^ x4

    x0 += k[7]
    x1 += k8
    x2 += k[0]
    x3 += k[1]
    x4 += k[2]
    x5 += k[3] + t[1]
    x6 += k[4] + t2
    x7 += k[5] + 7

    x0 += x1
    x1 = (x1<<39 | x1>>(64-39)) ^ x0
    x2 += x3
    x3 = (x3<<30 | x3>>(64-30)) ^ x2
    x4 += x5
    x5 = (x5<<34 | x5>>(64-34)) ^ x4
    x6 += x7
    x7 = (x7<<24 | x7>>(64-24)) ^ x6

    x2 += x1
    x1 = (x1<<13 | x1>>(64-13)) ^ x2
    x4 += x7
    x7 = (x7<<50 | x7>>(64-50)) ^ x4
    x6 += x5
    x5 = (x5<<10 | x5>>(64-10)) ^ x6
    x0 += x3
    x3 = (x3<<17 | x3>>(64-17)) ^ x0

    x4 += x1
    x1 = (x1<<25 | x1>>(64-25)) ^ x4
    x6 += x3
    x3 = (x3<<29 | x3>>(64-29)) ^ x6
    x0 += x5
    x5 = (x5<<39 | x5>>(64-39)) ^ x0
    x2 += x7
    x7 = (x7<<43 | x7>>(64-43)) ^ x2

    x6 += x1
    x1 = (x1<<8 | x1>>(64-8)) ^ x6
    x0 += x7
    x7 = (x7<<35 | x7>>(64-35)) ^ x0
    x2 += x5
    x5 = (x5<<56 | x5>>(64-56)) ^ x2
    x4 += x3
    x3 = (x3<<22 | x3>>(64-22)) ^ x4

    x0 += k8
    x1 += k[0]
    x2 += k[1]
    x3 += k[2]
    x4 += k[3]
    x5 += k[4] + t2
    x6 += k[5] + t[0]
    x7 += k[6] + 8

    x0 += x1
    x1 = (x1<<46 | x1>>(64-46)) ^ x0
    x2 += x3
    x3 = (x3<<36 | x3>>(64-36)) ^ x2
    x4 += x5
    x5 = (x5<<19 | x5>>(64-19)) ^ x4
    x6 += x7
    x7 = (x7<<37 | x7>>(64-37)) ^ x6

    x2 += x1
    x1 = (x1<<33 | x1>>(64-33)) ^ x2
    x4 += x7
    x7 = (x7<<27 | x7>>(64-27)) ^ x4
    x6 += x5
    x5 = (x5<<14 | x5>>(64-14)) ^ x6
    x0 += x3
    x3 = (x3<<42 | x3>>(64-42)) ^ x0

    x4 += x1
    x1 = (x1<<17 | x1>>(64-17)) ^ x4
    x6 += x3
    x3 = (x3<<49 | x3>>(64-49)) ^ x6
    x0 += x5
    x5 = (x5<<36 | x5>>(64-36)) ^ x0
    x2 += x7
    x7 = (x7<<39 | x7>>(64-39)) ^ x2

    x6 += x1
    x1 = (x1<<44 | x1>>(64-44)) ^ x6
    x0 += x7
    x7 = (x7<<9 | x7>>(64-9)) ^ x0
    x2 += x5
    x5 = (x5<<54 | x5>>(64-54)) ^ x2
    x4 += x3
    x3 = (x3<<56 | x3>>(64-56)) ^ x4

    x0 += k[0]
    x1 += k[1]
    x2 += k[2]
    x3 += k[3]
    x4 += k[4]
    x5 += k[5] + t[0]
    x6 += k[6] + t[1]
    x7 += k[7] + 9

    x0 += x1
    x1 = (x1<<39 | x1>>(64-39)) ^ x0
    x2 += x3
    x3 = (x3<<30 | x3>>(64-30)) ^ x2
    x4 += x5
    x5 = (x5<<34 | x5>>(64-34)) ^ x4
    x6 += x7
    x7 = (x7<<24 | x7>>(64-24)) ^ x6

    x2 += x1
    x1 = (x1<<13 | x1>>(64-13)) ^ x2
    x4 += x7
    x7 = (x7<<50 | x7>>(64-50)) ^ x4
    x6 += x5
    x5 = (x5<<10 | x5>>(64-10)) ^ x6
    x0 += x3
    x3 = (x3<<17 | x3>>(64-17)) ^ x0

    x4 += x1
    x1 = (x1<<25 | x1>>(64-25)) ^ x4
    x6 += x3
    x3 = (x3<<29 | x3>>(64-29)) ^ x6
    x0 += x5
    x5 = (x5<<39 | x5>>(64-39)) ^ x0
    x2 += x7
    x7 = (x7<<43 | x7>>(64-43)) ^ x2

    x6 += x1
    x1 = (x1<<8 | x1>>(64-8)) ^ x6
    x0 += x7
    x7 = (x7<<35 | x7>>(64-35)) ^ x0
    x2 += x5
    x5 = (x5<<56 | x5>>(64-56)) ^ x2
    x4 += x3
    x3 = (x3<<22 | x3>>(64-22)) ^ x4

    x0 += k[1]
    x1 += k[2]
    x2 += k[3]
    x3 += k[4]
    x4 += k[5]
    x5 += k[6] + t[1]
    x6 += k[7] + t2
    x7 += k8 + 10

    x0 += x1
    x1 = (x1<<46 | x1>>(64-46)) ^ x0
    x2 += x3
    x3 = (x3<<36 | x3>>(64-36)) ^ x2
    x4 += x5
    x5 = (x5<<19 | x5>>(64-19)) ^ x4
    x6 += x7
    x7 = (x7<<37 | x7>>(64-37)) ^ x6

    x2 += x1
    x1 = (x1<<33 | x1>>(64-33)) ^ x2
    x4 += x7
    x7 = (x7<<27 | x7>>(64-27)) ^ x4
    x6 += x5
    x5 = (x5<<14 | x5>>(64-14)) ^ x6
    x0 += x3
    x3 = (x3<<42 | x3>>(64-42)) ^ x0

    x4 += x1
    x1 = (x1<<17 | x1>>(64-17)) ^ x4
    x6 += x3
    x3 = (x3<<49 | x3>>(64-49)) ^ x6
    x0 += x5
    x5 = (x5<<36 | x5>>(64-36)) ^ x0
    x2 += x7
    x7 = (x7<<39 | x7>>(64-39)) ^ x2

    x6 += x1
    x1 = (x1<<44 | x1>>(64-44)) ^ x6
    x0 += x7
    x7 = (x7<<9 | x7>>(64-9)) ^ x0
    x2 += x5
    x5 = (x5<<54 | x5>>(64-54)) ^ x2
    x4 += x3
    x3 = (x3<<56 | x3>>(64-56)) ^ x4

    x0 += k[2]
    x1 += k[3]
    x2 += k[4]
    x3 += k[5]
    x4 += k[6]
    x5 += k[7] + t2
    x6 += k8 + t[0]
    x7 += k[0] + 11

    x0 += x1
    x1 = (x1<<39 | x1>>(64-39)) ^ x0
    x2 += x3
    x3 = (x3<<30 | x3>>(64-30)) ^ x2
    x4 += x5
    x5 = (x5<<34 | x5>>(64-34)) ^ x4
    x6 += x7
    x7 = (x7<<24 | x7>>(64-24)) ^ x6

    x2 += x1
    x1 = (x1<<13 | x1>>(64-13)) ^ x2
    x4 += x7
    x7 = (x7<<50 | x7>>(64-50)) ^ x4
    x6 += x5
    x5 = (x5<<10 | x5>>(64-10)) ^ x6
    x0 += x3
    x3 = (x3<<17 | x3>>(64-17)) ^ x0

    x4 += x1
    x1 = (x1<<25 | x1>>(64-25)) ^ x4
    x6 += x3
    x3 = (x3<<29 | x3>>(64-29)) ^ x6
    x0 += x5
    x5 = (x5<<39 | x5>>(64-39)) ^ x0
    x2 += x7
    x7 = (x7<<43 | x7>>(64-43)) ^ x2

    x6 += x1
    x1 = (x1<<8 | x1>>(64-8)) ^ x6
    x0 += x7
    x7 = (x7<<35 | x7>>(64-35)) ^ x0
    x2 += x5
    x5 = (x5<<56 | x5>>(64-56)) ^ x2
    x4 += x3
    x3 = (x3<<22 | x3>>(64-22)) ^ x4

    x0 += k[3]
    x1 += k[4]
    x2 += k[5]
    x3 += k[6]
    x4 += k[7]
    x5 += k8 + t[0]
    x6 += k[0] + t[1]
    x7 += k[1] + 12

    x0 += x1
    x1 = (x1<<46 | x1>>(64-46)) ^ x0
    x2 += x3
    x3 = (x3<<36 | x3>>(64-36)) ^ x2
    x4 += x5
    x5 = (x5<<19 | x5>>(64-19)) ^ x4
    x6 += x7
    x7 = (x7<<37 | x7>>(64-37)) ^ x6

    x2 += x1
    x1 = (x1<<33 | x1>>(64-33)) ^ x2
    x4 += x7
    x7 = (x7<<27 | x7>>(64-27)) ^ x4
    x6 += x5
    x5 = (x5<<14 | x5>>(64-14)) ^ x6
    x0 += x3
    x3 = (x3<<42 | x3>>(64-42)) ^ x0

    x4 += x1
    x1 = (x1<<17 | x1>>(64-17)) ^ x4
    x6 += x3
    x3 = (x3<<49 | x3>>(64-49)) ^ x6
    x0 += x5
    x5 = (x5<<36 | x5>>(64-36)) ^ x0
    x2 += x7
    x7 = (x7<<39 | x7>>(64-39)) ^ x2

    x6 += x1
    x1 = (x1<<44 | x1>>(64-44)) ^ x6
    x0 += x7
    x7 = (x7<<9 | x7>>(64-9)) ^ x0
    x2 += x5
    x5 = (x5<<54 | x5>>(64-54)) ^ x2
    x4 += x3
    x3 = (x3<<56 | x3>>(64-56)) ^ x4

    x0 += k[4]
    x1 += k[5]
    x2 += k[6]
    x3 += k[7]
    x4 += k8
    x5 += k[0] + t[1]
    x6 += k[1] + t2
    x7 += k[2] + 13

    x0 += x1
    x1 = (x1<<39 | x1>>(64-39)) ^ x0
    x2 += x3
    x3 = (x3<<30 | x3>>(64-30)) ^ x2
    x4 += x5
    x5 = (x5<<34 | x5>>(64-34)) ^ x4
    x6 += x7
    x7 = (x7<<24 | x7>>(64-24)) ^ x6

    x2 += x1
    x1 = (x1<<13 | x1>>(64-13)) ^ x2
    x4 += x7
    x7 = (x7<<50 | x7>>(64-50)) ^ x4
    x6 += x5
    x5 = (x5<<10 | x5>>(64-10)) ^ x6
    x0 += x3
    x3 = (x3<<17 | x3>>(64-17)) ^ x0

    x4 += x1
    x1 = (x1<<25 | x1>>(64-25)) ^ x4
    x6 += x3
    x3 = (x3<<29 | x3>>(64-29)) ^ x6
    x0 += x5
    x5 = (x5<<39 | x5>>(64-39)) ^ x0
    x2 += x7
    x7 = (x7<<43 | x7>>(64-43)) ^ x2

    x6 += x1
    x1 = (x1<<8 | x1>>(64-8)) ^ x6
    x0 += x7
    x7 = (x7<<35 | x7>>(64-35)) ^ x0
    x2 += x5
    x5 = (x5<<56 | x5>>(64-56)) ^ x2
    x4 += x3
    x3 = (x3<<22 | x3>>(64-22)) ^ x4

    x0 += k[5]
    x1 += k[6]
    x2 += k[7]
    x3 += k8
    x4 += k[0]
    x5 += k[1] + t2
    x6 += k[2] + t[0]
    x7 += k[3] + 14

    x0 += x1
    x1 = (x1<<46 | x1>>(64-46)) ^ x0
    x2 += x3
    x3 = (x3<<36 | x3>>(64-36)) ^ x2
    x4 += x5
    x5 = (x5<<19 | x5>>(64-19)) ^ x4
    x6 += x7
    x7 = (x7<<37 | x7>>(64-37)) ^ x6

    x2 += x1
    x1 = (x1<<33 | x1>>(64-33)) ^ x2
    x4 += x7
    x7 = (x7<<27 | x7>>(64-27)) ^ x4
    x6 += x5
    x5 = (x5<<14 | x5>>(64-14)) ^ x6
    x0 += x3
    x3 = (x3<<42 | x3>>(64-42)) ^ x0

    x4 += x1
    x1 = (x1<<17 | x1>>(64-17)) ^ x4
    x6 += x3
    x3 = (x3<<49 | x3>>(64-49)) ^ x6
    x0 += x5
    x5 = (x5<<36 | x5>>(64-36)) ^ x0
    x2 += x7
    x7 = (x7<<39 | x7>>(64-39)) ^ x2

    x6 += x1
    x1 = (x1<<44 | x1>>(64-44)) ^ x6
    x0 += x7
    x7 = (x7<<9 | x7>>(64-9)) ^ x0
    x2 += x5
    x5 = (x5<<54 | x5>>(64-54)) ^ x2
    x4 += x3
    x3 = (x3<<56 | x3>>(64-56)) ^ x4

    x0 += k[6]
    x1 += k[7]
    x2 += k8
    x3 += k[0]
    x4 += k[1]
    x5 += k[2] + t[0]
    x6 += k[3] + t[1]
    x7 += k[4] + 15

    x0 += x1
    x1 = (x1<<39 | x1>>(64-39)) ^ x0
    x2 += x3
    x3 = (x3<<30 | x3>>(64-30)) ^ x2
    x4 += x5
    x5 = (x5<<34 | x5>>(64-34)) ^ x4
    x6 += x7
    x7 = (x7<<24 | x7>>(64-24)) ^ x6

    x2 += x1
    x1 = (x1<<13 | x1>>(64-13)) ^ x2
    x4 += x7
    x7 = (x7<<50 | x7>>(64-50)) ^ x4
    x6 += x5
    x5 = (x5<<10 | x5>>(64-10)) ^ x6
    x0 += x3
    x3 = (x3<<17 | x3>>(64-17)) ^ x0

    x4 += x1
    x1 = (x1<<25 | x1>>(64-25)) ^ x4
    x6 += x3
    x3 = (x3<<29 | x3>>(64-29)) ^ x6
    x0 += x5
    x5 = (x5<<39 | x5>>(64-39)) ^ x0
    x2 += x7
    x7 = (x7<<43 | x7>>(64-43)) ^ x2

    x6 += x1
    x1 = (x1<<8 | x1>>(64-8)) ^ x6
    x0 += x7
    x7 = (x7<<35 | x7>>(64-35)) ^ x0
    x2 += x5
    x5 = (x5<<56 | x5>>(64-56)) ^ x2
    x4 += x3
    x3 = (x3<<22 | x3>>(64-22)) ^ x4

    x0 += k[7]
    x1 += k8
    x2 += k[0]
    x3 += k[1]
    x4 += k[2]
    x5 += k[3] + t[1]
    x6 += k[4] + t2
    x7 += k[5] + 16

    x0 += x1
    x1 = (x1<<46 | x1>>(64-46)) ^ x0
    x2 += x3
    x3 = (x3<<36 | x3>>(64-36)) ^ x2
    x4 += x5
    x5 = (x5<<19 | x5>>(64-19)) ^ x4
    x6 += x7
    x7 = (x7<<37 | x7>>(64-37)) ^ x6

    x2 += x1
    x1 = (x1<<33 | x1>>(64-33)) ^ x2
    x4 += x7
    x7 = (x7<<27 | x7>>(64-27)) ^ x4
    x6 += x5
    x5 = (x5<<14 | x5>>(64-14)) ^ x6
    x0 += x3
    x3 = (x3<<42 | x3>>(64-42)) ^ x0

    x4 += x1
    x1 = (x1<<17 | x1>>(64-17)) ^ x4
    x6 += x3
    x3 = (x3<<49 | x3>>(64-49)) ^ x6
    x0 += x5
    x5 = (x5<<36 | x5>>(64-36)) ^ x0
    x2 += x7
    x7 = (x7<<39 | x7>>(64-39)) ^ x2

    x6 += x1
    x1 = (x1<<44 | x1>>(64-44)) ^ x6
    x0 += x7
    x7 = (x7<<9 | x7>>(64-9)) ^ x0
    x2 += x5
    x5 = (x5<<54 | x5>>(64-54)) ^ x2
    x4 += x3
    x3 = (x3<<56 | x3>>(64-56)) ^ x4

    x0 += k8
    x1 += k[0]
    x2 += k[1]
    x3 += k[2]
    x4 += k[3]
    x5 += k[4] + t2
    x6 += k[5] + t[0]
    x7 += k[6] + 17

    x0 += x1
    x1 = (x1<<39 | x1>>(64-39)) ^ x0
    x2 += x3
    x3 = (x3<<30 | x3>>(64-30)) ^ x2
    x4 += x5
    x5 = (x5<<34 | x5>>(64-34)) ^ x4
    x6 += x7
    x7 = (x7<<24 | x7>>(64-24)) ^ x6

    x2 += x1
    x1 = (x1<<13 | x1>>(64-13)) ^ x2
    x4 += x7
    x7 = (x7<<50 | x7>>(64-50)) ^ x4
    x6 += x5
    x5 = (x5<<10 | x5>>(64-10)) ^ x6
    x0 += x3
    x3 = (x3<<17 | x3>>(64-17)) ^ x0

    x4 += x1
    x1 = (x1<<25 | x1>>(64-25)) ^ x4
    x6 += x3
    x3 = (x3<<29 | x3>>(64-29)) ^ x6
    x0 += x5
    x5 = (x5<<39 | x5>>(64-39)) ^ x0
    x2 += x7
    x7 = (x7<<43 | x7>>(64-43)) ^ x2

    x6 += x1
    x1 = (x1<<8 | x1>>(64-8)) ^ x6
    x0 += x7
    x7 = (x7<<35 | x7>>(64-35)) ^ x0
    x2 += x5
    x5 = (x5<<56 | x5>>(64-56)) ^ x2
    x4 += x3
    x3 = (x3<<22 | x3>>(64-22)) ^ x4

    x0 += k[0]
    x1 += k[1]
    x2 += k[2]
    x3 += k[3]
    x4 += k[4]
    x5 += k[5] + t[0]
    x6 += k[6] + t[1]
    x7 += k[7] + 18

    dst[0] = src[0] ^ x0
    dst[1] = src[1] ^ x1
    dst[2] = src[2] ^ x2
    dst[3] = src[3] ^ x3
    dst[4] = src[4] ^ x4
    dst[5] = src[5] ^ x5
    dst[6] = src[6] ^ x6
    dst[7] = src[7] ^ x7
}