package probe

import (
	"github.com/wlynxg/chardet/consts"
)

type HebrewProbe struct {
	CharSetProbe

	Space, NormalKaf, FinalKaf, FinalMem, NormalMem, FinalNun, NormalNun,
	FinalPe, NormalPe, FinalTsadi, NormalTsadi byte
	MinFinalCharDistance                int
	MinModelDistance                    float64
	VisualHebrewName, LogicalHebrewName string

	finalCharLogicalScore int
	finalCharVisualScore  int
	prev                  byte
	beforePrev            byte
	logicalProbe          Probe
	visualProbe           Probe
}

func NewHebrewProbe() *HebrewProbe {
	p := &HebrewProbe{
		CharSetProbe: NewCharSetProbe(consts.UnknownLangFilter),
		Space:        0x20,
		// windows-1255 / ISO-8859-8 code points of interest
		FinalKaf:    0xEA,
		NormalKaf:   0xEB,
		FinalMem:    0xED,
		NormalMem:   0xEE,
		FinalNun:    0xEF,
		NormalNun:   0xF0,
		FinalPe:     0xF3,
		NormalPe:    0xF4,
		FinalTsadi:  0xF5,
		NormalTsadi: 0xF6,
		// Minimum Visual vs Logical final letter score difference.
		// If the difference is below this, don't rely solely on the final letter score
		// distance.
		MinFinalCharDistance: 5,
		// Minimum Visual vs Logical model score difference.
		// If the difference is below this, don't rely at all on the model score
		// distance.
		MinModelDistance:      0.01,
		VisualHebrewName:      consts.ISO88598,
		LogicalHebrewName:     consts.Windows1255,
		finalCharLogicalScore: 0,
		finalCharVisualScore:  0,
		logicalProbe:          nil,
		visualProbe:           nil,
	}
	p.prev = p.Space
	p.beforePrev = p.Space
	return p
}

func (h *HebrewProbe) Reset() {
	h.finalCharLogicalScore = 0
	h.finalCharVisualScore = 0
	// The two last characters seen in the previous buffer,
	// mPrev and mBeforePrev are initialized to space in order to simulate
	// a word delimiter at the beginning of the data
	h.prev = h.Space
	h.beforePrev = h.Space
	// These probers are owned by the group prober.
}

func (h *HebrewProbe) SetModelProbe(logical, visual Probe) {
	h.logicalProbe = logical
	h.visualProbe = visual
}

func (h *HebrewProbe) IsFinal(c byte) bool {
	switch c {
	case h.FinalKaf, h.FinalMem, h.FinalNun, h.FinalPe, h.FinalTsadi:
		return true
	default:
		return false
	}
}

func (h *HebrewProbe) IsNonFinal(c byte) bool {
	// The normal Tsadi is not a good Non-Final letter due to words like
	// 'lechotet' (to chat) containing an apostrophe after the tsadi. This
	// apostrophe is converted to a space in FilterWithoutEnglishLetters
	// causing the Non-Final tsadi to appear at an end of a word even
	// though this is not the case in the original text.
	// The letters Pe and Kaf rarely display a related behavior of not being
	// a good Non-Final letter. Words like 'Pop', 'Winamp' and 'Mubarak'
	// for example legally end with a Non-Final Pe or Kaf. However, the
	// benefit of these letters as Non-Final letters outweighs the damage
	// since these words are quite rare.
	switch c {
	case h.NormalKaf, h.NormalMem, h.NormalNun, h.NormalPe:
		return true
	default:
		return false
	}
}

func (h *HebrewProbe) Feed(buf []byte) consts.ProbingState {
	// Final letter analysis for logical-visual decision.
	// Look for evidence that the received buffer is either logical Hebrew
	// or visual Hebrew.
	// The following cases are checked:
	// 1) A word longer than 1 letter, ending with a final letter. This is
	//    an indication that the text is laid out "naturally" since the
	//    final letter really appears at the end. +1 for logical score.
	// 2) A word longer than 1 letter, ending with a Non-Final letter. In
	//    normal Hebrew, words ending with Kaf, Mem, Nun, Pe or Tsadi,
	//    should not end with the Non-Final form of that letter. Exceptions
	//    to this rule are mentioned above in isNonFinal(). This is an
	//    indication that the text is laid out backwards. +1 for visual
	//    score
	// 3) A word longer than 1 letter, starting with a final letter. Final
	//    letters should not appear at the beginning of a word. This is an
	//    indication that the text is laid out backwards. +1 for visual
	//    score.
	//
	// The visual score and logical score are accumulated throughout the
	// text and are finally checked against each other in GetCharSetName().
	// No checking for final letters in the middle of words is done since
	// that case is not an indication for either Logical or Visual text.
	//
	// We automatically filter out all 7-bit characters (replace them with
	// spaces) so the word boundary detection works properly. [MAP]

	if h.state == consts.NotMeProbingState {
		// Both model probers say it's not them. No reason to continue.
		return consts.NotMeProbingState
	}

	buf = h.FilterHighByteOnly(buf)
	for _, b := range buf {
		if b == h.Space {
			// We stand on a space - a word just ended
			if h.beforePrev != h.Space {
				// next-to-last char was not a space so self._prev is not a 1-letter word
				if h.IsFinal(h.prev) {
					// case (1) [-2:not space][-1:final letter][cur:space]
					h.finalCharLogicalScore++
				} else if h.IsNonFinal(h.prev) {
					// case (2) [-2:not space][-1:Non-Final letter][cur:space]
					h.finalCharVisualScore++
				}
			}
		} else {
			if h.beforePrev == h.Space && h.IsFinal(h.prev) && b != h.Space {
				// case (3) [-2:space][-1:final letter][cur:not space]
				h.finalCharVisualScore++
			}
		}
		h.beforePrev = h.prev
		h.prev = b
	}
	// Forever detecting, till the end or until both model probes return
	// ProbingState.NOT_ME (handled above)
	return consts.DetectingProbingState
}

func (h *HebrewProbe) CharSetName() string {
	// Make the decision: is it Logical or Visual?
	// If the final letter score distance is dominant enough, rely on it.
	finalSub := h.finalCharLogicalScore - h.finalCharVisualScore
	if finalSub >= h.MinFinalCharDistance {
		return h.LogicalHebrewName
	} else if finalSub <= -h.MinFinalCharDistance {
		return h.VisualHebrewName
	}

	// It's not dominant enough, try to rely on the model scores instead.
	modelSub := h.logicalProbe.GetConfidence() - h.visualProbe.GetConfidence()
	if modelSub > h.MinModelDistance {
		return h.LogicalHebrewName
	} else if modelSub < -h.MinModelDistance {
		return h.VisualHebrewName
	}

	// Still no good, back to final letter distance, maybe it'll save the day.
	if finalSub < 0 {
		return h.VisualHebrewName
	}

	// (finalSub > 0 - Logical) or (don't know what to do) default to Logical.
	return h.LogicalHebrewName
}

func (h *HebrewProbe) Language() string {
	return consts.Hebrew
}

func (h *HebrewProbe) State() consts.ProbingState {
	// TODO: _logical_prober
	return consts.DetectingProbingState
}

func (h *HebrewProbe) GetConfidence() float64 {
	return 0
}
