package consts

var legacyToCanonical = map[string]string{
	Ascii:       "US-ASCII",
	ShiftJis:    "Shift_JIS",
	Johab:       "KS_C_5601-1987",
	MacRoman:    "macintosh",
	MacCyrillic: "x-mac-cyrillic",
}

var canonicalToLegacy map[string]string

func init() {
	canonicalToLegacy = make(map[string]string, len(legacyToCanonical))
	for legacy, canonical := range legacyToCanonical {
		canonicalToLegacy[canonical] = legacy
	}
}

// CanonicalCharset returns the IANA-compliant charset name for the provided legacy encoding name.
func CanonicalCharset(name string) string {
	if canonical, ok := legacyToCanonical[name]; ok {
		return canonical
	}
	return name
}

// LegacyCharset returns the legacy encoding name for the provided canonical charset.
func LegacyCharset(name string) string {
	if legacy, ok := canonicalToLegacy[name]; ok {
		return legacy
	}
	return name
}
