package probe

import (
	"github.com/wlynxg/chardet/consts"
)

func MBCGroupProbe(filter consts.LangFilter) *CharSetGroupProbe {
	return &CharSetGroupProbe{
		filter: filter,
		probes: []ICharSetProbe{
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
	}
}
