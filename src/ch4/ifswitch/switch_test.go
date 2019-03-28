package ifswitch1_test

import "testing"

func TestIfSwitch(t *testing.T)  {
	s0 := [] int {1,2,3,4,5,6};

	for _,val := range s0{
		switch {
		case val%2==0:
			t.Log("偶数");
		case val%2 == 1:
			t.Log("奇数")
		default:
			t.Log("其他")
		}
	}

	i := 0

	for i<5 {
		t.Log(i)
		i++
	}
}