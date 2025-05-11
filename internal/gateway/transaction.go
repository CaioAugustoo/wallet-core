package gateway

import "github.com/CaioAugustoo/wallet-core/internal/entity"

type Transaction interface {
	Create(transaction *entity.Transaction) error
}
