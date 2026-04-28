+`Result.Encoding` continues to expose the legacy value (e.g. `Ascii`, `SHIFT_JIS`). For new applications use `Result.Charset`, which follows IANA naming.
+
+## Decoding text
+
+Use the optional `github.com/wlynxg/chardet/lookup` helper to map `Result.Charset` to `golang.org/x/text/encoding`:
+
+```go
+package main
+
+import (
+	"fmt"
+
+	"github.com/wlynxg/chardet"
+	"github.com/wlynxg/chardet/lookup"
+)
+
+func main() {
+	data := []byte("Your text data here...")
+	result := chardet.Detect(data)
+
+	enc, err := lookup.LookupEncoding(result.Charset)
+	if err != nil {
+		panic(err)
+	}
+	if enc == nil {
+		fmt.Printf("no decoder for %s\n", result.Charset)
+		return
+	}
+
+	decoded, err := enc.NewDecoder().String(string(data))
+	if err != nil {
+		panic(err)
+	}
+
+	fmt.Println(decoded)
+}
+
*** End Patch
