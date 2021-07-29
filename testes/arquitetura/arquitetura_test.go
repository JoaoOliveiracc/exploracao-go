package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	t.Parallel()
	if runtime.GOARCH == "amd64" {
		t.Skip("N�o funciona em arquitetura amd64")
	}
	// ...
	t.Fail()
}
