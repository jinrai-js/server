package global_cashe

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
)

func GetRequestKey(request *http.Request) string {
	hash := sha256.New()

	hash.Write([]byte(request.Method))
	hash.Write([]byte(request.URL.String()))

	if request.Body != nil && request.Body != http.NoBody {
		body, _ := io.ReadAll(request.Body)
		hash.Write(body)
		request.Body = io.NopCloser(bytes.NewReader(body))
	}

	return hex.EncodeToString(hash.Sum(nil))
}
