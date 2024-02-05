package validation

type Rule func(key, value string) error

type Rules []Rule
