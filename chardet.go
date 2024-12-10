package chardet

import (
	"github.com/wlynxg/chardet/consts"
)

func Detect(buf []byte) Result {
	d := NewUniversalDetector(consts.UnknownLangFilter)
	d.Feed(buf)
	return d.GetResult()
}
