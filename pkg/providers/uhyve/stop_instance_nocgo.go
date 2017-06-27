// +build !cgo

package uhyve

import "github.com/emc-advanced-dev/pkg/errors"

func (p *UhyveProvider) StopInstance(id string) error {

	return errors.New("Stopping qemu instance is not supported without cgo", nil)
}
