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

type Task struct {
	ft C.friso_task_t
	f C.friso_t
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
	println("frigo freed")
	if f.friso != nil {
		C.friso_free(f.friso)
		f.friso = nil
	}
}


func (f *Frigo) NewTask() *Task {
	t := &Task{f: f.friso}
	t.ft = C.friso_new_task()
	runtime.SetFinalizer(t, (*Task).Free)
	return t
}

func (t *Task) Free() {
	println("task freed")
	if t.ft != nil {
		C.friso_free_task(t.ft)
		t.ft = nil
	}
	t.f = nil
}

func (t *Task) Parse(text string) []string {
	input := C.CString(text)
	defer C.free(unsafe.Pointer(input))
	C.friso_set_text(t.ft, input)
	words := make([]string, 0, 1)
	for ; C.friso_next(t.f, t.f.mode, t.ft) != nil ; {
		word := C.GoString(t.ft.hits.word)
		words = append(words, word)
		if t.ft.hits.wtype == C.__FRISO_NEW_WORDS__ {
			C.free(unsafe.Pointer(t.ft.hits.word))
		}
	}

	return words
}