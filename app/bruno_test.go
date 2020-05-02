package main

import "testing"

// Nessa função de teste, no t.Errorf usei o primeiro %s que faz referencia ao html; e o segundo %s faz referencia ao texto esperado
func TestGreeting(t *testing.T) {
	html := greeting("Bruno Leal - Code.education Rocks!")
	if html != "<b>Bruno Leal - Code.education Rocks!</b>" {
		t.Errorf("O texto está incorreto, valor recebido: %s, valor desejado: %s.", html, "<b>Bruno Leal - Code.education Rocks!</b>")
	}
}
