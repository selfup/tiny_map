package tinymap

import (
	"testing"
)

func Test_TinyStrMap_Get(t *testing.T) {
	tinyStrMap := new(TinyStrMap)

	result, err := tinyStrMap.Get("a")

	if result != "" {
		t.Errorf("failed to return string default")
	}

	if err == nil {
		t.Errorf("Get without known key should have failed")
	}
}

func Benchmark_TinyStrMap_Get_Single_Lower_Bound(b *testing.B) {
	tinyStrMap := new(TinyStrMap)

	tinyStrMap.Set("a", "_")

	for n := 0; n < b.N; n++ {
		tinyStrMap.Get("a")
	}
}

func Benchmark_TinyStrMap_Get_Max_Size_Upper_Bound(b *testing.B) {
	tinyStrMap := new(TinyStrMap)

	for i := 0; i < 100; i++ {
		tinyStrMap.Set(string(i), "_")
	}

	tinyStrMap.Set("a", "_")

	for n := 0; n < b.N; n++ {
		tinyStrMap.Get("a")
	}
}
