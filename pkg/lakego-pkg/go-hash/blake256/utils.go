package blake256

import (
    "math/bits"
    "encoding/binary"
)

// Endianness option
const littleEndian bool = false

func putu64(ptr []byte, a uint64) {
    if littleEndian {
        binary.LittleEndian.PutUint64(ptr, a)
    } else {
        binary.BigEndian.PutUint64(ptr, a)
    }
}

func bytesToUint32s(b []byte) []uint32 {
    size := len(b) / 4
    dst := make([]uint32, size)

    for i := 0; i < size; i++ {
        j := i * 4

        if littleEndian {
            dst[i] = binary.LittleEndian.Uint32(b[j:])
        } else {
            dst[i] = binary.BigEndian.Uint32(b[j:])
        }
    }

    return dst
}

func uint32sToBytes(w []uint32) []byte {
    size := len(w) * 4
    dst := make([]byte, size)

    for i := 0; i < len(w); i++ {
        j := i * 4

        if littleEndian {
            binary.LittleEndian.PutUint32(dst[j:], w[i])
        } else {
            binary.BigEndian.PutUint32(dst[j:], w[i])
        }
    }

    return dst
}

func rotr(x uint32, n int) uint32 {
    return bits.RotateLeft32(x, 32 - n)
}

func G(v *[16]uint32, m []uint32, i int, a, b, c, d, e int) {
    v[a] += (m[sigma[i][e]] ^ u256[sigma[i][e+1]]) + v[b]
    v[d] = rotr(v[d] ^ v[a], 16)
    v[c] += v[d]
    v[b] = rotr(v[b] ^ v[c], 12)
    v[a] += (m[sigma[i][e+1]] ^ u256[sigma[i][e]])+v[b]
    v[d] = rotr(v[d] ^ v[a], 8)
    v[c] += v[d]
    v[b] = rotr(v[b] ^ v[c], 7)
}
