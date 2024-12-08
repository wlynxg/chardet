package probe

import (
	"github.com/wlynxg/chardet/cda"
	"github.com/wlynxg/chardet/consts"
	"github.com/wlynxg/chardet/smm"
)

type EUCKRProbe struct {
	MultiByteCharSetProbe
}

func NewEUCKRProbe() *EUCKRProbe {
	return &EUCKRProbe{
		MultiByteCharSetProbe: NewMultiByteCharSetProbe(
			consts.EucKrModelName,
			consts.KoreanLanguage,
			consts.UnknownLangFilter,
			cda.NewEUCKRDistributionAnalysis(),
			smm.NewCodingStateMachine(smm.EucKrSmModel()),
		),
	}
}
