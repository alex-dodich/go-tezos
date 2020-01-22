package operations

import (
	"github.com/alex-dodich/go-tezos/v2/account"
	"github.com/alex-dodich/go-tezos/v2/delegate"
)

type TezosOperationsService interface {
	CreateBatchPayment(payments []delegate.Payment, wallet account.Wallet, paymentFee int, gaslimit int, batchSize int) ([]string, error)
	InjectOperation(op string) ([]byte, error)
	GetBlockOperationHashes(id interface{}) ([]string, error)
}
