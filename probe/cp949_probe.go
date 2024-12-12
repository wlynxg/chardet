package probe

import (
	"github.com/wlynxg/chardet/cda"
	"github.com/wlynxg/chardet/consts"
)

type CP949Probe struct {
	MultiByteCharSetProbe
}

func NewCP949Probe() *CP949Probe {
	return &CP949Probe{
		MultiByteCharSetProbe: NewMultiByteCharSetProbe(
			consts.CP949,
			consts.Korean,
			consts.UnknownLangFilter,
			cda.NewEUCKRDistributionAnalysis(),
			NewCodingStateMachine(CP949SmModel()),
		),
	}
}
