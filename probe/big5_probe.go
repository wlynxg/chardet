package probe

import (
	"github.com/wlynxg/chardet/cda"
	"github.com/wlynxg/chardet/consts"
)

type Big5Probe struct {
	MultiByteCharSetProbe
}

func NewBig5Probe() *Big5Probe {
	return &Big5Probe{
		MultiByteCharSetProbe: NewMultiByteCharSetProbe(
			consts.Big5,
			consts.Chinese,
			consts.UnknownLangFilter,
			cda.NewBig5DistributionAnalysis(),
			NewCodingStateMachine(Big5SmModel()),
		),
	}
}
