package chardet

import (
	"testing"

	"github.com/wlynxg/chardet/consts"
)

func TestNewResultSetsCharset(t *testing.T) {
	tests := map[string]string{
		consts.Ascii:    "US-ASCII",
		consts.ShiftJis: "Shift_JIS",
		consts.Johab:    "KS_C_5601-1987",
		consts.MacRoman: "macintosh",
	}

	for legacy, canonical := range tests {
		res := newResult(legacy, 1.0, "")
		if res.Encoding != legacy {
			t.Fatalf("expected legacy encoding %s, got %s", legacy, res.Encoding)
		}
		if res.Charset != canonical {
			t.Fatalf("expected charset %s, got %s", canonical, res.Charset)
		}
	}
}
