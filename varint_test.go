package sqlite4_bench

import (
	"bytes"
	"testing"

	"github.com/dchest/varuint"
	"github.com/mohae/varint"
)

func TestVarInt(t *testing.T) {
	for i, test := range tests {
		b := make([]byte, len(test.encoded))
		n := varint.PutUint64(b, test.decoded)
		if n != test.n {
			t.Errorf("encode %d: got %d want %d", i, n, test.n)
		}
		if !bytes.Equal(b, test.encoded) {
			t.Errorf("encode %d: got %v want %v", i, b[0:n], test.encoded)
		}
		v, n := varint.Uint64(test.encoded)
		if n != test.n {
			t.Errorf("decode %d: got %d want %d", i, n, test.n)
		}
		if v != test.decoded {
			t.Errorf("decode %d: got %d want %d", i, v, test.decoded)
		}
	}
}

func TestVarUint(t *testing.T) {
	for i, test := range tests {
		b := make([]byte, len(test.encoded))
		n := varuint.PutUint64(b, test.decoded)
		if n != test.n {
			t.Errorf("encode %d: got %d want %d", i, n, test.n)
		}
		if !bytes.Equal(b, test.encoded) {
			t.Errorf("encode %d: got %v want %v", i, b[0:n], test.encoded)
		}
		v, n := varint.Uint64(test.encoded)
		if n != test.n {
			t.Errorf("decode %d: got %d want %d", i, n, test.n)
		}
		if v != test.decoded {
			t.Errorf("decode %d: got %d want %d", i, v, test.decoded)
		}
	}
}
