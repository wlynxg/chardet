package chardet

type CharSetGroupProbe struct {
	filter LangFilter

	probes []ICharSetProbe
}

func (c *CharSetGroupProbe) Reset() {

}

func MBCGroupProbe(filter LangFilter) *CharSetGroupProbe {
	return &CharSetGroupProbe{
		filter: filter,
		probes: []ICharSetProbe{
			NewUTF8Probe(),
		},
	}
}
