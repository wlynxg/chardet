package probe

import (
	"github.com/wlynxg/chardet/cda"
	"github.com/wlynxg/chardet/consts"
)

type GB2312Probe struct {
	MultiByteCharSetProbe
}

func NewGB2312Probe() *GB2312Probe {
	return &GB2312Probe{
		MultiByteCharSetProbe: NewMultiByteCharSetProbe(
			consts.GB2312,
			consts.Chinese,
			consts.UnknownLangFilter,
			cda.NewGB2312DistributionAnalysis(),
			NewCodingStateMachine(GB2312SmModel()),
		),
	}
}
