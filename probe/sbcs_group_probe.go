package probe

type SBCSGroupProbe struct {
	CharSetGroupProbe
}

func NewSBCSGroupProbe() *SBCSGroupProbe {
	p := &SBCSGroupProbe{}
	return p
}
