package ch4

import "testing"

func TestSlice2 (t *testing.T){

	//var slice []int
	//slice = []int{1,2,3}
	//slice1 := [3]int{1,2,3}

	//var slice []int;
	slice := make([] int, 2, 10)

	for i:=0;i<10;i++ {
		slice = append(slice, i)
		//t.Log(slice, len(slice), cap(slice))
	}

	s0 := slice[3:6];

	t.Log(s0, len(s0), cap(s0))
}
