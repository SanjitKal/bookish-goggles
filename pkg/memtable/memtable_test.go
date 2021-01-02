package memtable

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestInsert(t *testing.T) {
	memt := Memtable{capacity: 5}
	err := memt.Insert("k1", "v1")
	if err != nil {
		t.Fatalf(err.Error())
	}
	val, err := memt.Lookup("k1")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if val != "v1" {
		t.Fatal(fmt.Sprintf("Expected to find v1, but found %s instead", val))
	}
}

func BenchmarkInsertSequential(b *testing.B) {
	b.StopTimer()
	memt := Memtable{capacity: -1}
	keyArr := make([]string, b.N)
	valArr := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		keyArr[i] = fmt.Sprintf("k%d", i)
		valArr[i] = fmt.Sprintf("v%d", i)
	}
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		err := memt.Insert(keyArr[n], valArr[n])
		if err != nil {
			b.Fatalf(err.Error())
		}
	}
}

func BenchmarkInsertRandom(b *testing.B) {
	b.StopTimer()
	memt := Memtable{capacity: -1}
	keyArr := make([]string, b.N)
	valArr := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		randInt := rand.Intn(b.N)
		keyArr[i] = fmt.Sprintf("k%d", randInt)
		valArr[i] = fmt.Sprintf("v%d", randInt)
	}
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		err := memt.Insert(keyArr[n], valArr[n])
		if err != nil {
			b.Fatalf(err.Error())
		}
	}
}

func BenchmarkLookupAfterManySequentialInsert(b *testing.B) {
	b.StopTimer()
	memt := Memtable{capacity: -1}
	keyArr := make([]string, b.N)
	valArr := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		keyArr[i] = fmt.Sprintf("k%d", i)
		valArr[i] = fmt.Sprintf("v%d", i)
		err := memt.Insert(keyArr[i], valArr[i])
		if err != nil {
			b.Fatalf(err.Error())
		}
	}
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		val, err := memt.Lookup(keyArr[n])
		if val != valArr[n] {
			b.Fatalf(fmt.Sprintf(`Expected to find val "%s" during lookup for key "%s", but found "%s"`, valArr[n], keyArr[n], val))
		}
		if err != nil {
			b.Fatalf(err.Error())
		}
	}
}

func BenchmarkLookupAfterManyRandomInsert(b *testing.B) {
	b.StopTimer()
	memt := Memtable{capacity: -1}
	keyArr := make([]string, b.N)
	valArr := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		randInt := rand.Intn(b.N)
		keyArr[i] = fmt.Sprintf("k%d", randInt)
		valArr[i] = fmt.Sprintf("v%d", randInt)
		err := memt.Insert(keyArr[i], valArr[i])
		if err != nil {
			b.Fatalf(err.Error())
		}
	}

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		val, err := memt.Lookup(keyArr[n])
		if val != valArr[n] {
			b.Fatalf(fmt.Sprintf(`Expected to find val "%s" during lookup for key "%s", but found "%s"`, valArr[n], keyArr[n], val))
		}
		if err != nil {
			b.Fatalf(err.Error())
		}
	}
}
