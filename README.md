# sqlite4-bench
Benchmarks for sqlite4 variable length integer encoding of 64-bit unsigned integer implementations in Go.

See https://sqlite.org/src4/doc/trunk/www/varint.wiki for details of the encoding and decoding rules.


## Benchmarks
Three benchmarks are run against the Encode and Decode operations for each implementation.  

### All
The `*All` benchmarks, use a range of values that represent the min and max values, along with values that are at the encoding/decoding rule boundaries and a couple extra cases.

```
BenchmarkCockroachDBPutUvarintAll   	 5000000	       255 ns/op	  31.27 MB/s	       0 B/op	       0 allocs/op
BenchmarkCockroachDBUvarintAll      	 5000000	       282 ns/op	  28.32 MB/s	       0 B/op	       0 allocs/op
BenchmarkDchestPutUint64All         	10000000	       156 ns/op	  51.18 MB/s	       0 B/op	       0 allocs/op
BenchmarkDchestUint64All            	10000000	       159 ns/op	  50.27 MB/s	       0 B/op	       0 allocs/op
BenchmarkMohaePutUvarintAll         	10000000	       154 ns/op	  51.69 MB/s	       0 B/op	       0 allocs/op
BenchmarkMohaeUvarintAll            	10000000	       167 ns/op	  47.66 MB/s	       0 B/op	       0 allocs/op
```

### MinLen
The `*MinLen` benchmarks use a value that encodes to 1 byte.  This shows the best-case scenario for encoding/decoding operations.  This is usually the most common case.

```
BenchmarkCockroachDBPutUvarintMinLen	500000000	         3.61 ns/op	2215.80 MB/s	       0 B/op	       0 allocs/op
BenchmarkCockroachDBUvarintMinLen   	300000000	         4.17 ns/op	1918.88 MB/s	       0 B/op	       0 allocs/op
BenchmarkDchestPutUint64MinLen      	500000000	         3.33 ns/op	2399.68 MB/s	       0 B/op	       0 allocs/op
BenchmarkDchestUint64MinLen         	500000000	         3.52 ns/op	2270.33 MB/s	       0 B/op	       0 allocs/op
BenchmarkMohaePutUvarintMinLen      	500000000	         3.31 ns/op	2415.08 MB/s	       0 B/op	       0 allocs/op
BenchmarkMohaeUvarintMinLen         	500000000	         3.56 ns/op	2246.06 MB/s	       0 B/op	       0 allocs/op
```

### MaxLen
The `*MaxLen` benchmarks use a value that encodes to 9 bytes.  This shows the worst-case scenario for encoding/decoding operations.  This is usually the least common case.

```
BenchmarkCockroachDBPutUvarintMaxLen	100000000	        22.4 ns/op	 357.25 MB/s	       0 B/op	       0 allocs/op
BenchmarkCockroachDBUvarintMaxLen   	100000000	        22.8 ns/op	 350.79 MB/s	       0 B/op	       0 allocs/op
BenchmarkDchestPutUint64MaxLen      	100000000	        10.9 ns/op	 732.82 MB/s	       0 B/op	       0 allocs/op
BenchmarkDchestUint64MaxLen         	100000000	        11.4 ns/op	 704.60 MB/s	       0 B/op	       0 allocs/op
BenchmarkMohaePutUvarintMaxLen      	200000000	         9.56 ns/op	 836.79 MB/s	       0 B/op	       0 allocs/op
BenchmarkMohaeUvarintMaxLen         	100000000	        10.6 ns/op	 757.42 MB/s	       0 B/op	       0 allocs/op
```

## Notes:
The results of [varuint](https://github.com/dchest/varuint) and [uvarint](https://github.com/mohae/uvarint) should be about the same.  The main difference between the two is that Varuint does boundary checks on the slice while Uvarint will panic if the slice isn't of sufficient size.  Uvarint behaves this way because a slice of insufficient length is considered a programmer error.

The [CockroachDB](https://github.com/cockroachdb/cockroach) implementation of SQLite4's varint is copied from https://github.com/cockroachdb/cockroach/blob/77296d6b6c8c7a6c0d26456067a033a0d33bb730/util/encoding/varint.go and can be found in the `cockroachdb_varint.go` file.  This is necessary because the funcs are not exported.  The CockroachDB implementation is copyrighted by The Cockroach Authors under the Apache License, Version 2.0 (the "License").  A copy of the license may be obtained at

    http://www.apache.org/licenses/LICENSE-2.0

A complete copy of the original copyright notification may be found in the `cockroachdb_varint.go` file.
