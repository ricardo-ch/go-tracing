package main

import (
	"context"
	"errors"
	"strings"
)

type StringService interface {
	Uppercase(context.Context, string) (string, error)
}

type stringService struct{}

func (stringService) Uppercase(_ context.Context, s string) (string, error) {
	if s == "" {
		return "", errors.New("Empty string")
	}
	return strings.ToUpper(s), nil
}
