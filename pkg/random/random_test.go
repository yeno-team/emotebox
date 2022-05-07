package random_test

import (
	"testing"

	"github.com/yeno-team/emotebox/pkg/random"
)

func TestRandString(t *testing.T) {
	for i := 0; i < 20; i += 5 {
		r, err := random.RandString(i + 1)
		if err != nil {
			t.Fatal(err)
		}

		if len(r) != i {
			t.Fatalf("Length of generated string '%s' expected to be '%d'", r, i)
		}
	}
}
