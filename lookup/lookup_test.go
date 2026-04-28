package lookup

import "testing"

func TestLookupEncoding(t *testing.T) {
	tests := map[string]bool{
		"US-ASCII":  true,
		"Shift_JIS": true,
		"csGB2312":  true,
		"cp932":     false, // Supported charset but no decoder available
	}

	for name, expectDecoder := range tests {
		enc, err := LookupEncoding(name)
		if err != nil {
			t.Fatalf("LookupEncoding(%s) returned error: %v", name, err)
		}
		if expectDecoder && enc == nil {
			t.Fatalf("expected decoder for %s", name)
		}
		if !expectDecoder && enc != nil {
			t.Fatalf("did not expect decoder for %s", name)
		}
	}
}
