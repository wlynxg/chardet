package cda

type EUCJPContextAnalysis struct {
	JapaneseContextAnalysis
}

func NewEUCJPContextAnalysis() *EUCJPContextAnalysis {
	e := &EUCJPContextAnalysis{}
	e.JapaneseContextAnalysis = NewJapaneseContextAnalysis(e.GetOrder)
	return e
}

func (e *EUCJPContextAnalysis) CharSetName() string {
	return ""
}

func (e *EUCJPContextAnalysis) GetOrder(buf []byte) (int, int) {
	if len(buf) == 0 {
		return -1, -1
	}

	var (
		firstChar = buf[0]
		charLen   = 1
	)

	if firstChar == 0x8E || (firstChar >= 0xA1 && firstChar <= 0xFE) {
		charLen = 2
	} else if firstChar == 0x8F {
		charLen = 3
	}
	if len(buf) > 1 {
		secondChar := buf[1]
		if firstChar == 0xA4 && (secondChar >= 0xA1 && secondChar <= 0xF3) {
			return int(secondChar) - 0xA1, charLen
		}
	}
	return -1, charLen
}
