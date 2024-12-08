package probe

import (
	"github.com/wlynxg/chardet/cda"
	"github.com/wlynxg/chardet/consts"
	"github.com/wlynxg/chardet/smm"
)

type CP949Probe struct {
	MultiByteCharSetProbe
}

func NewCP949Probe() *CP949Probe {
	return &CP949Probe{
		MultiByteCharSetProbe: NewMultiByteCharSetProbe(
			consts.CP949ModelName,
			consts.KoreanLanguage,
			consts.UnknownLangFilter,
			cda.NewEUCKRDistributionAnalysis(),
			smm.NewCodingStateMachine(smm.CP949SmModel()),
		),
	}
}
