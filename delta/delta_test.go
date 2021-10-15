package delta

import (
	"testing"

	"github.com/iimrudy/vUtils/random"
)

func TestHelloName(t *testing.T) {
	s1 := random.GenerateString(90000)
	st := random.GenerateString(90000)
	s2 := s1 + st + s1

	dt := cDelta([]byte(s1), []byte(s2))
	st = ""
	t.Log(dt.SourceSize())
	t.Log(dt.TargetSize())
	x, err := dt.Apply([]byte(s1))
	if err != nil {
		t.Fatal(err)
		t.Fail()
	} else {
		if string(x) == s2 {
			t.Log("Test OK")
		} else {
			t.Fatal("X != s2 (patched string different from original)")
			t.Fail()
		}
	}
}
