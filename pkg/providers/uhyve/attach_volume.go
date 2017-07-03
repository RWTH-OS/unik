package uhyve

import "github.com/emc-advanced-dev/pkg/errors"

func (p *UhyveProvider) AttachVolume(id, instanceId, mntPoint string) error {
	return errors.New("not yet supportded for uhyve", nil)
}
