package container

import "github.com/go-kit/kit/log"

type Container struct {

}

func New(
	logger log.Logger,
) *Container {
	return &Container{}
}
