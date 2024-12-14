package chardet

import (
	"github.com/wlynxg/chardet/consts"
	"github.com/wlynxg/chardet/probe"
	"sort"
)

// Detect the encoding of the given byte string.
func Detect(buf []byte) Result {
	d := NewUniversalDetector(consts.UnknownLangFilter)
	d.Feed(buf)
	return d.GetResult()
}

// DetectAll the possible encodings of the given byte string.
func DetectAll(buf []byte) []Result {
	d := NewUniversalDetector(consts.UnknownLangFilter)
	d.Feed(buf)
	result := d.GetResult()

	if d.inputState == consts.HighByteInputState {
		var (
			results []Result
			probes  []probe.Probe
		)

		for _, p := range d.charsetProbes {
			switch rp := p.(type) {
			case *probe.CharSetGroupProbe:
				probes = append(probes, rp.Probes()...)
			default:
				probes = append(probes, p)
			}
		}

		for _, setProbe := range probes {
			if setProbe.GetConfidence() > d.MinimumThreshold {
				charsetName := setProbe.CharSetName()
				if d.hasWinBytes {
					// Use Windows encoding name instead of ISO-8859 if we saw any
					// extra Windows-specific bytes
					if n, ok := d.IsoWinMap[setProbe.CharSetName()]; ok {
						charsetName = n
					}
				}
				results = append(results, Result{
					Encoding:   charsetName,
					Confidence: setProbe.GetConfidence(),
					Language:   setProbe.Language(),
				})
			}
		}

		if len(results) > 0 {
			sort.Slice(results, func(i, j int) bool {
				return results[i].Confidence > results[j].Confidence
			})
			return results
		}
	}
	return []Result{result}
}
