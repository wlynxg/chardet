package probe

import (
	"github.com/wlynxg/chardet/cda"
	"github.com/wlynxg/chardet/consts"
	"github.com/wlynxg/chardet/smm"
)

type JOHABProbe struct {
	MultiByteCharSetProbe
}

func NewJOHABProbe() *JOHABProbe {
	return &JOHABProbe{
		MultiByteCharSetProbe: NewMultiByteCharSetProbe(
			consts.JohabName,
			consts.KoreanLanguage,
			consts.UnknownLangFilter,
			cda.NewJOHABDistributionAnalysis(),
			smm.NewCodingStateMachine(smm.JohabSmModel()),
		),
	}
}
