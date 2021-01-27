package neo

import (
	"fmt"
	"testing"
)

func TestNep5_One(t *testing.T) {
	nep5Hash := "658cabf9c1f71ba0fa64098a7c17e52b94046ece"
	nep5 := NewNep5(nep5Hash)
	decimal, err := nep5.Decimals()
	if err != nil {
		panic(err)
	}
	name, err := nep5.Name()
	if err != nil {
		panic(err)
	}
	fmt.Printf("nep5: %s, decimal: %d, name: %s\n", nep5Hash, decimal, name)
}

func TestNep5(t *testing.T) {
	tokens := []string{"b951ecbbc5fe37a9c280a76cb0ce0014827294cf","ab38352559b8b203bde5fddfa0b07d8b2525e132","17c76859c11bc14da5b3e9c88fa695513442c606","271e1e4616158c7440ffd1d5ca51c0c12c792833","0df563008be710f3e0130208f8adc95ed7e5518d","f46719e2d16bf50cddcef9d4bbfece901f73cbb6","c277117879af3197fbef92c71e95800aa3b89d9a","282e3340d5a1cd6a461d5f558d91bc1dbc02a07b","534dcac35b0dfadc7b2d716a7a73a7067c148b37"}
	for _, token := range tokens {
		nep5Hash := token
		nep5 := NewNep5(nep5Hash)
		decimal, err := nep5.Decimals()
		if err != nil {
			panic(err)
		}
		name, err := nep5.Name()
		if err != nil {
			panic(err)
		}

		fmt.Printf("nep5: %s, decimal: %d, name: %s\n", nep5Hash, decimal, name)
	}
}
