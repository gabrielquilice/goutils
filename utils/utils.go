package utils

import (
	"log"
	"reflect"
)

// All valid types
type AnyValue interface {
	any | *any
}

// Simulates the behavior of a ternary if
func TernaryIf[T AnyValue](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	} else {
		return falseValue
	}
}

// Checks whether the variable was instantiated or not
func IsZero[T AnyValue](v T) bool {
	return reflect.ValueOf(&v).Elem().IsZero()
}

// Returns the first value if it's not null, otherwise it returns the second
func GetValueOrElse[T AnyValue](value, replacement T) T {
	if IsZero(value) {
		return replacement
	} else {
		return value
	}
}

// Display formatted string in console
func PrintString(s string, a ...any) {
	log.Printf(s+"\n", a)
}

// Returns nil if the values are equal, otherwise returns a pointer with the result of "value"
func NullIf[T comparable](value, compareValue T) *T {
	if value == compareValue {
		return nil
	}

	return &compareValue
}

