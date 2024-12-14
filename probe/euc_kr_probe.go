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
			consts.EucKr,
			consts.Korean,
			consts.UnknownLangFilter,
			cda.NewEUCKRDistributionAnalysis(),
			NewCodingStateMachine(EucKrSmModel()),
		),
	}
}
