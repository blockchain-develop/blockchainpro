package ethereum

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/shopspring/decimal"
    "golang.org/x/crypto/sha3"
    "io/ioutil"
    "math/big"
    "strconv"
    "strings"
    "testing"
)

type Hold struct {
    Balance decimal.Decimal
    Address string
    Symbol string
    Contract string
    Decimal uint
    CorrectBalance decimal.Decimal
}

func MustDecimal(value string) decimal.Decimal {
    a, err := decimal.NewFromString(value)
    if err != nil {
        panic(err)
    }
    return a
}

func MustUint(value string) uint {
    a, err := strconv.ParseUint(value, 10, 64)
    if err != nil {
        panic(err)
    }
    return uint(a)
}

func TestEthereumBalance(t *testing.T) {
    //
    data, err := ioutil.ReadFile("./pg_0314.csv")
    if err != nil {
        panic(err)
    }

    //
    holds := make([]*Hold, 0)
    records := bytes.Split(data, []byte{'\n'})
    for _, record := range records {
        items := bytes.Split(record, []byte{','})
        if len(items) != 5 {
            panic("invalid records")
        }
        balance := MustDecimal(string(items[0]))
        dec := MustUint(string(items[4]))
        decX := decimal.New(1, int32(dec))
        balanceX := balance.Mul(decX)
        contract :=  string(items[3])
        if contract == "\"\"" {
            contract = ""
        }
        holds = append(holds, &Hold{
            Balance:  balanceX,
            Address:  string(items[1]),
            Symbol:   string(items[2]),
            Contract: contract,
            Decimal:  MustUint(string(items[4])),
        })
    }

    //
    clients := make([]*ethclient.Client, 0)
    clientx, err := ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/oZjqYK7lYFYbEJetOg5XduI2PgS5SN0g")
    if err != nil {
        panic(err)
    }
    clients = append(clients, clientx)
    clientx, err = ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/Z4JRVYhdGm_X2p69UM8o3L456Scz4Olr")
    if err != nil {
        panic(err)
    }
    clients = append(clients, clientx)
    clientx, err = ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/VxcJR8wb-vUv34QIvIumAvIxGB7rNQhY")
    if err != nil {
        panic(err)
    }
    clients = append(clients, clientx)

    //
    slice := 9
    exits := make(chan bool, 0)
    for j := 0;j < slice;j ++ {
        go func(index int) {
            for i, hold := range holds {
                if i % slice != index {
                    continue
                }
                client := clients[index % 3]
                if hold.Contract == "" { // ETH
                    balance, err := client.BalanceAt(context.Background(), common.HexToAddress(hold.Address), nil)
                    if err != nil {
                        hold.CorrectBalance = decimal.NewFromInt(-1)
                        continue
                    }
                    hold.CorrectBalance = decimal.NewFromBigInt(balance, 0)
                } else {
                    funcSignature := []byte("balanceOf(address)")
                    hash := sha3.NewLegacyKeccak256()
                    hash.Write(funcSignature)
                    methodId := hash.Sum(nil)[:4]

                    userAddress := common.HexToAddress(hold.Address)
                    paddedUserAddress := common.LeftPadBytes(userAddress.Bytes(), 32)

                    var callData []byte
                    callData = append(callData, methodId...)
                    callData = append(callData, paddedUserAddress...)

                    contractAddress := common.HexToAddress(hold.Contract)
                    res, err := client.CallContract(context.Background(), ethereum.CallMsg{
                        From:       common.Address{},
                        To:         &contractAddress,
                        Gas:        0,
                        GasPrice:   big.NewInt(0),
                        GasFeeCap:  big.NewInt(0),
                        GasTipCap:  big.NewInt(0),
                        Value:      big.NewInt(0),
                        Data:       callData,
                        AccessList: nil,
                    }, nil)
                    if err != nil {
                        hold.CorrectBalance = decimal.NewFromInt(-1)
                        continue
                    }
                    balance := new(big.Int).SetBytes(res)
                    hold.CorrectBalance = decimal.NewFromBigInt(balance, 0)
                }
                fmt.Printf("%d\n", i)
            }
            //
            exits <- true
        }(j)
    }

    for j := 0;j < slice;j ++ {
        <- exits
    }

    holdsJson, _ := json.MarshalIndent(holds, "", "    ")
    ioutil.WriteFile("./holds.json", holdsJson, 0666)
}

func TestEthereumBalance1(t *testing.T) {
    //
    data, err := ioutil.ReadFile("./holds.json")
    if err != nil {
        panic(err)
    }

    //
    holds := make([]*Hold, 0)
    json.Unmarshal(data, &holds)

    //
    holdsX := make([]*Hold, 0)
    for _, hold := range holds {
        if hold.CorrectBalance.IsNegative() {
            continue
        }
        if hold.CorrectBalance.Cmp(hold.Balance) == 0 {
            continue
        }
        holdsX = append(holdsX, hold)
    }

    fmt.Printf("%d\n", len(holdsX))

    holdsXJson, _ := json.MarshalIndent(holdsX, "", "    ")
    ioutil.WriteFile("./holdsX.json", holdsXJson, 0666)
}

func TestEthereumBalance2(t *testing.T) {
    //
    data, err := ioutil.ReadFile("./holds.json")
    if err != nil {
        panic(err)
    }

    //
    holds := make([]*Hold, 0)
    json.Unmarshal(data, &holds)

    //
    content := make([]string, 0)
    for _, hold := range holds {
        if hold.CorrectBalance.Cmp(hold.Balance) == 0 {
            continue
        }
        //
        curl := fmt.Sprintf(`curl --location --request POST 'http://localhost:8083/api/v1/tokencollect/balance_correct' \
--header 'Content-Type: application/json' \
--data-raw '{
    "chain": "ethereum",
    "address":"%s",
    "coin":"%s",
    "wallet":"Phemex_old"
}'

sleep 0.5

`, hold.Address, hold.Symbol)
        content = append(content, curl)
    }

    contentx := strings.Join(content, "")
    ioutil.WriteFile("./content.json", []byte(contentx), 0666)
}