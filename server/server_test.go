package server

import (
	"io"
	"net/http"
	"testing"
)

func Test_server(t *testing.T) {
	StartServer()
	if result, callErr := http.Get("http://localhost:8080/test"); callErr != nil {
		t.Errorf("call to web server failed with %v", callErr)
	} else {
		defer result.Body.Close()
		if responseText, responseTextErr := io.ReadAll(result.Body); responseTextErr != nil {
			t.Errorf("response could not be read into a string %v", responseTextErr)
		} else {
			textResult := string(responseText)
			if textResult != "hi" {
				t.Errorf("the response text shoult be \"Hi\" it is acturally %v", responseText)
			}
		}
	}
}
