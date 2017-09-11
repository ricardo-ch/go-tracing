package main

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/ricardo-ch/go-tracing"
)

type StringService interface {
	Uppercase(context.Context, string) (string, error)
}

type stringService struct{}

func (stringService) Uppercase(ctx context.Context, s string) (string, error) {
	if s == "" {
		return "", errors.New("Empty string")
	}
	time.Sleep(1 * time.Second)
	nestedFunc(ctx)

	return strings.ToUpper(s), nil
}

func nestedFunc(ctx context.Context) {
	span, ctx := tracing.CreateSpan(ctx, "nestedFunc", nil)
	defer span.Finish()

	time.Sleep(1 * time.Second)
}
