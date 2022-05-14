package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"runtime"
	"strings"
	"testing"
)

type casesT []struct {
	method   string
	url      string
	body     interface{}
	wantCode int
}

const pass = "\u2713"
const fail = "\u2717"

func ctrlerTest(ctrler func(http.ResponseWriter, *http.Request), cases casesT, t *testing.T) {
	_, file, no, _ := runtime.Caller(1)
	fmt.Printf("From: %s:%d\n", file, no)
	for _, c := range cases {
		fmt.Printf("%s %s ", c.method, c.url)

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
			fmt.Errorf(`%s
	have: %d
	want: %d
`,
				fail,
				resp.StatusCode,
				c.wantCode)
		}

		fmt.Printf(`%s
	status: %d
	body: %s
`,
			pass,
			c.wantCode,
			body)
	}
}
