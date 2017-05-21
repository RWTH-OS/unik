package uhyve

import (
	"os"
	"path/filepath"

	"github.com/cf-unik/unik/pkg/config"
	"github.com/cf-unik/unik/pkg/state"
)

var debuggerTargetImageName string

type UhyveProvider struct {
	config config.Uhyve
	state  state.State
}

func UhyveStateFile() string {
	return filepath.Join(config.Internal.UnikHome, "uhyve/state.json")

}
func uhyveImagesDirectory() string {
	return filepath.Join(config.Internal.UnikHome, "uhyve/images/")
}

func uhyveInstancesDirectory() string {
	return filepath.Join(config.Internal.UnikHome, "uhyve/instances/")
}

func uhyveVolumesDirectory() string {
	return filepath.Join(config.Internal.UnikHome, "uhyve/volumes/")
}

func NewUhyveProvider(config config.Uhyve) (*UhyveProvider, error) {

	os.MkdirAll(uhyveImagesDirectory(), 0777)
	os.MkdirAll(uhyveInstancesDirectory(), 0777)
	os.MkdirAll(uhyveVolumesDirectory(), 0777)

	p := &UhyveProvider{
		config: config,
		state:  state.NewBasicState(UhyveStateFile()),
	}

	return p, nil
}

func (p *UhyveProvider) WithState(state state.State) *UhyveProvider {
	p.state = state
	return p
}

func getImagePath(imageName string) string {
	return filepath.Join(uhyveImagesDirectory(), imageName, "boot.img")
}

func getKernelPath(imageName string) string {
	return filepath.Join(uhyveImagesDirectory(), imageName, "program.bin")
}

func getCmdlinePath(imageName string) string {
	return filepath.Join(uhyveImagesDirectory(), imageName, "cmdline")
}

func getVolumePath(volumeName string) string {
	return filepath.Join(uhyveVolumesDirectory(), volumeName, "data.img")
}

func getHermitLoaderPath(imageName string) string {
	return filepath.Join(uhyveImagesDirectory(), imageName, "ldhermit.elf")
}
