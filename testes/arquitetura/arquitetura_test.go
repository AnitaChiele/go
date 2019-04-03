package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	// indica que o teste pode ser executado de forma paralela.
	t.Parallel()

	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquitetura amd64")
	}

	t.Fail()
}
