package fabric

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"testing"
)

func newFabSdk() *fabsdk.FabricSDK {
	sdk, err := fabsdk.New(config.FromFile("./config_e2e.yaml"))
	if err != nil {
		panic(err)
	}
	return sdk
}

func newResMgmt(sdk *fabsdk.FabricSDK) *resmgmt.Client {
	rcp := sdk.Context(fabsdk.WithUser("Admin"), fabsdk.WithOrg("Org1"))
	rc, err := resmgmt.New(rcp)
	if err != nil {
		panic(err)
	}
	return rc
}

func newChannelClient(sdk *fabsdk.FabricSDK) *channel.Client {
	ccp := sdk.ChannelContext("mychannel", fabsdk.WithUser("User1"))
	cc, err := channel.New(ccp)
	if err != nil {
		panic(err)
	}
	return cc
}

func packArgs(args []string) [][]byte {
	ret := make([][]byte, 0)
	for _, arg := range args {
		ret = append(ret, []byte(arg))
	}
	return ret
}

func TestCCQuery(t *testing.T) {
	sdk := newFabSdk()
	channelClient := newChannelClient(sdk)
	req := channel.Request{
		ChaincodeID: "basic",
		Fcn: "query",
		Args: packArgs([]string{"GetAllAssets"}),
	}
	response, err := channelClient.Query(req, channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		panic(err)
	}
	fmt.Printf("response: %v\n", response)
}

func TestCCInvoke(t *testing.T) {
	sdk := newFabSdk()
	channelClient := newChannelClient(sdk)
	req := channel.Request{
		ChaincodeID: "basic",
		Fcn: "TransferAsset",
		Args: packArgs([]string{"asset6","Christopher"}),
	}
	response, err := channelClient.Execute(req, channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		panic(err)
	}
	fmt.Printf("response: %v\n", response)
}
