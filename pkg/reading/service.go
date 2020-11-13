package reading

//  Repository .
type Repository interface {
	GetAllCandyNames() ([]string, error)
}

// Service .
type Service interface {
	GetAllCandyNames() ([]string, error)
}
type service struct {
	r Repository
}

// NewService .
func NewService(r Repository) *service {
	return &service{r}
}

// GetAllCandyNames .
func (s *service) GetAllCandyNames() ([]string, error) {
	cs, err := s.r.GetAllCandyNames()
	if err != nil {
		return nil, err
	}
	return cs, nil
}