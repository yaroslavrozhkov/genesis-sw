package repository

type RateRepo struct {
}

func NewRateRepo() *RateRepo {
	return &RateRepo{}
}

func (r *RateRepo) RateCurrent() (int, error) {
	mess, err := getCurrentExchange()
	return mess, err
}
