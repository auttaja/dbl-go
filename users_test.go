package dbl_go

import (
	"net/http"
	"os"
	"testing"
)

func TestDBL_GetUser(t *testing.T) {
	dbl := NewDBL(os.Getenv("TOKEN"), &http.Client{})
	_, err := dbl.GetUser("109710323094683648")
	if err != nil {
		t.Errorf("GetUser threw an error: %v", err)
	}
}
