package gateway

import "github.com/CaioAugustoo/wallet-core/internal/entity"

type Account interface {
	Save(account *entity.Account) error
	FindById(id string) (*entity.Account, error)
}
