package gateway

import "github.com/CaioAugustoo/wallet-core/internal/entity"

type Client interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
