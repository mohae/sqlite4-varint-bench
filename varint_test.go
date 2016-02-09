package sqlite4_varint_bench

import (
	"bytes"
	"testing"

	"github.com/dchest/varuint"
	"github.com/mohae/uvarint"
)

func TestUvarint(t *testing.T) {
	for i, test := range tests {
		b := make([]byte, len(test.encoded))
		n := uvarint.PutUvarint(b, test.decoded)
		if n != test.n {
			t.Errorf("encode %d: got %d want %d", i, n, test.n)
		}
		if !bytes.Equal(b, test.encoded) {
			t.Errorf("encode %d: got %v want %v", i, b[0:n], test.encoded)
		}
		v, n := uvarint.Uvarint(test.encoded)
		if n != test.n {
			t.Errorf("decode %d: got %d want %d", i, n, test.n)
		}
		if v != test.decoded {
			t.Errorf("decode %d: got %d want %d", i, v, test.decoded)
		}
	}
}

func TestVaruint(t *testing.T) {
	for i, test := range tests {
		b := make([]byte, len(test.encoded))
		n := varuint.PutUint64(b, test.decoded)
		if n != test.n {
			t.Errorf("encode %d: got %d want %d", i, n, test.n)
		}
		if !bytes.Equal(b, test.encoded) {
			t.Errorf("encode %d: got %v want %v", i, b[0:n], test.encoded)
		}
		v, n := varuint.Uint64(test.encoded)
		if n != test.n {
			t.Errorf("decode %d: got %d want %d", i, n, test.n)
		}
		if v != test.decoded {
			t.Errorf("decode %d: got %d want %d", i, v, test.decoded)
		}
	}
}

func TestCoackroach(t *testing.T) {
	for i, test := range tests {
		b := make([]byte, len(test.encoded))
		n := putUvarint(b, test.decoded)
		if n != test.n {
			t.Errorf("encode %d: got %d want %d", i, n, test.n)
		}
		if !bytes.Equal(b, test.encoded) {
			t.Errorf("encode %d: got %v want %v", i, b[0:n], test.encoded)
		}
		v, n := getUvarint(test.encoded)
		if n != test.n {
			t.Errorf("decode %d: got %d want %d", i, n, test.n)
		}
		if v != test.decoded {
			t.Errorf("decode %d: got %d want %d", i, v, test.decoded)
		}
	}
}
