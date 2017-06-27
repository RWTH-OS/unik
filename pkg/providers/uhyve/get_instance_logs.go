package uhyve

import (
	"github.com/cf-unik/unik/pkg/providers/common"
	"github.com/emc-advanced-dev/pkg/errors"
)

func (p *UhyveProvider) GetInstanceLogs(id string) (string, error) {
	instance, err := p.GetInstance(id)
	if err != nil {
		return "", errors.New("retrieving instance "+id, err)
	}
	return common.GetInstanceLogs(instance)
}
