package lib

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOpen(t *testing.T) {
	fs := mockFS{}
	_, err := fs.Open("fake.file")
	if err != nil {
		t.Error()
	}
}

func TestCopy(t *testing.T) {
	fs := mockFS{}
	dstFile := fs.NewFile(0, "dstFile")
	srcFile := fs.NewFile(0, "srcFile")
	_, err := fs.Copy(dstFile, srcFile)
	if err != nil {
		t.Error()
	}
}

func TestCreate(t *testing.T) {
	fs := mockFS{}
	_, err := fs.Create("fake.file")
	if err != nil {
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