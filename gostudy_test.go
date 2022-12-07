package gostudy

import (
	"testing"
)

func TestVersion(t *testing.T) {
	ver := Version()
	if ver == "" {
		t.Fatalf(`Version() = %v, want not empty`, ver)
	}
}
