package main
import "testing"
func TestSwap(t *testing.T){
	a,b:="Hello","Zuri"
	got1,got2:=swap(a,b)
	want1,want2:="Zuri","Hello"
	if got1!=want1&&got2!=want2{
		t.Errorf("got1 %q got2 %q want1 %q want2 %q",got1,got2,want1,want2)
	}
}