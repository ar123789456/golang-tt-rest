package email

type EmailUseCase interface {
	ValidEmail(string) ([]string, error)
}
