package chardet

const (
	ChineseLanguage  = "Chinese"
	JapaneseLanguage = "Japanese"
)

const (
	HzModelName        = "HZ-GB-2312"
	Iso2022cnModelName = "ISO-2022-CN"
	Iso2022jpModelName = "ISO-2022-JP"
)

type LangFilter byte

const (
	ChineseSimplifiedLangFilter = 1 << iota
	ChineseTraditionalLangFilter
	JapaneseLangFilter
	KoreanLangFilter
	NonCjkLangFilter
	ChineseLangFilter = ChineseSimplifiedLangFilter | ChineseTraditionalLangFilter
	CjkLangFilter     = ChineseLangFilter | JapaneseLangFilter | KoreanLangFilter
	AllLangFilter     = ChineseSimplifiedLangFilter | ChineseTraditionalLangFilter |
		JapaneseLangFilter | KoreanLangFilter | NonCjkLangFilter
)

type InputState byte

const (
	PureAsciiInputState = iota + 1
	EcsAsciiInputState
	HighByteInputState
)

type ProbingState byte

const (
	DetectingProbingState = iota + 1
	FoundItProbingState
	NotMeProbingState
)

type MachineState byte

const (
	StartMachineState = iota
	ErrorMachineState
	ItsMeMachineState
)
