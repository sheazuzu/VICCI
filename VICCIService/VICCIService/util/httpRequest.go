package util

import (
	"bytes"
	"encoding/base64"
	"net/http"
	"net/http/httptest"
)

func PerformRequest(handler http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {

	request, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	hash := base64.StdEncoding.EncodeToString([]byte("vicci:HeyHeyVicci_HeyVicciHey"))
	request.Header.Set("Authorization","Basic "+hash)

	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)
	return recorder
}


