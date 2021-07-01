package bloom

import (
	"fmt"
	"testing"
)

func TestBloom(t *testing.T) {
	f := New(10000, 0.01)
	f.Add("Test")
	b1 := f.Exist("Test")
	fmt.Println(b1)
	b2 := f.Exist("Test2")
	t.Log(b2)
}
