package uhyve

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/cf-unik/unik/pkg/config" //add hermitcore proxy
	"github.com/cf-unik/unik/pkg/types"
	"github.com/emc-advanced-dev/pkg/errors"
)

func (p *UhyveProvider) RunInstance(params types.RunInstanceParams) (_ *types.Instance, err error) {
	logrus.WithFields(logrus.Fields{
		"image-id": params.ImageId,
		"mounts":   params.MntPointsToVolumeIds,
		"env":      params.Env,
	}).Infof("running instance %s", params.Name)

	if _, err := p.GetInstance(params.Name); err == nil {
		return nil, errors.New("instance with name "+params.Name+" already exists. uhyve provider requires unique names for instances", nil)
	}

	image, err := p.GetImage(params.ImageId)
	if err != nil {
		return nil, errors.New("getting image", err)
	}

	cmdName := filepath.Join(config.Internal.UnikHome, "hermitcoreproxy")
	cmdArgs := []string{getImagePath(image.Name)}
	cmd := exec.Command(cmdName, cmdArgs...)
	env := os.Environ()
	env = append(env, "HERMIT_ISLE=uhyve")
	cmd.Env = env

	if err := cmd.Start(); err != nil {
		return nil, errors.New("can't start HermitCore's proxy", nil)
	}

	instance := &types.Instance{
		Id:             fmt.Sprintf("%d", cmd.Process.Pid),
		Name:           params.Name,
		State:          types.InstanceState_Running,
		IpAddress:      params.Ip,
		Infrastructure: types.Infrastructure_UHYVE,
		ImageId:        image.Id,
		Created:        time.Now(),
	}

	if err := p.state.ModifyInstances(func(instances map[string]*types.Instance) error {
		instances[instance.Id] = instance
		return nil
	}); err != nil {
		return nil, errors.New("modifying instance map in state", err)
	}

	logrus.WithField("instance", instance).Infof("HermitCore instance created successfully")

	return instance, nil
}
