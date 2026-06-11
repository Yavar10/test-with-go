package main
import(
	"testing"
	"math"
)

func TestMath(t *testing.T){
	got:=mathTest()
	want:=float32(math.Pi)
	if got!=want{
		t.Errorf("got %f want %f",got,want)
	}
}