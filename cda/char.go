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

	orderFunc OrderFunc
}

func NewCharDistributionAnalysis(charToFreqOrder []int, tableSize int,
	typicalDistributionRatio float64, order OrderFunc) CharDistributionAnalysis {
	if order == nil {
		order = func(bytes []byte) int { return -1 }
	}

	return CharDistributionAnalysis{
		EnoughDataThreshold:  1024,
		SureYes:              0.99,
		SureNo:               0.11,
		MinimumDataThreshold: 3,

		// Mapping table to get frequency order from char order (get from GetOrder())
		charToFreqOrder: charToFreqOrder,
		tableSize:       tableSize,

		// This is a constant value which varies from language to language,
		// used in calculating confidence.  See
		// http://www.mozilla.org/projects/intl/UniversalCharsetDetection.html
		// for further detail.
		typicalDistributionRatio: typicalDistributionRatio,
		done:                     false,
		totalChars:               0,
		freqChars:                0,

		orderFunc: order,
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
	var (
		order int
	)

	if length == 2 && c.orderFunc != nil {
		// we only care about 2-bytes character in our distribution analysis
		order = c.orderFunc(buff)
	} else {
		order = -1
	}

	if order >= 0 {
		c.totalChars++
		// order is valid
		if order < c.tableSize && c.charToFreqOrder[order] < 512 {
			c.freqChars++
		}
	}
}

// GetConfidence return confidence based on existing data
func (c *CharDistributionAnalysis) GetConfidence() float64 {
	// if we didn't receive any character in our consideration range,
	// return negative answer
	if c.totalChars <= 0 || c.freqChars <= c.MinimumDataThreshold {
		return c.SureNo
	}

	if c.totalChars != c.freqChars {
		r := float64(c.freqChars) / (float64(c.totalChars-c.freqChars) * c.typicalDistributionRatio)
		if r < c.SureYes {
			return r
		}
	}

	// normalize confidence (we don't want to be 100% sure)
	return c.SureYes
}

func (c *CharDistributionAnalysis) GotEnoughData() bool {
	// It is not necessary to receive all data to draw conclusion.
	// For charset detection, certain amount of data is enough
	return c.totalChars > c.EnoughDataThreshold
}

func (c *CharDistributionAnalysis) CharSetName() string {
	return ""
}
