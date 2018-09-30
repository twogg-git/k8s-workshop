package main

import (
	"fmt"
	"net"
	"net/http"
)

func getServerIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	return conn.LocalAddr().(*net.UDPAddr).IP.String()
}

func playHome(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html><html><body><center>
		<img src="https://raw.githubusercontent.com/twogg-git/k8s-workshop/master/src/1.2.png">
		<h1 style="color:green">Playing with Kubernetes</h1>
		<h2 style="color:blue">Your server IP:` + getServerIP() + `</h2>
		<h3 style="color:blue">Version: twogghub/k8s-workshop:1.2-yaml</h3>	
		</center></body></html>`
	fmt.Fprintf(w, html)
}

func main() {
	http.HandleFunc("/", playHome)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
