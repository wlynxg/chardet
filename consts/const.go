package consts

const (
	Chinese   = "Chinese"
	Japanese  = "Japanese"
	Korean    = "Korean"
	Hebrew    = "Hebrew"
	Russian   = "Russian"
	Greek     = "Greek"
	Bulgarian = "Bulgarian"
	Thai      = "Thai"
	Turkish   = "Turkish"
)

const (
	Ascii   = "Ascii"
	UTF8    = "UTF-8"
	UTF8SIG = "UTF-8-SIG"
	UTF16   = "UTF-16"
	UTF16Le = "UTF-16LE"
	UTF16Be = "UTF-16BE"
	UTF32   = "UTF-32"
	UTF32Be = "UTF-32BE"
	UTF32Le = "UTF-32LE"

	GB2312   = "GB2312"
	HzGB2312 = "HZ-GB-2312"
	ShiftJis = "SHIFT_JIS"
	Big5     = "Big5"
	Johab    = "Johab"
	Koi8R    = "KOI8-R"
	TIS620   = "TIS-620"

	MacCyrillic = "MacCyrillic"
	MacRoman    = "MacRoman"

	EucTw = "EUC-TW"
	EucKr = "EUC-KR"
	EucJp = "EUC-JP"

	CP932 = "CP932"
	CP949 = "CP949"

	Windows1250 = "Windows-1250"
	Windows1251 = "Windows-1251"
	Windows1252 = "Windows-1252"
	Windows1253 = "Windows-1253"
	Windows1254 = "Windows-1254"
	Windows1255 = "Windows-1255"
	Windows1256 = "Windows-1256"
	Windows1257 = "Windows-1257"

	ISO88591  = "ISO-8859-1"
	ISO88592  = "ISO-8859-2"
	ISO88595  = "ISO-8859-5"
	ISO88596  = "ISO-8859-6"
	ISO88597  = "ISO-8859-7"
	ISO88598  = "ISO-8859-8"
	ISO88599  = "ISO-8859-9"
	ISO885913 = "ISO-8859-13"
	ISO2022CN = "ISO-2022-CN"
	ISO2022JP = "ISO-2022-JP"
	ISO2022KR = "ISO-2022-KR"
	UCS43412  = "X-ISO-10646-UCS-4-3412"
	UCS42143  = "X-ISO-10646-UCS-4-2143"

	IBM855 = "IBM855"
	IBM866 = "IBM866"
)

const (
	UTF8BOM     = "\xEF\xBB\xBF"
	UTF32LEBOM  = "\xFF\xFE\x00\x00"
	UTF32BEBOM  = "\x00\x00\xFE\xFF"
	UTF16LEBOM  = "\xFE\xFF"
	UTF16BEBOM  = "\xFF\xFE"
	UCS43412BOM = "\xFE\xFF\x00\x00"
	UCS42143BOM = "\x00\x00\xFF\xFE"
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
type CharacterCategory int

const (
	UndefinedCharacterCategory CharacterCategory = 255 - iota
	LineBreakCharacterCategory
	SymbolCharacterCategory
	DigitCharacterCategory
	ControlCharacterCategory
)
