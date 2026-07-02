package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestEmptyHeader(t *testing.T) {
	headers := make(http.Header)
	headers.Set("", "*")
	_, err := GetAPIKey(headers)
	if !reflect.DeepEqual(ErrNoAuthHeaderIncluded, err) {
		t.Fatalf("expected: %v, got: %v", ErrNoAuthHeaderIncluded, err)
	}
}
