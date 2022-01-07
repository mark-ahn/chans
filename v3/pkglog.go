package chans

import (
	"io"
	"reflect"

	"github.com/rs/zerolog/log"
)

type _anchor struct{}

var pkglog = log.With().
	Str("package", reflect.TypeOf(_anchor{}).PkgPath()).
	Logger()

func SetPkgLogWriter(w io.Writer) {
	pkglog = pkglog.Output(w)
}
