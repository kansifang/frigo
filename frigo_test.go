package frigo

import "testing"

func Test1(t *testing.T) {
	f := NewFrigo("/etc/friso/friso.ini")
	task := f.NewTask()
	words := task.Parse("我爱北京天安门")
	if len(words) != 4 {
		t.Errorf("# of words does not match %d != 4\n", len(words))
	}
}

func Test2(t *testing.T) {
	f := NewFrigo("/etc/friso/friso.ini")
	task := f.NewTask()
	words := task.Parse("面包 凯撒大帝 芝士球 照烧小丸子 85冰咖啡 小墨鱼 蓝莓奶酪 枫糖方块 丹麦奶酥菠萝 胚芽奶茶 起士球 抹茶蛋糕 泡芙")
	for _, word := range(words) {
		println(word)
	}
}

func TestMemory(t *testing.T) {
	f := NewFrigo("./friso.ini")
	for i := 0; i < 4000000; i ++ {
		task := f.NewTask()
		task.Parse("我爱北京天安门")
	}
}