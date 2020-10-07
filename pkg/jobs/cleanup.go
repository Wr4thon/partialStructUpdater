package jobs

import (
	"github.com/docker/docker/pkg/integration/cmd"
	"github.com/pkg/errors"
)

type cleanup struct {
	file string
}

func NewCleanupJob(file string) Job {
	return &cleanup{
		file: file,
	}
}

func (c *cleanup) Run() error {
	res := cmd.RunCommand("go", "fmt", c.file)
	if res.Error != nil {
		return errors.Wrap(res.Error, "error while executing command")
	}

	return nil
}
