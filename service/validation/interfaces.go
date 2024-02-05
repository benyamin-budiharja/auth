package validation

type Validator interface {
	Add(rule Rule)
	Validate(key, data string) []error
}
