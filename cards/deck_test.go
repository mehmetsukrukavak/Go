package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("52'lik deste oluşturulması beklenirken %v'lik deste oluştu", len(d))
	}

	if d[0] != "Karo As" {
		t.Errorf("İlk kart Karo As olmalı fakat %v", d[0])
	}

	if d[len(d)-1] != "Maça Papaz" {
		t.Errorf("Son kart Maça Papaz olmalı fakat %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("52'lik deste oluşturulması beklenirken %v'lik deste oluştu", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
