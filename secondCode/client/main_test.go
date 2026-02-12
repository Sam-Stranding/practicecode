package main

import (
	"io"
	"net/http"
	"testing"
)

func Test_client(t *testing.T) {
	resq, err := http.Post("http://localhost:8080/ping", "", nil)
	if err != nil {
		t.Error(err)
		return
	}

	body, _ := io.ReadAll(resq.Body)
	defer resq.Body.Close()

	//t.Log(string(body))
	t.Logf("%s", body)
	t.Error("test")
}
