package clients

import (
	"bytes"
	"crypto/hmac"

	// "crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	headerNameXSign = "X-Hub-Signature"
	headerNameXSign256 = "X-Hub-Signature-256"
	signaturePrefix = "sha1="
	signaturePrefix256 = "sha256="
)

// errors
var (
	errNoXSignHeader      = errors.New("there is no x-sign header")
	errInvalidXSignHeader = errors.New("invalid x-sign header")
)

// Authorize authorizes web hooks from FB
// https://developers.facebook.com/docs/graph-api/webhooks/getting-started/#validate-payloads
func Authorize(r *http.Request) error {
	signature := r.Header.Get(headerNameXSign256)
	if !strings.HasPrefix(signature, signaturePrefix256) {
		return errNoXSignHeader
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("read all: %w", err)
	}

	// We read the request body and now it's empty. We have to rewrite it for further reads.
	r.Body.Close() //nolint:errcheck
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	validSignature, err := isValidSignature(signature, body)
	if err != nil {
		return fmt.Errorf("is valid signature: %w", err)
	}
	if !validSignature {
		return errInvalidXSignHeader
	}

	return nil
}

func signBody(body []byte) []byte {
	h := hmac.New(sha256.New, []byte(AppSecret))
	h.Reset()
	h.Write(body)

	return h.Sum(nil)
}

func isValidSignature(signature string, body []byte) (bool, error) {
	actualSign, err := hex.DecodeString(signature[len(signaturePrefix256):])
	if err != nil {
		return false, fmt.Errorf("decode string: %w", err)
	}

	return hmac.Equal(signBody(body), actualSign), nil
}