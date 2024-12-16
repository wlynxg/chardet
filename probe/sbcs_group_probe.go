package probe

import (
	"github.com/wlynxg/chardet/consts"
)

type SBCSGroupProbe struct {
	CharSetGroupProbe
}

func NewSBCSGroupProbe() *SBCSGroupProbe {
	p := &SBCSGroupProbe{}
	p.CharSetGroupProbe = NewCharSetGroupProbe(consts.UnknownLangFilter, nil)
	hp := NewHebrewProbe()
	model := NewWindows1255HebrewModel()
	logical := NewSingleByteCharSetProbe(model, false, hp)
	visual := NewSingleByteCharSetProbe(model, true, hp)
	hp.SetModelProbe(logical, visual)

	p.probes = []Probe{
		NewSingleByteCharSetProbe(NewWindows1251RussianModel(), false, nil),
		NewSingleByteCharSetProbe(NewKoi8RRussianModel(), false, nil),
		NewSingleByteCharSetProbe(NewISO88595RussianModel(), false, nil),
		NewSingleByteCharSetProbe(NewMacCyrillicRussianModel(), false, nil),
		NewSingleByteCharSetProbe(NewIBM866RussianModel(), false, nil),
		NewSingleByteCharSetProbe(NewIBM855RussianModel(), false, nil),

		NewSingleByteCharSetProbe(NewISO88597GreekModel(), false, nil),
		NewSingleByteCharSetProbe(NewWindows1253GreekModel(), false, nil),

		NewSingleByteCharSetProbe(NewISO88595BulgarianModel(), false, nil),
		NewSingleByteCharSetProbe(NewWindows1251BulgarianModel(), false, nil),

		NewSingleByteCharSetProbe(NewTis620ThaiModel(), false, nil),
		NewSingleByteCharSetProbe(NewIso88599TurkishModel(), false, nil),
		logical,
		visual,
	}
	p.Reset()
	return p
}
