// This file contains the repository implementation layer.
package validation

type ValidatorImpl struct {
	rules Rules
}

type UserRepository struct {
}

type NewRepositoryOptions struct {
	Dsn string
}

func NewValidator(rule []Rule) *ValidatorImpl {
	return &ValidatorImpl{
		rules: rule,
	}
}
