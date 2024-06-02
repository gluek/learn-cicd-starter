package auth

import (
	"testing"
	"net/http"
	"reflect"
)

func TestGetAPIKey(t *testing.T) {
	testHeaderValid := http.Header{"Authorization": []string{"ApiKey 123456"}}
	testHeaderInvalid := http.Header{"Authorization": []string{"ApiKe"}}
	testHeaderEmpty := http.Header{"Authorization": []string{""}}
	key, err := GetAPIKey(testHeaderValid)
	if !reflect.DeepEqual("", key) {
		t.Fatalf("expected: 123456, got: %v, err: %v", key, err)
	}
	key, err = GetAPIKey(testHeaderInvalid)
	if !reflect.DeepEqual("", key) {
		t.Fatalf("expected: %v, got: %v, err: %v", "", key, err)
	}
	if reflect.DeepEqual(err, nil) {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	key, err = GetAPIKey(testHeaderEmpty)
	if !reflect.DeepEqual("", key) {
		t.Fatalf("expected: %v, got: %v, err: %v", "", key, err)
	}
	if reflect.DeepEqual(err, nil) {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

}
