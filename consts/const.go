package consts

const (
	UTF8ProbeCharsetName     = "utf-8"
	ShiftJISProbeCharsetName = "SHIFT_JIS"
	CP932CharsetName         = "CP932"
)

const (
	ChineseLanguage   = "Chinese"
	JapaneseLanguage  = "Japanese"
	KoreanLanguage    = "Korean"
	HebrewLanguage    = "Hebrew"
	RussianLanguage   = "Russian"
	GreekLanguage     = "Greek"
	BulgarianLanguage = "Bulgarian"
	ThaiLanguage      = "Thai"
	TurkishLanguage   = "Turkish"
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
	EucTwModelName     = "EUC-TW"
	EucKrModelName     = "EUC-KR"
	EucJpModelName     = "EUC-JP"
	CP949ModelName     = "CP949"
	Big5ModelName      = "Big5"
	JohabName          = "Johab"
	Windows1251        = "windows-1251"
	Windows1255        = "windows-1255"
	Windows1253        = "windows-1253"
	Koi8R              = "KOI8-R"
	ISO88595           = "ISO-8859-5"
	ISO88591           = "ISO-8859-1"
	IBM855             = "IBM855"
	IBM866             = "IBM866"
	MacCyrillic        = "MacCyrillic"
	ISO88597           = "ISO-8859-7"
	TIS620             = "TIS-620"
)

// LangFilter represents the different language filters we can apply to a "UniversalDetector".
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

// InputState represents the different states a universal detector can be in.
type InputState byte

const (
	PureAsciiInputState InputState = iota + 1
	EcsAsciiInputState
	HighByteInputState
)

// ProbingState represents the different states a probe can be in.
type ProbingState byte

const (
	DetectingProbingState ProbingState = iota + 1
	FoundItProbingState
	NotMeProbingState
)

// MachineState represents the different states a state machine can be in.
type MachineState byte

const (
	StartMachineState MachineState = iota
	ErrorMachineState
	ItsMeMachineState
)

// SequenceLikelihood represents the likelihood of a character following the previous one.
type SequenceLikelihood byte

const (
	NgativeSequenceLikelihood SequenceLikelihood = iota
	UnlikelySequenceLikelihood
	LikelySequenceLikelihood
	PositiveSequenceLikelihood
	LikelihoodCategories = 4
)

// CharacterCategory represents the different categories language models for
//
//	``SingleByteCharsetProbe`` put characters into.
//
//	Anything less than CONTROL is considered a letter.
type CharacterCategory byte

const (
	UndefinedCharacterCategory CharacterCategory = 255 - iota
	LineBreakCharacterCategory
	SymbolCharacterCategory
	DigitCharacterCategory
	ControlCharacterCategory
)
