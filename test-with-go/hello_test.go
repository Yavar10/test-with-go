package main
import (
	"testing"
)

func TestAdd(t *testing.T){
	got:=Hello()
	want:="Hello Zuri"
	got2:=Hello2("Yavar")
	want2:="Hello Yavar"
	if got!=want&&got2!=want2{
		t.Errorf("got %q want %q got2 %q want2 %q",got,want,got2,want2)
	}
}
