package uhyve

import "github.com/emc-advanced-dev/pkg/errors"

func (p *UhyveProvider) DeleteVolume(id string, force bool) error {
	return errors.New("not yet supportded for uhyve", nil)
}
