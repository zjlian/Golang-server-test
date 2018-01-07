package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"time"
	"voting-system-api/tool/rand"
)

var portCode = flag.String("p", "80", "web服务监听的端口")

func main() {
	flag.Parse()
	var portStr = ":" + (*portCode)

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
	//logConnInfo(r)
	fmt.Fprintln(w, "Golang server is running")
}

func randstrHandler(w http.ResponseWriter, r *http.Request) {
	//logConnInfo(r)
	const n = 8192

	tt := time.Now()
	set := make(map[string]int)
	for i := 0; i < n; i++ {
		set[rand.GetRS32()]++
	}
	rt := time.Since(tt)
	fmt.Fprintf(w, "预计生成 %d 个字符， 实际生成 %d 个字符\n耗时 %v",
		n, len(set), rt)
}

func ipHandler(w http.ResponseWriter, r *http.Request) {
	//logConnInfo(r)
	fmt.Fprintf(w, "你的IP是: %v", getRequestIP(r))
}
