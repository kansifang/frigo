package frigo

/*
#include "friso_API.h"
#include "friso.h"
*/
import "C"
import "unsafe"
import "runtime"

type Frigo struct {
	friso C.friso_t
}

type task struct {
	ft C.friso_task_t
	input *C.char
}

func newTask(text string) *task {
	t := &task{}
	t.ft = C.friso_new_task()
	t.input = C.CString(text)
	C.friso_set_text(t.ft, t.input)
	runtime.SetFinalizer(t, (*task).Free)
	return t
}

func (t *task) Free() {
	if t.ft != nil {
		C.friso_free_task(t.ft)
		t.ft = nil
	}
	if t.input != nil {
		C.free(unsafe.Pointer(t.input))
		t.input = nil
	}
}

func NewFrigo(inifile string) *Frigo {
	f := &Frigo{}
	fname := C.CString(inifile)
	defer C.free(unsafe.Pointer(fname))
	f.friso = C.friso_new_from_ifile(fname);
	runtime.SetFinalizer(f, (*Frigo).Free)
	return f
}

func (f *Frigo) Free() {
	if f.friso != nil {
		C.friso_free(f.friso)
		f.friso = nil
	}
}

func (f *Frigo) Parse(text string) []string {
	task := newTask(text)
	words := make([]string, 0, 1)
	for ; C.friso_next2(f.friso, task.ft) != nil ; {
		word := C.GoString(task.ft.hits.word)
		words = append(words, word)
		C.free(unsafe.Pointer(task.ft.hits.word))
	}
	return words
}