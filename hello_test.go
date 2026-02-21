package hello
import (
	"testing"
)

func TestAdd(t *testing.T){
	got:=Hello()
	want:="Hello Zuri"
	if got!=want{
		t.Errorf("got %q want %q",got,want)
	}
}