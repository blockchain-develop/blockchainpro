package ontology

import (
	"fmt"
	"testing"
)

func TestOep4(t *testing.T) {
	tokens := []string{"8037dd7161401417d3571b92b86846d34309129a","a2f89e531e55636d4af1cd044237d2fd5a616c72","3e931f60f2cd1387b52f1889dfcaf02a54b2c6a0","061a07cd393aac289b8ecfda2c3784b637a2fb33","3f0def1945d7129c5f6625147dcbbaaee402e751","33ae7eae016193ba0fe238b223623bc78faac158","46c3051c553aaeb3724ea69336ec483f39fa91b1","00c59fcd27a562d6397883eab1f2fff56e58ef80","17a58a4a65959c2f567e5063c560f9d09fb81284","0dabee6055a1c17e3b4bcb15af1a713605b7fcfc","9a576d927dda934b8ce69f35ec2c1025ceb10e6f","7b956c0c11fcffb9c9227ca1925ba4c3486b36f1","df19600d334bb13c6a9e3e9777aa8ec6ed6a4a79","ac654837a90eee8fccabd87a2d4fc7637484f01a"}
	sdk := newOntologySdk()
	for _, token := range tokens {
		contract := token
		name, err := OEP4Name(sdk, contract)
		if err != nil {
			panic(err)
		}
		symbol, err := OEP4Symbol(sdk, contract)
		if err != nil {
			panic(err)
		}
		decimal, err := OEP4Decimals(sdk, contract)
		if err != nil {
			panic(err)
		}
		supply, err := OEP4Supply(sdk, contract)
		if err != nil {
			panic(err)
		}
		fmt.Printf("oep4: %s, name: %s, symbol: %s, decimal: %d, supply: %s\n", contract, name, symbol, decimal, supply)
	}
}

func TestOep4_One(t *testing.T) {
	sdk := newOntologySdk()
	contract := "8037dd7161401417d3571b92b86846d34309129a"
	name, err := OEP4Name(sdk, contract)
	if err != nil {
		panic(err)
	}
	symbol, err := OEP4Symbol(sdk, contract)
	if err != nil {
		panic(err)
	}
	decimal, err := OEP4Decimals(sdk, contract)
	if err != nil {
		panic(err)
	}
	supply, err := OEP4Supply(sdk, contract)
	if err != nil {
		panic(err)
	}
	fmt.Printf("oep4: %s, name: %s, symbol: %s, decimal: %d, supply: %s\n", contract, name, symbol, decimal, supply)
}
