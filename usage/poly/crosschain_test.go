package poly

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/polynetwork/poly/common"
	common2 "github.com/polynetwork/poly/native/service/cross_chain_manager/common"
	"testing"
)

func TestParamDecode(t *testing.T) {
	param := &common2.MakeTxParam{}
	data,_ := hex.DecodeString("20000000000000000000000000000000000000000000000000000000000000426620e38a4281ae12d9a2378ffa92cc65a562ed8e749d6d820c22cd5731c52d725e8114250e76987d838a75310c34bf422ea9f1ac4cc906070000000000000014020c15e7d08a8ec7d35bcf3ac3ccbf0bbf2704e606756e6c6f636b4a14c38072aa3f8e049de541223a9c9772132bb48634148b11149cab2d2dfa608aaeae4c6611fc5d0afc34009cec25a899e4c9f4f4ef040000000000000000000000000000000000000000")
	_ = param.Deserialization(common.NewZeroCopySource(data))
	xx, _ := json.Marshal(param)
	fmt.Printf("%s", string(xx))
}
