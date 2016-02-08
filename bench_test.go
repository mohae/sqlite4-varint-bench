package sqlite4_bench

import (
	"testing"

	"github.com/dchest/varuint"
	"github.com/mohae/uvarint"
)

var tests = []struct {
	decoded uint64
	n       int
	encoded []byte
}{
	{0, 1, []byte{0x00}},
	{1, 1, []byte{0x01}},
	{240, 1, []byte{0xF0}},
	{241, 2, []byte{0xF1, 0x01}},
	{2287, 2, []byte{0xF8, 0xFF}},

	{2288, 3, []byte{0xF9, 0x00, 0x00}},
	{67823, 3, []byte{0xF9, 0xFF, 0xFF}},
	{67824, 4, []byte{0xFA, 0x01, 0x08, 0xF0}},
	{1<<24 - 1, 4, []byte{0xFA, 0xFF, 0xFF, 0xFF}},
	{1 << 24, 5, []byte{0xFB, 0x01, 0x00, 0x00, 0x00}},

	{1<<32 - 1, 5, []byte{0xFB, 0xFF, 0xFF, 0xFF, 0xFF}},
	{1 << 32, 6, []byte{0xFC, 0x01, 0x00, 0x00, 0x00, 0x00}},
	{1<<40 - 1, 6, []byte{0xFC, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
	{1 << 40, 7, []byte{0xFD, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00}},
	{1<<48 - 1, 7, []byte{0xFD, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},

	{1 << 48, 8, []byte{0xFE, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
	{1<<56 - 1, 8, []byte{0xFE, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
	{1 << 56, 9, []byte{0xFF, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
	{1<<64 - 1, 9, []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
}

// Bench for all tests
func BenchmarkCockroachDBPutUvarintAll(b *testing.B) {
	buf := make([]byte, 9)
	var n int
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			n = putUvarint(buf, test.decoded)
		}
	}
	_ = n
}

func BenchmarkCockroachDBUvarintAll(b *testing.B) {
	var res uint64
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			res, _ = getUvarint(test.encoded)
		}
	}
	_ = res
}

func BenchmarkDchestPutUint64All(b *testing.B) {
	buf := make([]byte, 9)
	var n int
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			n = varuint.PutUint64(buf, test.decoded)
		}
	}
	_ = n
}

func BenchmarkDchestUint64All(b *testing.B) {
	var res uint64
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			res, _ = varuint.Uint64(test.encoded)
		}
	}
	_ = res
}

func BenchmarkMohaePutUvarintAll(b *testing.B) {
	buf := make([]byte, 9)
	var n int
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			n = uvarint.PutUvarint(buf, test.decoded)
		}
	}
	_ = n
}

func BenchmarkMohaeUvarintAll(b *testing.B) {
	var res uint64
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			res, _ = uvarint.Uvarint(test.encoded)
		}
	}
	_ = res
}

// Single encode: use < 241 (fastest)
func BenchmarkCockroachDBPutUvarintMinLen(b *testing.B) {
	buf := make([]byte, 9)
	var n int
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		n = putUvarint(buf, tests[2].decoded)
	}
	_ = n
}

func BenchmarkCockroachDBUvarintMinLen(b *testing.B) {
	var res uint64
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		res, _ = getUvarint(tests[2].encoded)
	}
	_ = res
}

func BenchmarkDchestPutUint64MinLen(b *testing.B) {
	buf := make([]byte, 9)
	var n int
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		n = varuint.PutUint64(buf, tests[2].decoded)
	}
	_ = n
}

func BenchmarkDchestUint64MinLen(b *testing.B) {
	var res uint64
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		res, _ = varuint.Uint64(tests[2].encoded)
	}
	_ = res
}

func BenchmarkMohaePutUvarintMinLen(b *testing.B) {
	buf := make([]byte, 9)
	var n int
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		n = uvarint.PutUvarint(buf, tests[2].decoded)
	}
	_ = n
}

func BenchmarkMohaeUvarintMinLen(b *testing.B) {
	var res uint64
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		res, _ = uvarint.Uvarint(tests[2].encoded)
	}
	_ = res
}

// Single encode: use > 1<<56 (slowest)
func BenchmarkCockroachDBPutUvarintMaxLen(b *testing.B) {
	buf := make([]byte, 9)
	var n int
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		n = putUvarint(buf, tests[17].decoded)
	}
	_ = n
}

func BenchmarkCockroachDBUvarintMaxLen(b *testing.B) {
	var res uint64
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		res, _ = getUvarint(tests[17].encoded)
	}
	_ = res
}

func BenchmarkDchestPutUint64MaxLen(b *testing.B) {
	buf := make([]byte, 9)
	var n int
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		n = varuint.PutUint64(buf, tests[17].decoded)
	}
	_ = n
}

func BenchmarkDchestUint64MaxLen(b *testing.B) {
	var res uint64
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		res, _ = varuint.Uint64(tests[17].encoded)
	}
	_ = res
}

func BenchmarkMohaePutUvarintMaxLen(b *testing.B) {
	buf := make([]byte, 9)
	var n int
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		n = uvarint.PutUvarint(buf, tests[17].decoded)
	}
	_ = n
}

func BenchmarkMohaeUvarintMaxLen(b *testing.B) {
	var res uint64
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		res, _ = uvarint.Uvarint(tests[17].encoded)
	}
	_ = res
}
