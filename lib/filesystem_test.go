package lib

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOpen(t *testing.T) {
	fs := mockFS{}
	_, error := fs.Open("fake.file")
	if error != nil {
		t.Error()
	}
}

func TestDownloadFile(t *testing.T) {
	want := "Success!"
	fs := mockFS{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(want))
	}))

	client := NewHTTPClient(srv.Client(), srv.URL)
	_, resultError := client.DownloadFile(fs)
	if resultError != nil {
		t.Error()
	}
}