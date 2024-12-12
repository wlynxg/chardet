package probe

import (
	"github.com/wlynxg/chardet/cda"
	"github.com/wlynxg/chardet/consts"
)

type JOHABProbe struct {
	MultiByteCharSetProbe
}

func NewJOHABProbe() *JOHABProbe {
	return &JOHABProbe{
		MultiByteCharSetProbe: NewMultiByteCharSetProbe(
			consts.Johab,
			consts.Korean,
			consts.UnknownLangFilter,
			cda.NewJOHABDistributionAnalysis(),
			NewCodingStateMachine(JohabSmModel()),
		),
	}
}
