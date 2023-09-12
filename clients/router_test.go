package clients

import (
	"bytes"
	"encoding/json"
	"io"
	// "log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	// "github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func GetResponseRecorder() *httptest.ResponseRecorder {
	return httptest.NewRecorder()
}
func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
    gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)
    ctx.Request = &http.Request{
      Header: make(http.Header),
      URL:    &url.URL{},
    }
    return ctx
}
// func MockJsonGet(c *gin.Context) {
// 	c.Request.Method = "GET"
// 	c.Request.Header.Set("Content-Type", "application/json")
// 	c.Set("user_id", 1)

// 	// set path params
// 	c.Params = []gin.Param{
// 		{
// 			Key:   "id",
// 			Value: "1",
// 		},
// 	}
// }
func MockJsonPost(c *gin.Context, content interface{}) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")
	// c.Set("message", "")
	// c.Set("customer_id", 3)

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}
// func TestInitRouter(t *testing.T) {
// 	router := InitRouter()
// 	w := GetResponseRecorder()
// 	reqInfo := RequestInfo{
//     CustomerID: 2,
//     Message: "",
// 	}
// 	// marshall data to json (like json_encode)
// 	marshalled, err := json.Marshal(reqInfo)
// 	if err != nil {
//     log.Fatalf("impossible to marshall reqInfo: %s", err)
// 	}
// 	req, err := http.NewRequest("POST", "/review", bytes.NewReader(marshalled))
// 	if err != nil {
//     log.Fatalf("impossible to build request: %s", err)
// 	}
// 	// handler := http.HandlerFunc(getReview)
// 	// handler.ServeHttp(w, req)

// 	assert.Equal(t, "Hi Bob, Welcome to Connectly.ai! How can I help you? You can ask me something like below: 1. How to subscribe service from connectly.ai?  2. I would like to know more about connectly.ai 3. I would like to provide feedback for connectly.ai ", w.Body.message)
// }

func TestGetUUID(t *testing.T) {
    result := getUUID()
    if len(result) <= 0 {
    t.Errorf("getUUID() doesn't generate UUID correctly.")
    }
		result2 := getUUID()
    if result2 == result {
    t.Errorf("getUUID() generates the same UUID.")
    }
}
// func TestGetReview(t *testing.T) {

// }