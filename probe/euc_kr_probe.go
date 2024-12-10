package probe

import (
	"github.com/wlynxg/chardet/cda"
	"github.com/wlynxg/chardet/consts"
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
			NewCodingStateMachine(EucKrSmModel()),
		),
	}
}
