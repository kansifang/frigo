package frigo

import "testing"

func Test1(t *testing.T) {
	f := NewFrigo("./friso.ini")
	words := f.Parse("我爱北京天安门")
	if len(words) != 4 {
		t.Errorf("# of words does not match\n")
	}
}