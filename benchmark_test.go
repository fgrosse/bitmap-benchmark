package benchmark

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/RoaringBitmap/roaring"
	"github.com/pilosa/pilosa"
	pilosaroaring "github.com/pilosa/pilosa/roaring"
)

// BENCHMARKS, to run them type "go test -bench Benchmark -run -"

var c9 uint

// go test -bench BenchmarkIntersection -run -
func BenchmarkIntersectionPilosa(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := pilosa.NewBitmap()
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s1.SetBit(uint64(r.Int31n(int32(sz))))
	}
	s2 := pilosa.NewBitmap()
	sz = 100000000
	initsize = 65000
	for i := 0; i < initsize; i++ {
		s2.SetBit(uint64(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	card := uint64(0)
	for j := 0; j < b.N; j++ {
		s3 := s1.Intersect(s2)
		card = card + s3.Count()
	}
}

// go test -bench BenchmarkIntersection -run -
func BenchmarkIntersectionPilosaInternal(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := pilosaroaring.NewBitmap()
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s1.Add(uint64(r.Int31n(int32(sz))))
	}
	s2 := pilosaroaring.NewBitmap()
	sz = 100000000
	initsize = 65000
	for i := 0; i < initsize; i++ {
		s2.Add(uint64(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	card := uint64(0)
	for j := 0; j < b.N; j++ {
		s3 := s1.Intersect(s2)
		card = card + s3.Count()
	}
}

// go test -bench BenchmarkIntersection -run -
func BenchmarkIntersectionRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := roaring.NewBitmap()
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s1.Add(uint32(r.Int31n(int32(sz))))
	}
	s2 := roaring.NewBitmap()
	sz = 100000000
	initsize = 65000
	for i := 0; i < initsize; i++ {
		s2.Add(uint32(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	card := uint64(0)
	for j := 0; j < b.N; j++ {
		s3 := roaring.And(s1, s2)
		card = card + s3.GetCardinality()
	}
}

// go test -bench BenchmarkUnion -run -
func BenchmarkUnionPilosa(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := pilosa.NewBitmap()
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s1.SetBit(uint64(r.Int31n(int32(sz))))
	}
	s2 := pilosa.NewBitmap()
	sz = 100000000
	initsize = 65000
	for i := 0; i < initsize; i++ {
		s2.SetBit(uint64(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	card := uint64(0)
	for j := 0; j < b.N; j++ {
		s3 := s1.Union(s2)
		card = card + s3.Count()
	}
}

// go test -bench BenchmarkUnion -run -
func BenchmarkUnionPilosaInternal(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := pilosaroaring.NewBitmap()
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s1.Add(uint64(r.Int31n(int32(sz))))
	}
	s2 := pilosaroaring.NewBitmap()
	sz = 100000000
	initsize = 65000
	for i := 0; i < initsize; i++ {
		s2.Add(uint64(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	card := uint64(0)
	for j := 0; j < b.N; j++ {
		s3 := s1.Union(s2)
		card = card + s3.Count()
	}
}

// go test -bench BenchmarkUnion -run -
func BenchmarkUnionRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := roaring.NewBitmap()
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s1.Add(uint32(r.Int31n(int32(sz))))
	}
	s2 := roaring.NewBitmap()
	sz = 100000000
	initsize = 65000
	for i := 0; i < initsize; i++ {
		s2.Add(uint32(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	card := uint64(0)
	for j := 0; j < b.N; j++ {
		s3 := roaring.Or(s1, s2)
		card = card + s3.GetCardinality()
	}
}

// go test -bench BenchmarkSize -run -
func BenchmarkSizeRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s1 := roaring.NewBitmap()
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s1.Add(uint32(r.Int31n(int32(sz))))
	}
	s2 := roaring.NewBitmap()
	sz = 100000000
	initsize = 65000
	for i := 0; i < initsize; i++ {
		s2.Add(uint32(r.Int31n(int32(sz))))
	}
	fmt.Printf("%.1f MB ", float32(s1.GetSerializedSizeInBytes()+s2.GetSerializedSizeInBytes())/(1024.0*1024))
}

// TODO get size for pilosa bitmap

// go test -bench BenchmarkSet -run -
func BenchmarkSetRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	sz := 1000000
	s := roaring.NewBitmap()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Add(uint32(r.Int31n(int32(sz))))
	}
}

// go test -bench BenchmarkSet -run -
func BenchmarkSetPilosa(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	sz := 1000000
	s := pilosa.NewBitmap()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.SetBit(uint64(r.Int31n(int32(sz))))
	}
}

// go test -bench BenchmarkSet -run -
func BenchmarkSetPilosaInternal(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	sz := 1000000
	s := pilosaroaring.NewBitmap()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Add(uint64(r.Int31n(int32(sz))))
	}
}

// go test -bench BenchmarkGetTest -run -
func BenchmarkGetTestRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	sz := 1000000
	initsize := 50000
	s := roaring.NewBitmap()
	for i := 0; i < initsize; i++ {
		s.Add(uint32(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Contains(uint32(r.Int31n(int32(sz))))
	}
}

// go test -bench BenchmarkGetTest -run -
func BenchmarkGetTestPilosa(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	sz := 1000000
	initsize := 50000
	s := pilosa.NewBitmap()
	for i := 0; i < initsize; i++ {
		s.SetBit(uint64(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = pilosa.NewBitmap(uint64(r.Int31n(int32(sz)))).IntersectionCount(s) > 0
	}
}

// go test -bench BenchmarkGetTest -run -
func BenchmarkGetTestPilosaInternal(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	sz := 1000000
	initsize := 50000
	s := pilosaroaring.NewBitmap()
	for i := 0; i < initsize; i++ {
		s.Add(uint64(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Contains(uint64(sz))
	}
}

// go test -bench BenchmarkCount -run -
func BenchmarkCountRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s := roaring.NewBitmap()
	sz := 1000000
	initsize := 50000
	for i := 0; i < initsize; i++ {
		s.Add(uint32(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.GetCardinality()
	}
}

// go test -bench BenchmarkCount -run -
func BenchmarkCountPilosa(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s := pilosa.NewBitmap()
	sz := 1000000
	initsize := 50000
	for i := 0; i < initsize; i++ {
		s.SetBit(uint64(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Count()
	}
}

// go test -bench BenchmarkCount -run -
func BenchmarkCountPilosaInternal(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s := pilosaroaring.NewBitmap()
	sz := 1000000
	initsize := 50000
	for i := 0; i < initsize; i++ {
		s.Add(uint64(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Count()
	}
}

// go test -bench BenchmarkIterate -run -
func BenchmarkIterateRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s := roaring.NewBitmap()
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s.Add(uint32(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		c9 = uint(0)
		i := s.Iterator()
		for i.HasNext() {
			i.Next()
			c9++
		}
	}
}

// go test -bench BenchmarkSparseIterate -run -
func BenchmarkSparseIterateRoaring(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s := roaring.NewBitmap()
	sz := 100000000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s.Add(uint32(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		c9 = uint(0)
		i := s.Iterator()
		for i.HasNext() {
			i.Next()
			c9++
		}
	}
}

// go test -bench BenchmarkIterate -run -
func BenchmarkIteratePilosaInternal(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s := pilosaroaring.NewBitmap()
	sz := 150000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s.Add(uint64(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		c9 = uint(0)
		i := s.Iterator()
		var eof bool
		for !eof {
			_, eof = i.Next()
			c9++
		}
	}
}

// go test -bench BenchmarkSparseIterate -run -
func BenchmarkSparseIteratePilosaInternal(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0))
	s := pilosaroaring.NewBitmap()
	sz := 100000000
	initsize := 65000
	for i := 0; i < initsize; i++ {
		s.Add(uint64(r.Int31n(int32(sz))))
	}
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		c9 = uint(0)
		i := s.Iterator()
		var eof bool
		for !eof {
			_, eof = i.Next()
			c9++
		}
	}

}
