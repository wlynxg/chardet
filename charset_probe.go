package chardet

type ICharSetProbe interface {
	Feed([]byte) ProbingState
	Reset()
}
