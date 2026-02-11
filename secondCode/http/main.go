package main

import (
	"fmt"
	"net/http"
)

func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, `[{"id":1,"name":张三,"age":18}]`)
		return
	}
	http.Error(w, "Method not allowed", http.StatusNotFound) //方法不允许
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("请求开始", r.RequestURI)
		next.ServeHTTP(w, r)
		fmt.Println("请求结束", r.Body)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/students", getStudents)

	handle := loggingMiddleware(mux)

	http.ListenAndServe(":8080", handle)
}
