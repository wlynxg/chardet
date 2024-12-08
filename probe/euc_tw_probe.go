package probe

import (
	"github.com/wlynxg/chardet/cda"
	"github.com/wlynxg/chardet/consts"
	"github.com/wlynxg/chardet/smm"
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
			smm.NewCodingStateMachine(smm.EucTwSmModel()),
		),
	}
}
