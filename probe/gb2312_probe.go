package probe

import (
	"github.com/wlynxg/chardet/cda"
	"github.com/wlynxg/chardet/consts"
	"github.com/wlynxg/chardet/smm"
)

type GB2312Probe struct {
	MultiByteCharSetProbe
}

func NewGB2312Probe() *GB2312Probe {
	return &GB2312Probe{
		MultiByteCharSetProbe: NewMultiByteCharSetProbe(
			consts.GB2312ModelName,
			consts.ChineseLanguage,
			consts.UnknownLangFilter,
			cda.NewGB2312DistributionAnalysis(),
			smm.NewCodingStateMachine(smm.GB2312SmModel()),
		),
	}
}
