package utils

import (
	"app1/pkg/utils"
	"context"
	"errors"
	"os"
)

type EnvironmentHelper interface {
	Check(ctx context.Context)
}

func NewEnvironmentHelper() EnvironmentHelper {
	return &environmentHelper{}
}

type environmentHelper struct {
}

func (eh environmentHelper) Check(ctx context.Context) {

	envs := []string{
		"PORT",
	}

	for _, env := range envs {
		if os.Getenv(env) == "" {
			err := errors.New("environment " + env + " must to be defined")
			utils.GetLogger().Error(ctx, err).Send()
			panic(err)
		}
	}

}
