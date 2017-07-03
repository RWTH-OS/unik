package uhyve

import (
	"github.com/cf-unik/unik/pkg/types"
	"github.com/emc-advanced-dev/pkg/errors"
)

func (p *UhyveProvider) CreateVolume(params types.CreateVolumeParams) (_ *types.Volume, err error) {
	return nil, errors.New("creating directory for volume file", nil)
}
