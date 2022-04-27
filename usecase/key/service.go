package key

import "Atlan_Collect_KGS/entity"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetLastKey() (int, error) {
	last, err := s.repo.Get()
	if err != nil {
		return -1, err
	}
	return last, nil
}

func (s *Service) AddKeys(last int, amount int) (error, int) {
	list, newLastkey, err := entity.GetKeyPack(last, amount)
	if err != nil {
		return err, -1
	}
	return s.repo.Create(list, newLastkey), newLastkey
}
