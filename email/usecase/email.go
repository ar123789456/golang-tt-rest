package usecase

import "regexp"

type EmailUseCase struct {
}

func NewEmailUseCase() *EmailUseCase {
	return &EmailUseCase{}
}

func (em *EmailUseCase) ValidEmail(s string) ([]string, error) {
	emailRegexp, err := regexp.Compile("[a-zA-Z0-9]{1,}@[a-zA-Z0-9]{1,}.[a-z]{1,}")
	if err != nil {
		return []string{}, err
	}
	return emailRegexp.FindAllString(s, -1), err
}
