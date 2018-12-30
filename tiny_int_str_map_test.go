package tinymap

import (
	"testing"
)

func Test_IntStrMap_Get(t *testing.T) {
	intStrMap := new(IntStrMap)

	result, err := intStrMap.Get(42)

	if result != "" {
		t.Errorf("failed to return string default")
	}

	if err == nil {
		t.Errorf("Get without known key should have failed")
	}
}

func Benchmark_IntStrMap_Get_Single_Lower_Bound(b *testing.B) {
	intStrMap := new(IntStrMap)

	intStrMap.Set(0, "_")

	for n := 0; n < b.N; n++ {
		intStrMap.Get(0)
	}
}

func Benchmark_IntStrMap_Get_Max_Size_Upper_Bound(b *testing.B) {
	intStrMap := new(IntStrMap)

	upperBound := 100

	for i := 0; i < upperBound; i++ {
		intStrMap.Set(i, "_")
	}

	intStrMap.Set(upperBound, "_")

	for n := 0; n < b.N; n++ {
		intStrMap.Get(upperBound)
	}
}