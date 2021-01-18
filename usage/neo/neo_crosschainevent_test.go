package neo

import (
	"encoding/hex"
	"fmt"
	"github.com/blockchainpro/usage/utiles/common"
	"math/big"
	"strconv"
	"testing"
)

func TestNeoCrossChainEvent_Scan(t *testing.T) {
	client := NewNeoClient()

	for i := 6735756;i < 6735758;i ++ {
		blockResp := client.GetBlockByIndex(uint32(i))
		block := blockResp.Result
		for _, tx := range block.Tx {
			if tx.Type != "InvocationTransaction" {
				continue
			}
			//fmt.Printf("========================================= txid: %s\n", tx.Txid)
			logResp := client.GetApplicationLog(tx.Txid)

			if logResp.ErrorResponse.Error.Message != "" {
				fmt.Printf("GetApplicationLog err: %s\n", logResp.ErrorResponse.Error.Message)
			}

			appLog := logResp.Result
			for _, item := range appLog.Executions {
				//fmt.Printf("saveNeoCrossTxsByHeight height: %d, tx contract: %s, gas: %s\n",i, item.Contract, item.GasConsumed)
				for _, notify := range item.Notifications {
					value := notify.State.Value
					if len(value) <= 0 {
						continue
					}
					method := value[0].Value
					xx, _ := hex.DecodeString(method)
					method = string(xx)
					fmt.Printf("xxxxxxxx: method: %s\n", method, )
					if method == "CrossChainLockEvent" {
						fmt.Printf("Notifications contract: %s, hash: %s\n\n", notify.Contract, tx.Txid)
						fmt.Printf("CrossChainLockEvent: %s, from address: %s, contract address: %s, to chainid: %s, key: %s, param: %s\n",
							value[0].Value, value[1].Value, value[2].Value, value[3].Value, value[4].Value, value[5].Value)
						//parseNotifyData(notify.State.Value[6].Value)
					} else if method == "LockEvent" {
						fmt.Printf("LockEvent: %s, from asset: %s, from address: %s, to chainid: %s, to asset: %s, to address: %s, amount: %s\n",
							value[0].Value, value[1].Value, value[2].Value, value[3].Value, value[4].Value, value[5].Value, value[6].Value)
						amount, _ := strconv.ParseUint(value[6].Value, 16, 32)
						fmt.Printf("========================================= %d\n", amount)
					}
				}
			}
		}
	}
}

func TestNeoCrossChainEvent(t *testing.T) {
	client := NewNeoClient()
	countrep := client.GetBlockCount()
	fmt.Printf("xxxxxxxxxxx %d\n", countrep.Result)
	txhash := "542e6319cf85a8ca662dfdc2dcf515c9ef0d5871a7587a091704e5c7461f015a"
	//txhash := "21c6ae12471611d06682f47863f6771acefc0edffd9a4a5eb3fe8ca2c57c72ef"
	logResp := client.GetApplicationLog(txhash)

	if logResp.ErrorResponse.Error.Message != "" {
		fmt.Printf("GetApplicationLog err: %s\n", logResp.ErrorResponse.Error.Message)
	}

	appLog := logResp.Result
	for _, item := range appLog.Executions {
		fmt.Printf("saveNeoCrossTxsByHeight height: %d tx contract: %s, gas: %s\n", 1, item.Contract, item.GasConsumed)
		for _, notify := range item.Notifications {
			value := notify.State.Value
			method := value[0].Value
			xx, _ := hex.DecodeString(method)
			method = string(xx)
			fmt.Printf("notify, contract: %s, method: %s, values: %v\n", notify.Contract, method, value)
			if method == "CrossChainLockEvent" {
				fmt.Printf("xx: %s, from address: %s, contract address: %s, to chainid: %s, key: %s, param: %s\n",
					value[0].Value, value[1].Value, value[2].Value, value[3].Value, value[4].Value, value[5].Value)
				//parseNotifyData(notify.State.Value[6].Value)
			} else if method == "LockEvent" {
				fmt.Printf("xx: %s, from asset: %s, from address: %s, to chainid: %s, to asset: %s, to address: %s, amount: %s\n",
					value[0].Value, value[1].Value, value[2].Value, value[3].Value, value[4].Value, value[5].Value, value[6].Value)
				amount, _ := strconv.ParseUint(common.HexStringReverse(value[6].Value), 16, 64)
				fmt.Printf("==================================== %d\n", amount)
				if value[6].Type == "Integer" {
					data, _ := strconv.ParseUint((value[6].Value), 10, 64)
					newAmount := new(big.Int).SetInt64(int64(data))
					fmt.Printf("xx==================================== %s\n", newAmount.String())
				}
				{
					newAmount, _ := new(big.Int).SetString(common.HexStringReverse(value[6].Value), 16)
					fmt.Printf("==================================== %s\n", newAmount.String())
				}
			}
		}
	}
}

func TestNeoCrossChainEvent1(t *testing.T) {
	client := NewNeoClient()
	txhash := "0xdd2596b880253db7d5b81487dbd2efdf9829160ce22ea710669a2aaa31ded35a"
	logResp := client.GetApplicationLog(txhash)

	if logResp.ErrorResponse.Error.Message != "" {
		fmt.Printf("GetApplicationLog err: %s\n", logResp.ErrorResponse.Error.Message)
	}

	appLog := logResp.Result
	for _, item := range appLog.Executions {
		fmt.Printf("saveNeoCrossTxsByHeight tx contract: %s, gas: %s\n", item.Contract, item.GasConsumed)
		for _, notify := range item.Notifications {
			value := notify.State.Value
			method := value[0].Value
			xx, _ := hex.DecodeString(method)
			method = string(xx)
			fmt.Printf("notify, contract: %s, method: %s, values: %v\n", notify.Contract, method, value)
			if method == "CrossChainUnlockEvent" {
				fchainId, _ := strconv.ParseUint(notify.State.Value[1].Value, 10, 32)
				fmt.Printf("contract address: %s, from chainid: %d\n",
					common.HexStringReverse(notify.State.Value[2].Value), fchainId)
				//parseNotifyData(notify.State.Value[6].Value)
			} else if method == "UnlockEvent" {
				amount := big.NewInt(0)
				if value[3].Type == "Integer1" {
					//data, _ := strconv.ParseUint(value[3].Value, 10, 64)
					amount, _ = new(big.Int).SetString((value[3].Value), 10)
					//amount = amount.SetInt64(int64(data))
				} else {
					amount, _ = new(big.Int).SetString(common.HexStringReverse(value[3].Value), 16)
				}
				fmt.Printf("token address: %s, amount: %s\n", common.HexStringReverse(value[1].Value), amount.String())
			}
		}
	}
}

func TestCrossChainAmount(t *testing.T)  {
	client := NewNeoClient()
	txhashs := []string{"2c1c1e01884c5978580ba8d9cd351c75edb8fc1f274666b0c820e0fcea3a4392","2c4a4de7677ba339fbbd7d505c527ff54fc294abd5c5dddf72d038950815608f","2c744cfa31aa378862bdbbe696a6435ac88bb4990c594ac5f0fa08a573778765","2c78ff12edb7ca3878845b59087a16e998877066b0eb470c1b5e3adbc21cdf93","2c7a5bea8dafce215c08643de77238e5be19535ddad281dd146ed6e823f109d1","2c08c8522ba13a80cc793d459704277bd31fa2c4320706e05a38e5fd975d8fc4","2c0f76b4f5db759324e01bab004aac37ce0a034377b86dc5adcc864671513f75","2d255655611f3c5ac32ea83acfa9ef620e22cb029c4e427ee9e0bad2eb6395dd","2d3c31eddcbf291d22f493ff8b77a4f8222f8c3662f2ba22ddf5dc2a6bef546c","2c1fd25c9b8bf6638b3740ecfa6ee49dc6d55eb2a3cacb7e89d18df4d2827ef6","454b5a252f9e76f213ab4d5e25f43416d6c4f1f4f9641be114bd16ff9432828a","2c30a6ba80a2308b72739b59d429a3ca5d71201c6516cc8acd82212ee9b74e6b","2c681e5b57b79d058351c6159aab0d16323d1da66f3b8daf0ccc7848a5876a75","2c73726a0c3352fe3706cf8f1c2bd69f3f431527f9ca31b07b3ab116462d7621","2c8b6e3764bc603bb9f43dfccaad709872903581d2fc31dc76d79322a5f22551","2cb62b75984825304cfff12ad79e2802215218923144725c908f39e88b835383","2cbcad1633095b04da7c135befff95ba5c08fc40e3b898ff62b9b6d47d3839e8","2cbe25a22cd23930ea4b8d42a5634b12de2ce8be5ac81d801d2d99c5032ef0e0","2cd819a499c6f121cdff3b09b5ec1cc4775234e44556b25ecf667d04e2063fda","2ce6899d97bbfa90d81252376e0d009880144b3ab6a730746d556aaaceff1c3d","2ceabbbb0faad928e7bcfa90a91fa9eb8c270de11194858a6f2b10fb33413d63","2ceb435fa01c28088ec1f42ac9ef7609cb3cd66542007a82450ec4c62c64118b","2cff205e92e0b253370bb5a730626f3bf6f4eeced01c3fa73969a11fcde4e09b","2d03a7405c3d37a8cd9b7a4b6a48804ea57d5180e367cebfbe8bd64ae5570bd8","2c6e66599a44500a3ee0b2120d93fd540ab6cc9a04c2be0ec560e51aa002c926","2cf5a87a2e0c061f15484a32794c534aadfdc5fc7a333451c0e57d8fc078ecbb","2d2359b683fa3d1ee91c436236a5261a578ec8cd5d55dcf5b38fb9003a394a43","2c1729dc3f8ae56f9afc6a054461a0677939135eaeb8fbea9649f249ebaf6df1","2c9804aab943fb38560afd9ecd944d67270be4a444e38824299396b78dbce603","2c53c18a014513cfacf65020c92453a15634c710cdc9a8e8cde95e152e8e5ac9","2c5fd8b8165622b5b0ec439f8d83fd10533ab34f09a9481d494cec07ee55412f","2c66773aaf3a7cbf3cff09ba39c1123a6f584c1c3bbdfade584f4d51a2f14214","2d4063b69194d084749d19a8bcc6b74911b7e27ec650e05307c9d4f71c864827","96d16bf926017182e92446ddf4e6d1ab7556cbffe742843c350f10f9942ca75f","2c38266f92e2baaae39e8363bf444eaf86d3a78acf4df771ca87e8f6f612ecd9","2c56158e97cd001bee14424772d23fbb9c945ca431e47dd86464317e739cec02","2c68e5b69ba2ec649c7567e1d3c1de175e985b84735a74cdb25f6dd8eb2ec0c1","2cebabe49fd7997891172812289d5567e770dd4c59292755ca05e2a3c9593498","2c051ff8abeceef85eb0ff4dc18b47ec717d736120802b6ca7212a5341afc056","2d116c183b2ae00ed43863cbbee757b5c7906170dec42868d9529156e41963e7"}

	for _, txhash := range txhashs {
		logResp := client.GetApplicationLog(txhash)
		if logResp.ErrorResponse.Error.Message != "" {
			fmt.Printf("GetApplicationLog err: %s\n", logResp.ErrorResponse.Error.Message)
		}
		appLog := logResp.Result
		for _, item := range appLog.Executions {
			for _, notify := range item.Notifications {
				value := notify.State.Value
				method := value[0].Value
				xx, _ := hex.DecodeString(method)
				method = string(xx)
				if method == "UnlockEvent" {
					amount := big.NewInt(0)
					if value[3].Type == "Integer" {
						//data, _ := strconv.ParseUint(value[3].Value, 10, 64)
						amount, _ = new(big.Int).SetString((value[3].Value), 10)
						//amount = amount.SetInt64(int64(data))
					} else {
						amount, _ = new(big.Int).SetString(common.HexStringReverse(value[3].Value), 16)
					}
					fmt.Printf("update tchain_transfer set amount = %s where txhash = \"%s\";\n", amount.String(), txhash)
				}
			}
		}
	}
}
