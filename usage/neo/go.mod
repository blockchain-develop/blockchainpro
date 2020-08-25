module github.com/blockchainpro/usage/neo

go 1.14

require (
	github.com/blockchainpro/usage/utiles v0.0.0
	github.com/joeqian10/neo-gogogo v0.0.0
)

replace (
	github.com/blockchainpro/usage/utiles => github.com/blockchain-develop/blockchainpro/usage/utiles v0.0.0-20200825033423-d1be9e2d6aa1
	github.com/joeqian10/neo-gogogo => github.com/blockchain-develop/neo-gogogo v0.0.0-20200824102609-fddf06a45f66
)
