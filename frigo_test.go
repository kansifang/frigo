package frigo

import "testing"

func Test1(t *testing.T) {
	f := NewFrigo("./friso.ini")
	task := f.NewTask()
	words := task.Parse("我爱北京天安门")
	if len(words) != 4 {
		t.Errorf("# of words does not match\n")
	}
}

func TestMemory(t *testing.T) {
	f := NewFrigo("./friso.ini")
	for i := 0; i < 4000000; i ++ {
		task := f.NewTask()
		task.Parse("我爱北京天安门")
	}
}