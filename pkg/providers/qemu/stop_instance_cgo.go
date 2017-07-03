// +build cgo

package qemu

import (
	"syscall"

	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/cf-unik/unik/pkg/compilers"
	"github.com/cf-unik/unik/pkg/types"
	"github.com/emc-advanced-dev/pkg/errors"
)

func (p *QemuProvider) StopInstance(id string) error {
	instance, err := p.GetInstance(id)
	if err != nil {
		return errors.New("retrieving instance "+id, err)
	}

	// kill qemu
	pid, err := strconv.Atoi(instance.Id)
	if err != nil {
		return errors.New("invalid instance id (should be qemu pid)", err)
	}

	image, err := p.GetImage(instance.ImageId)

	if compilers.CompilerType(image.RunSpec.Compiler).Base() == "hermitcore" { // kill proxy with SIGTERM so that the proxy quits qemu
		err = syscall.Kill(pid, syscall.SIGTERM)
	} else {
		err = syscall.Kill(pid, syscall.SIGKILL)
	}
	if err != nil {
		logrus.Warn("failed terminating instance, assuming instance has externally terminated", err)
	}

	volumesToDetach := []*types.Volume{}
	volumes, err := p.ListVolumes()
	if err != nil {
		return errors.New("getting volume list", err)
	}
	for _, volume := range volumes {
		if volume.Attachment == instance.Id {
			volumesToDetach = append(volumesToDetach, volume)
		}
	}

	return p.state.RemoveInstance(instance)
}
