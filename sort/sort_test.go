package sort

import (
	"math/rand"
	"testing"
)

/**
@IDE        :   GoLand
@Author     :   bytebody
@Date       :   2020/11/17 9:05 PM
@Description:   测试 sort
*/

func BenchmarkInts(b *testing.B) {
	array := make([][]int, 0, 10000)
	for i := 0; i < cap(array); i++ {
		values := make([]int, 0, 200)
		for j := 0; j < 2000; j++ {
			values = append(values, rand.Intn(cap(array)))
		}
		array = append(array, values)
	}
	b.Run("stable", func(b *testing.B) {

		for i := 0; i < b.N; i++ {
			values := array[i%len(array)]
			Stable(IntSlice(values))
		}
	})
	array = make([][]int, 0, 10000)
	for i := 0; i < cap(array); i++ {
		values := make([]int, 0, 200)
		for j := 0; j < 2000; j++ {
			values = append(values, rand.Intn(cap(array)))
		}
		array = append(array, values)
	}
	b.Run("quick", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			values := array[i%len(array)]
			Sort(IntSlice(values))
		}
	})
}

func TestInts(t *testing.T) {
	array := make([][]int, 0, 1000)
	for i := 0; i < cap(array); i++ {
		values := make([]int, 0, 64)
		for j := 0; j < cap(values); j++ {
			values = append(values, rand.Intn(cap(array)))
		}
		array = append(array, values)
	}
	t.Run("q", func(t *testing.T) {
		Ints(array[0])
	})
}
