package server

// testing left

import (
	"context"
	"fmt"

	"github.com/sap200/dvpn-node/utils"
	"github.com/sap200/vineyard/x/vineyard/types"
	"github.com/tendermint/starport/starport/pkg/cosmosclient"
)

func telemeterBandwidth(cc cosmosclient.Client, accountName string) {
	// 1st thing get bandwidth
	bdwth := utils.GetBandwidth()

	// get accountAddress
	accountAddress, err := cc.Address(accountName)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, _ = bdwth, accountAddress

	// get the server info by id
	// that is UUIDOfServer
	qc := types.NewQueryClient(cc.Context)
	qResp, err := qc.NodeAll(context.Background(), &types.QueryAllNodeRequest{})
	if err != nil {
		fmt.Println(err)
		return
	}

	// get my qResp
	// and find my Node too
	var nodeMetaData types.Node
	for _, v := range qResp.Node {
		if v.Index == UUIDOfServer {
			nodeMetaData = v
			break
		}
	}

	// update data
	msg := types.NewMsgUpdateNode(
		accountAddress.String(),
		nodeMetaData.Index,
		nodeMetaData.Address,
		nodeMetaData.Location,
		nodeMetaData.Bandwidth+bdwth,
		nodeMetaData.Uid,
	)

	// broadcast the update command
	txResp, err := cc.BroadcastTx(accountName, msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(txResp)
	}

	// done updating

}
