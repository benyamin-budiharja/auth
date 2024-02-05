package validation

import (
	"fmt"
	"unicode"
)

type Validator struct {
	rule Rules
}

func (v *Validator) Add(rule Rule) {
	v.rule = append(v.rule, rule)
}

func (v *Validator) Validate(key, data string) []error {
	var errors []error
	for _, rule := range v.rule {
		err := rule(key, data)
		if err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}

func ValidateMinLength(min int) Rule {
	return func(key, data string) error {
		if len(data) < min {
			return fmt.Errorf("%s is less than %d", key, min)
		}
		return nil
	}
}

func ValidateMaxLength(max int) Rule {
	return func(key, data string) error {
		if len(data) > max {
			return fmt.Errorf("%s is more than %d", key, max)
		}
		return nil
	}
}

func ValidatePrefix(prefix string) Rule {
	return func(key, data string) error {
		if data[:len(prefix)] != prefix {
			return fmt.Errorf("%s does not start with %s", key, prefix)
		}
		return nil
	}
}

func ValidateCapitalPresence(required int) Rule {
	return func(key, data string) error {
		count := 0
		for _, a := range data {
			if unicode.IsUpper(a) {
				count++
				if count >= required {
					return nil
				}
			}
		}
		return fmt.Errorf("%s required to have %d capital character(s)", key, required)
	}
}

func ValidateNumericPresence(required int) Rule {
	return func(key, data string) error {
		count := 0
		for _, a := range data {
			if unicode.IsNumber(a) {
				count++
				if count >= required {
					return nil
				}
			}
		}
		return fmt.Errorf("%s required to have %d numeric character(s)", key, required)
	}
}

func ValidateSpecialCharacterPresence(required int) Rule {
	return func(key, data string) error {
		count := 0
		for _, a := range data {
			if !unicode.IsLetter(a) && !unicode.IsNumber(a) {
				count++
				if count >= required {
					return nil
				}
			}
		}
		return fmt.Errorf("%s required to have %d special character(s)", key, required)
	}

}
