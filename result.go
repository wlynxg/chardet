package chardet

import "github.com/wlynxg/chardet/consts"

func newResult(name string, confidence float64, language string) Result {
	charset := consts.CanonicalCharset(name)
	legacy := consts.LegacyCharset(charset)
	return Result{
		Encoding:   legacy,
		Charset:    charset,
		Confidence: confidence,
		Language:   language,
	}
}
