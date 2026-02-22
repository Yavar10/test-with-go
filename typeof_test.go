package main
import "testing"
func TestTypeOf(t *testing.T){
	got:=typeOf()
	want:="int"
	if got!=want{
		t.Errorf("got %q want %q",got,want)
	}
}