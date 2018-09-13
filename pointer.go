package packr

import (
	"io/ioutil"

	"github.com/gobuffalo/packr/file"
	"github.com/gobuffalo/packr/file/resolver"
	"github.com/pkg/errors"
)

type Pointer struct {
	ForwardBox  string
	ForwardPath string
}

var _ resolver.Resolver = Pointer{}

func (p Pointer) Find(box string, path string) (file.File, error) {
	b := findBox(p.ForwardBox)
	f, err := b.Resolve(p.ForwardPath)
	if err != nil {
		return f, errors.WithStack(err)
	}
	x, err := ioutil.ReadAll(f)
	if err != nil {
		return f, errors.WithStack(err)
	}
	return file.NewFile(path, x), nil
}