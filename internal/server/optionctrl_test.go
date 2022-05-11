package server

// import (
// 	"net/http"
// 	"testing"
// )

// func TestOptionCtrl(t *testing.T) {
// 	cases := casesT{
// 		{
// 			http.MethodGet,
// 			"/options",
// 			nil,
// 			http.StatusOK,
// 		},
// 		{
// 			http.MethodPost,
// 			"/options",
// 			nil,
// 			http.StatusMethodNotAllowed,
// 		},
// 		{
// 			http.MethodPut,
// 			"/options",
// 			struct{ Save int }{Save: 6},
// 			http.StatusOK,
// 		},
// 		{
// 			http.MethodDelete,
// 			"/options",
// 			nil,
// 			http.StatusMethodNotAllowed,
// 		},
// 		{
// 			http.MethodGet,
// 			"/options",
// 			nil,
// 			http.StatusOK,
// 		},
// 	}
// 	ctrlerTest(optionCtrl, "optionCtrl", cases, t)
// }
