module github.com/blockchainpro/usage

go 1.14

require (
	github.com/cosmos/cosmos-sdk v0.0.0
	github.com/joeqian10/neo-gogogo v0.0.0-20200814072357-20f86305fd8b
	github.com/ontio/ontology v1.11.1-0.20200817111815-fb3d61fdc12e
	github.com/ontio/ontology-crypto v1.0.9
	github.com/ontio/ontology-go-sdk v0.0.0
	github.com/tendermint/iavl v0.14.0
	github.com/tendermint/tendermint v0.33.7
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/Switcheo/cosmos-sdk v0.39.2-0.20200814061308-474a0dbbe4ba
	github.com/ontio/ontology-go-sdk => github.com/blockchain-develop/ontology-go-sdk v1.11.9-0.20200819032651-0a1c83f57ac6
)
