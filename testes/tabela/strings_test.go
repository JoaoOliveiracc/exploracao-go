package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) - �ndices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Explorando Go", "Go", 0},
		{"", "", 0},
		{"Opa", "Opa", -1},
		{"Jo�o", "o", 2},
	}

	for _, teste := range testes {
		t.Logf("Massa: %v", teste)
		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, atual)
		}
	}
}
