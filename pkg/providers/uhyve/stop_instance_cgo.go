// +build cgo

package uhyve

import (
	"syscall"

	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/emc-advanced-dev/pkg/errors"
)

func (p *UhyveProvider) StopInstance(id string) error {
	instance, err := p.GetInstance(id)
	if err != nil {
		return errors.New("retrieving instance "+id, err)
	}

	// kill qemu
	pid, err := strconv.Atoi(instance.Id)
	if err != nil {
		return errors.New("invalid instance id (should be hermitcore's proxy pid)", err)
	}

	err = syscall.Kill(pid, syscall.SIGINT) //this ensures that the corresponding logfile is immediately deleted

	if err != nil {
		logrus.Warn("failed terminating instance, assuming instance has externally terminated", err)
	}

	return p.state.RemoveInstance(instance)
}
