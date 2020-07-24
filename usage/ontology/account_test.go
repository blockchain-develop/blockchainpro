package ontology

import (
	"fmt"
	"github.com/ontio/ontology/common"
	"testing"
)

func TestOntologyBalance1(t *testing.T) {
	// create alliance sdk
	ontSdk := newOntologySdk()
	account_user, _ := common.AddressFromBase58(LAYER2_CONTRACT)
	//
	balance := getOntologyOngBalance(ontSdk, account_user)
	fmt.Printf("amount of address %s is: %d\n", account_user.ToBase58(), balance)
}
