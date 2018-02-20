package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	// usando esse metodo ele executa esse teste de forma paralela
	t.Parallel()
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquitetura amd64")
	}
	// ..
	t.Fail()
}
