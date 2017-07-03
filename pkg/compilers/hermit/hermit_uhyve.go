package hermit

import (
	goerrors "errors"
	"os"
	"path"
	"path/filepath"

	"github.com/Sirupsen/logrus"
	"github.com/cf-unik/unik/pkg/compilers"
	"github.com/cf-unik/unik/pkg/types"
	unikutil "github.com/cf-unik/unik/pkg/util"
	"github.com/emc-advanced-dev/pkg/errors"
)

type HermitUhyveCompiler struct {
	CompilerType string
}

func (i *HermitUhyveCompiler) CompileRawImage(params types.CompileImageParams) (*types.RawImage, error) {
	sourcesDir := params.SourcesDir
	env := make(map[string]string)

	//compilation is done by the container + copy of HermitCore's loader
	if err := unikutil.NewContainer("compilers-hermitcore-common").WithVolume(sourcesDir, "/opt/code").WithEnvs(env).Run(); err != nil {
		return nil, err
	}

	localImageFile, err := i.findFirstImageFile(sourcesDir)
	if err != nil {
		logrus.Errorf("error getting local image file name")
	}

	res := &types.RawImage{}
	res.RunSpec.Compiler = i.CompilerType
	res.LocalImagePath = path.Join(sourcesDir, localImageFile)
	res.StageSpec.ImageFormat = types.ImageFormat_RAW
	res.RunSpec.StorageDriver = types.StorageDriver_IDE
	res.RunSpec.DefaultInstanceMemory = 256
	return res, nil
}

func (i *HermitUhyveCompiler) findFirstImageFile(directory string) (string, error) {
	dir, err := os.Open(directory)
	if err != nil {
		return "", errors.New("could not open dir", err)
	}
	defer dir.Close()
	files, err := dir.Readdir(-1)
	if err != nil {
		return "", errors.New("could not read dir", err)
	}
	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ".img" {
				return file.Name(), nil
			}
		}
	}
	return "", errors.New("no image file found", goerrors.New("end of dir"))
}

func (r *HermitUhyveCompiler) Usage() *compilers.CompilerUsage {
	return nil
}
