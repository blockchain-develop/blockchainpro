package ethtools

type ECCMLockEvent struct {
	Method   string
	Txid     string
	TxHash   string
	User     string
	Tchain   uint32
	Contract string
	Height   uint64
	Value    []byte
}
type ECCMUnlockEvent struct {
	Method string
	TxHash string
	RTxHash  string
	FChainId uint32
	Contract  string
	Height uint64
}
type LockEvent struct {
	Method          string
	TxHash          string
	FromAddress     string
	FromAssetHash   string
	ToChainId       uint32
	ToAssetHash     string
	ToAddress       string
	Amount          uint64
}
type UnlockEvent struct {
	Method          string
	TxHash          string
	ToAssetHash     string
	ToAddress       string
	Amount          uint64
}



