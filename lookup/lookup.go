package lookup

import (
	"strings"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode/utf32"
)

// Looks up an `golang.org/x/text/encoding` encoding by charset name
//
// Recognizes all encodings returned by this library for which a corresponding
// or compatible superset encoder exists.
//
// Like `golang.org/x/text/encoding/ianaindex.IANA.Encoding`, this returns
// `encoding, nil` on success, `nil, err` on error and `nil, nil` if the name
// was correct but has no corresponding encoder.
func LookupEncoding(name string) (encoding.Encoding, error) {
	name = strings.ToLower(name)

	// First try stdlib lookup function
	encoding, err := ianaindex.IANA.Encoding(name)
	if encoding != nil {
		return encoding, nil
	}

	switch name {
		// UTF-32 family appears to be an acidental omission
		case "utf-32", "csutf32":
			return utf32.UTF32(utf32.BigEndian, utf32.UseBOM), nil
		case "utf-32be", "csutf32be":
			return utf32.UTF32(utf32.BigEndian, utf32.IgnoreBOM), nil
		case "utf-32le", "csutf32le":
			return utf32.UTF32(utf32.LittleEndian, utf32.IgnoreBOM), nil

		// GB2312 is a subset of GBK which in turn is a subset of GB18030
		case "gb2312", "csgb2312",
		     "gbk", "cp936", "ms936", "windows-936", "csgbk":
			return simplifiedchinese.GB18030, nil

		// MacCyrillic is missing IANA designation
		case "maccyrillic", "x-mac-cyrillic":
			return charmap.MacintoshCyrillic, nil

		// Not supported and not in IANA are:
		case "euc-tw",
		     "cp932", "ms932", "windows-932", "windows-31j",  // Similar to Shift-JIS
		     "cp949", "ms949", "windows-949":                 // Similar to Johab
			return nil, nil

		// Not supported but in IANA are (`err` will be `nil` for these):
		//  * ISO-2022-CN
		//  * ISO-2022-KR
		//  * TIS-620
		default:
			return nil, err
	}
}