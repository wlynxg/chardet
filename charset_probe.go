package chardet

type CharSetProbe interface {
	Feed([]byte) ProbingState
}
