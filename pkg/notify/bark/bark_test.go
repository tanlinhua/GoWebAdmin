package bark

import (
	"testing"
)

func TestSend(t *testing.T) {
	o, c := Send("title", "content")
	t.Log(o, c)
}
