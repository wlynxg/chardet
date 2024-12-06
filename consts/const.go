package consts

const (
	UTF8ProbeCharsetName     = "utf-8"
	ShiftJISProbeCharsetName = "SHIFT_JIS"
	CP932CharsetName         = "CP932"
)

const (
	ChineseLanguage  = "Chinese"
	JapaneseLanguage = "Japanese"
)

const (
	HzModelName        = "HZ-GB-2312"
	Iso2022cnModelName = "ISO-2022-CN"
	Iso2022jpModelName = "ISO-2022-JP"
	UTF8ModelName      = "UTF-8"
	UTF16LeModelName   = "UTF-16LE"
	UTF16BeModelName   = "UTF-16BE"
	ShiftJisModelName  = "Shift_JIS"
	GB2312ModelName    = "GB2312"
	EucTwModelName     = "x-euc-tw"
	EucKrModelName     = "EUC-KR"
	EucJpModelName     = "EUC-JP"
	CP949ModelName     = "CP949"
	Big5ModelName      = "Big5"
)

type LangFilter byte

const (
	UnknownLangFilter           LangFilter = 0
	ChineseSimplifiedLangFilter LangFilter = 1 << iota
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
	PureAsciiInputState InputState = iota + 1
	EcsAsciiInputState
	HighByteInputState
)

type ProbingState byte

const (
	DetectingProbingState ProbingState = iota + 1
	FoundItProbingState
	NotMeProbingState
)

type MachineState byte

const (
	StartMachineState MachineState = iota
	ErrorMachineState
	ItsMeMachineState
)
