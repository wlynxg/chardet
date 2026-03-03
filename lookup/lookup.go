package lookup

import (
	"strings"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode/utf32"
)

// LookupEncoding returns a golang.org/x/text/encoding for the provided charset name.
// The helper understands all charset values returned by github.com/wlynxg/chardet.
// It returns (encoding, nil) on success, (nil, err) when the charset is unknown,
// and (nil, nil) when the charset is valid but no compatible decoder exists.
func LookupEncoding(name string) (encoding.Encoding, error) {
	name = strings.ToLower(name)

	enc, err := ianaindex.IANA.Encoding(name)
	if enc != nil {
		return enc, nil
	}

	switch name {
	case "utf-32", "csutf32":
		return utf32.UTF32(utf32.BigEndian, utf32.UseBOM), nil
	case "utf-32be", "csutf32be":
		return utf32.UTF32(utf32.BigEndian, utf32.IgnoreBOM), nil
	case "utf-32le", "csutf32le":
		return utf32.UTF32(utf32.LittleEndian, utf32.IgnoreBOM), nil

	case "gb2312", "csgb2312", "gbk", "cp936", "ms936", "windows-936", "csgbk":
		return simplifiedchinese.GB18030, nil

	case "maccyrillic", "x-mac-cyrillic":
		return charmap.MacintoshCyrillic, nil

	case "euc-tw",
		"cp932", "ms932", "windows-932", "windows-31j",
		"cp949", "ms949", "windows-949":
		return nil, nil
	}

	return nil, err
}
