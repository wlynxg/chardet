package probe

import (
	"github.com/wlynxg/chardet/cda"
	"github.com/wlynxg/chardet/consts"
	"github.com/wlynxg/chardet/smm"
)

type Big5Probe struct {
	MultiByteCharSetProbe
}

func NewBig5Probe() *Big5Probe {
	return &Big5Probe{
		MultiByteCharSetProbe: NewMultiByteCharSetProbe(
			consts.Big5ModelName,
			consts.ChineseLanguage,
			consts.UnknownLangFilter,
			cda.NewBig5DistributionAnalysis(),
			smm.NewCodingStateMachine(smm.Big5SmModel()),
		),
	}
}
