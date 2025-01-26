package toolkit

import (
	"fmt"
	"testing"
)

func TestTools_RandomString(t *testing.T) {
	var testTools Tools

	s := testTools.RandomString(10)
	if len(s) != 10 {
		t.Errorf("Random string length is not 10")
	}
	fmt.Println("Random string:", s)
}
