package consts

const (
	UTF8SIG  = "UTF-8-SIG"
	UTF32    = "UTF-32"
	UTF16    = "UTF-16"
	UCS43412 = "X-ISO-10646-UCS-4-3412"
	UCS42143 = "X-ISO-10646-UCS-4-2143"
)

var (
	UTF8BOM     = []byte{0xEF, 0xBB, 0xBF}
	UTF32LEBOM  = []byte{0xFF, 0xFE, 0x00, 0x00}
	UTF32BEBOM  = []byte{0x00, 0x00, 0xFE, 0xFF}
	UTF16LEBOM  = []byte{0xFE, 0xFF}
	UTF16BEBOM  = []byte{0xFF, 0xFE}
	UCS43412BOM = []byte{0xFE, 0xFF, 0x00, 0x00}
	UCS42143BOM = []byte{0x00, 0x00, 0xFF, 0xFE}
)
