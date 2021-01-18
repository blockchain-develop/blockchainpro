package ontology

import (
	"fmt"
	"github.com/ontio/ontology/common"
	"testing"
)

func TestOntologyBalance1(t *testing.T) {
	// create alliance sdk
	ontSdk := newOntologySdk()
	account_user, _ := common.AddressFromHexString("f05db35f619a991435b160b21e4d9371fdc0d9f2")
	//
	balance := getOntologyOngBalance(ontSdk, account_user)
	fmt.Printf("amount of address %s is: %d\n", account_user.ToBase58(), balance)
}

func TestCreateAccount(t *testing.T) {
	createOntologyAccount()
}
