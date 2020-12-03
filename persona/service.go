package persona

/*Service interface para los sercicios*/
type Service interface {
}

type service struct {
	repo Repository
}

/*NerService permite crear el sercicio*/
func NerService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
