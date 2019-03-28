package str

import "testing"

func TestString (t *testing.T) {
	var s1 string

	s1="中"

	t.Log(len(s1))
	t.Logf("utf8 %x", s1)

	c1:=[] rune(s1)

	t.Log(len(c1))
	t.Logf("unicode %x", c1[0])
	t.Logf("utf8 %x", s1)

}
