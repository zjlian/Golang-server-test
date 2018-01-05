package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"voting-system-api/tool/rand"
)

var portCode = flag.String("port", "8088", "web服务监听的端口")

func main() {
	// text := "zjlian"
	// fmt.Println(rand.MD5(text))
	// var l int
	// fmt.Scanf("%d", &l)
	// for i := 0; i < l; i++ {
	// 	fmt.Println(rand.GetRS64())
	// }
	flag.Parse()
	var portStr = ":" + (*portCode)
	fmt.Println("Server running on localhost" + portStr)

	http.HandleFunc("/", greet)
	http.HandleFunc("/randstr", randstrHandler)
	http.HandleFunc("/ip", ipHandler)
	http.ListenAndServe(portStr, nil)
}

func getRequestIP(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return ""
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		return ""
	}
	return r.RemoteAddr
}

func logConnInfo(r *http.Request) {
	fmt.Println("User " + getRequestIP(r) +
		" connected and visited this URL: \"" + r.URL.Path + "\"")
}

func greet(w http.ResponseWriter, r *http.Request) {
	logConnInfo(r)
	fmt.Fprintln(w, "Golang server is running")
}

func randstrHandler(w http.ResponseWriter, r *http.Request) {
	logConnInfo(r)
	var printStr string
	for i := 0; i < 10; i++ {
		printStr += rand.GetRS32() + "\n"
	}
	fmt.Fprintf(w, printStr)
}

func ipHandler(w http.ResponseWriter, r *http.Request) {
	logConnInfo(r)
	fmt.Fprintf(w, "你的IP是: %v", getRequestIP(r))
}
