package key

import "Atlan_Collect_KGS/entity"

type Reader interface {
	Get() (int, error)
}

type Writer interface {
	Create(list []*entity.Key, lastKey int) error
}
type Repository interface {
	Reader
	Writer
}
type UseCase interface {
	GetLastKey() (int, error)
	AddKeys(last int, amount int) (error, int)
}
