package gt

import (
	"github.com/alex-dodich/go-tezos/account"
	"github.com/alex-dodich/go-tezos/block"
	tzc "github.com/alex-dodich/go-tezos/client"
	"github.com/alex-dodich/go-tezos/contracts"
	"github.com/alex-dodich/go-tezos/cycle"
	"github.com/alex-dodich/go-tezos/delegate"
	"github.com/alex-dodich/go-tezos/network"
	"github.com/alex-dodich/go-tezos/node"
	"github.com/alex-dodich/go-tezos/operations"
	"github.com/alex-dodich/go-tezos/snapshot"
	"github.com/pkg/errors"
)

// GoTezos is the driver of the library, it inludes the several RPC services
// like Block, SnapSHot, Cycle, Account, Delegate, Operations, Contract, and Network
type GoTezos struct {
	Client    tzc.TezosClient
	Constants network.Constants
	Block     block.TezosBlockService
	Snapshot  snapshot.TezosSnapshotService
	Cycle     cycle.TezosCycleService
	Account   account.TezosAccountService
	Delegate  delegate.TezosDelegateService
	Network   network.TezosNetworkService
	Operation operations.TezosOperationsService
	Contract  contracts.TezosContractsService
	Node      node.TezosNodeService
}

// NewGoTezos is a constructor that returns a GoTezos object
func NewGoTezos(URL string) (*GoTezos, error) {
	gotezos := GoTezos{}

	gotezos.Client = tzc.NewClient(URL)
	gotezos.Network = network.NewNetworkService(gotezos.Client)
	var err error
	gotezos.Constants, err = gotezos.Network.GetConstants()
	if err != nil {
		return &gotezos, errors.Wrap(err, "could not get network constants")
	}
	gotezos.Block = block.NewBlockService(gotezos.Client)
	gotezos.Cycle = cycle.NewCycleService(gotezos.Block)
	gotezos.Snapshot = snapshot.NewSnapshotService(
		gotezos.Cycle,
		gotezos.Client,
		gotezos.Block,
		gotezos.Constants,
	)
	gotezos.Account = account.NewAccountService(
		gotezos.Client,
		gotezos.Block,
		gotezos.Snapshot,
	)
	gotezos.Delegate = delegate.NewDelegateService(
		gotezos.Client,
		gotezos.Block,
		gotezos.Snapshot,
		gotezos.Account,
		gotezos.Constants,
	)
	gotezos.Operation = operations.NewOperationService(gotezos.Block, gotezos.Client)
	gotezos.Contract = contracts.NewContractService(gotezos.Client)
	gotezos.Node = node.NewNodeService(gotezos.Client)

	return &gotezos, nil
}
