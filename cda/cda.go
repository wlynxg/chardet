package cda

type CharDistributionAnalysis struct {
	EnoughDataThreshold  int
	SureYes              float64
	SureNo               float64
	MinimumDataThreshold int

	charToFreqOrder          []int
	tableSize                int
	typicalDistributionRatio float64
	done                     bool
	totalChars               int
	freqChars                int
}

func NewCharDistributionAnalysis() CharDistributionAnalysis {
	return CharDistributionAnalysis{
		EnoughDataThreshold:  1024,
		SureYes:              0.99,
		SureNo:               0.11,
		MinimumDataThreshold: 3,

		// Mapping table to get frequency order from char order (get from GetOrder())
		charToFreqOrder: nil,
		tableSize:       0,

		// This is a constant value which varies from language to language,
		// used in calculating confidence.  See
		// http://www.mozilla.org/projects/intl/UniversalCharsetDetection.html
		// for further detail.
		typicalDistributionRatio: 0,
		done:                     false,
		totalChars:               0,
		freqChars:                0,
	}
}

// Reset analyser, clear any state
func (c *CharDistributionAnalysis) Reset() {
	// If this flag is set to True, detection is done and conclusion has
	// been made
	c.done = false
	c.totalChars = 0 // Total characters encountered
	// The number of characters whose frequency order is less than 512
	c.freqChars = 0
}

// Feed a character with known length
func (c *CharDistributionAnalysis) Feed(buff []byte, length int) {
	// var (
	// 	order int
	// )
	//
	// if length == 2 {
	// 	order = 1
	// }
}
