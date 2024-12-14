<div align=center>

## chardet: Go character encoding detector
[![Go Reference](https://pkg.go.dev/badge/github.com/wlynxg/chardet.svg)](https://pkg.go.dev/github.com/wlynxg/chardet)
[![License](https://img.shields.io/github/license/wlynxg/chardet.svg?style=flat)](https://github.com/wlynxg/chardet)
[![Go Report Card](https://goreportcard.com/badge/github.com/wlynxg/chardet)](https://goreportcard.com/report/github.com/wlynxg/chardet)

</div>

# Introduction

This is a Go port of the python's [chardet](https://github.com/chardet/chardet) library. Much respect and appreciation to the original authors for their excellent work.

chardet is a character encoding detector library written in Go. It helps you automatically detect the character encoding of text content.

# Installation

To install chardet, use `go get`:

```bash
go get github.com/wlynxg/chardet
```

## Supported Encodings & Languages

**Support Encodings**:

<details>
  <summary>Expand the list of supported encodings</summary>

- **Ascii**
- **UTF-8**
- **UTF-8-SIG**
- **UTF-16**
- **UTF-16LE**
- **UTF-16BE**
- **UTF-32**
- **UTF-32BE**
- **UTF-32LE**
- **GB2312**
- **HZ-GB-2312**
- **SHIFT_JIS**
- **Big5**
- **Johab**
- **KOI8-R**
- **TIS-620**
- **MacCyrillic**
- **MacRoman**
- **EUC-TW**
- **EUC-KR**
- **EUC-JP**
- **CP932**
- **CP949**
- **Windows-1250**
- **Windows-1251**
- **Windows-1252**
- **Windows-1253**
- **Windows-1254**
- **Windows-1255**
- **Windows-1256**
- **Windows-1257**
- **ISO-8859-1**
- **ISO-8859-2**
- **ISO-8859-5**
- **ISO-8859-6**
- **ISO-8859-7**
- **ISO-8859-8**
- **ISO-8859-9**
- **ISO-8859-13**
- **ISO-2022-CN**
- **ISO-2022-JP**
- **ISO-2022-KR**
- **X-ISO-10646-UCS-4-3412**
- **X-ISO-10646-UCS-4-2143**
- **IBM855**
- **IBM866**

</details>

**Support Languages**:
<details>
<summary>Expand the list of supported languages</summary>
- Chinese
- Japanese
- Korean
- Hebrew
- Russian
- Greek
- Bulgarian
- Thai
- Turkish

</details>

# Usage

## Basic Usage

The simplest way to use chardet is with the `Detect` function:

```go
package main

import (
	"fmt"
	"github.com/wlynxg/chardet"
)

func main() {
	data := []byte("Your text data here...")
	result := chardet.Detect(data)
	fmt.Printf("Detected result: %+v\n", result) 
    //Output: Detected result: {Encoding:Ascii Confidence:1 Language:}
}
```

## Advanced Usage

For handling large amounts of text, you can use the detector incrementally. This allows the detector to stop as soon as it reaches sufficient confidence in its result.
```go
package main

import (
	"fmt"
	"github.com/wlynxg/chardet"
)

func main() {
	// Create a detector instance
	detector := chardet.NewUniversalDetector(0)
	// Process text in chunks
	chunk1 := []byte("First chunk of text...")
	chunk2 := []byte("Second chunk of text...")
	detector.Feed(chunk1)
	detector.Feed(chunk2)
	// Get the result
	result := detector.GetResult()
	fmt.Printf("Detected result: %+v\n", result)
	// Output: Detected result: {Encoding:Ascii Confidence:1 Language:}
}
```

## Processing Multiple Files

You can reuse the same detector instance for multiple files by using the `Reset()` method:
```go
package main

import (
	"fmt"
	"os"
	"github.com/wlynxg/chardet"
)

func main() {
	detector := chardet.NewUniversalDetector(0)
	files := []string{"file1.txt", "file2.txt"}
	for _, file := range files {
		detector.Reset()
		data, err := os.ReadFile(file)
		if err != nil {
			continue
		}
		detector.Feed(data)
		result := detector.GetResult()
		fmt.Printf("File %s encoding: %+v\n", file, result)
	}
}
```

# License

`chardet` is licensed under the [MIT License](LICENSE), 100% free and open-source, forever.
