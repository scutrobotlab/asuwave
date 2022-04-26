package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type casesT []struct {
	method   string
	url      string
	body     interface{}
	wantCode int
}

func ctrlerTest(ctrler func(http.ResponseWriter, *http.Request), name string, cases casesT, t *testing.T) {
	for _, c := range cases {
		var reqBody io.Reader = nil
		if c.body != nil {
			b, _ := json.Marshal(c.body)
			reqBody = strings.NewReader(string(b))
		}
		req := httptest.NewRequest(c.method, c.url, reqBody)
		w := httptest.NewRecorder()
		ctrler(w, req)

		resp := w.Result()
		body, _ := io.ReadAll(resp.Body)
		if resp.StatusCode != c.wantCode {
			t.Errorf("Ctrler %v error: want '%d', got '%d'", name, c.wantCode, resp.StatusCode)
		}

		fmt.Printf("Ctrler %v response: %v\n", name, string(body))
	}
}
