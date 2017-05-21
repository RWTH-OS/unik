package uhyve

import (
	"github.com/cf-unik/unik/pkg/providers"
)

func (p *UhyveProvider) GetConfig() providers.ProviderConfig {
	return providers.ProviderConfig{
		UsePartitionTables: true,
	}
}
