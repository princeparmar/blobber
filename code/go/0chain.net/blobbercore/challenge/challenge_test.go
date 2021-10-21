package challenge

import (
	"context"
	"testing"

	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/datastore"
	"github.com/0chain/blobber/code/go/0chain.net/core/chain"
	"github.com/0chain/blobber/code/go/0chain.net/core/node"
	"github.com/0chain/gosdk/zcncore"
)

func TestSyncOpenChallenge(t *testing.T) {

	chain.SetServerChain(&chain.Chain{
		ID: "0afc093ffb509f059c55478bc1a60351cef7b4e9c008a53a6cc8241ca8617dfe",
	})

	node.Self.ID = "f3d819ef590dab90adfa37462796a974cb219b9ddc4b18f8f33733688c9e8b2b"
	node.Self.PublicKey = "2da9ac8e1c0a56af38064e300e8f4c5724d28cfde219ec7bb249f4bc4ad40c1ce45012ae287cb16577d6d604f3c8d7ed3af7cb04ca3d533f7adb5ac745de8611"

	zcncore.SetNetwork([]string{
		"https://huawei.0chain.net/miner01",
		"https://huawei.0chain.net/miner02",
		"https://huawei.0chain.net/miner03",
	}, []string{
		"https://huawei.0chain.net/sharder02",
		"https://huawei.0chain.net/sharder01",
		"https://127.0.0.1/sharder",
	})

	datastore.UseMocket(true)

	acceptChallenges(context.TODO())

}
