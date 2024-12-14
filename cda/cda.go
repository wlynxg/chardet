package cda

type OrderFunc func([]byte) int

type GetOrderFunc func([]byte) (int, int)

type Analyzer interface {
	Reset()
	Feed([]byte, int)
	GotEnoughData() bool
	GetConfidence() float64
	CharSetName() string
}
