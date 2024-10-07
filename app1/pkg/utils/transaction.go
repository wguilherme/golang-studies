package utils

import (
	"context"

	"github.com/google/uuid"
)

type ContextTransactionType string

var ContextTransactionKey = ContextTransactionType("transaction")

func ContextWithTransaction(ctx context.Context) context.Context {
	return context.WithValue(ctx, ContextTransactionKey, uuid.New().String())
}

func GetTransaction(ctx context.Context) string {
	return ctx.Value(ContextTransactionKey).(string)
}
