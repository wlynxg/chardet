package probe

import (
	"github.com/wlynxg/chardet/consts"
)

type MBCSGroupProbe struct {
	CharSetGroupProbe
}

func MBCGroupProbe(filter consts.LangFilter) *MBCSGroupProbe {
	return &MBCSGroupProbe{
		CharSetGroupProbe: NewCharSetGroupProbe(
			filter,
			[]ICharSetProbe{
				NewUTF8Probe(),
				NewSJISProbe(),
				NewEUCJPProbe(),
				NewGB2312Probe(),
				NewEUCKRProbe(),
				NewCP949Probe(),
				NewBig5Probe(),
				NewEUCTWProbe(),
				NewJOHABProbe(),
			},
		),
	}
}
