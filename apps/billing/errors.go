package main

type billingError string

func (be billingError) Error() string {
	return string(be)
}

const (
	ErrInsufficientBalance billingError = "insufficient balance"
)
