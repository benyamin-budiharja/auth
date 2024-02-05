package validation

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// create unit test for Register function
// Path: service/validation/validation_test.go
func TestValidationService_ValidateLength(t *testing.T) {
	const (
		min = 4
		max = 10
	)

	rules := []Rule{
		ValidateMinLength(min),
		ValidateMaxLength(max),
	}

	t.Run("Input Too Short", func(t *testing.T) {

		validator := NewValidator(rules)

		err := validator.Validate("Full Name", "123")
		assert.NotNil(t, err)
		assert.True(t, Contains("Full Name is less than "+strconv.FormatInt(min, 10), err))
	})

	t.Run("Input Too Long", func(t *testing.T) {
		validator := NewValidator(rules)

		err := validator.Validate("Full Name", "12345678901")
		assert.NotNil(t, err)
		assert.True(t, Contains("Full Name is more than "+strconv.FormatInt(max, 10), err))
	})

	t.Run("Input Length OK", func(t *testing.T) {
		validator := NewValidator(rules)

		err := validator.Validate("Full Name", "123456")
		assert.Nil(t, err)
	})
}

func TestValidationService_ValidatePrefix(t *testing.T) {
	const prefix = "+62"

	rules := []Rule{
		ValidatePrefix(prefix),
	}

	t.Run("Prefix Invalid", func(t *testing.T) {
		validator := NewValidator(rules)

		err := validator.Validate("Phone", "+611111")
		assert.NotNil(t, err)
		assert.True(t, Contains("Phone does not start with "+prefix, err))
	})

	t.Run("Prefix Valid", func(t *testing.T) {
		validator := NewValidator(rules)

		err := validator.Validate("Phone", "+62811231")
		assert.Nil(t, err)
	})
}

func TestValidationService_ValidatePassword(t *testing.T) {
	const (
		capital = 1
		numeric = 1
		special = 1
	)

	rules := []Rule{
		ValidateCapitalPresence(capital),
		ValidateNumericPresence(capital),
		ValidateSpecialCharacterPresence(capital),
	}

	t.Run("Password Only Lowercase", func(t *testing.T) {
		validator := NewValidator(rules)

		err := validator.Validate("Password", "asd")
		assert.NotNil(t, err)
		assert.True(t, Contains("Password required to have "+strconv.FormatInt(capital, 10)+" capital character(s)", err))
		assert.True(t, Contains("Password required to have "+strconv.FormatInt(numeric, 10)+" numeric character(s)", err))
		assert.True(t, Contains("Password required to have "+strconv.FormatInt(special, 10)+" special character(s)", err))
	})

	t.Run("Password Only Uppercase", func(t *testing.T) {
		validator := NewValidator(rules)

		err := validator.Validate("Password", "ASD")
		assert.NotNil(t, err)
		assert.True(t, Contains("Password required to have "+strconv.FormatInt(numeric, 10)+" numeric character(s)", err))
		assert.True(t, Contains("Password required to have "+strconv.FormatInt(special, 10)+" special character(s)", err))
	})

	t.Run("Password With Uppercase and Numeric", func(t *testing.T) {
		validator := NewValidator(rules)

		err := validator.Validate("Password", "ASD123")
		assert.NotNil(t, err)
		assert.True(t, Contains("Password required to have "+strconv.FormatInt(special, 10)+" special character(s)", err))
	})

	t.Run("Password With Uppercase, Numeric, and Special Character", func(t *testing.T) {
		validator := NewValidator(rules)

		err := validator.Validate("Password", "@ASD123")
		assert.Nil(t, err)
	})
}

func Contains(expected string, errs []error) bool {
	for _, e := range errs {
		if e.Error() == expected {
			return true
		}
	}
	return false
}
