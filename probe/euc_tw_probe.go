package probe

import (
	"github.com/wlynxg/chardet/cda"
	"github.com/wlynxg/chardet/consts"
)

type EUCTWProbe struct {
	MultiByteCharSetProbe
}

func NewEUCTWProbe() *EUCTWProbe {
	return &EUCTWProbe{
		MultiByteCharSetProbe: NewMultiByteCharSetProbe(
			consts.EucTwModelName,
			consts.ChineseLanguage,
			consts.UnknownLangFilter,
			cda.NewEUCTWDistributionAnalysis(),
			NewCodingStateMachine(EucTwSmModel()),
		),
	}
}
