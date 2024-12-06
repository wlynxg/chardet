package probe

import (
	"github.com/wlynxg/chardet/consts"
)

func MBCGroupProbe(filter consts.LangFilter) CharSetGroupProbe {
	return CharSetGroupProbe{
		filter: filter,
		probes: []ICharSetProbe{
			NewUTF8Probe(),
		},
	}
}
