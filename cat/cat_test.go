package cat

import (
	"testing"
)

func TestCatEmptyPathname(t *testing.T) {
	err := Cat("")
	if err == nil {
		t.Fatalf(`Cat("") = %v, want error`, err)
	}
}
