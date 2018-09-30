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
		<img src="https://raw.githubusercontent.com/twogg-git/k8s-workshop/master/src/1.1.1.png">
		<h1 style="color:red">Playing with Kubernetes</h1>
		<h2 style="color:red">Your server IP ` + getServerIP() + ` :9090</h2>
		<h3 style="color:red">Version twogghub/k8s-workshop:1.1-qaonly</h3>	
		</center></body></html>`
	fmt.Fprintf(w, html)
}

func main() {
	http.HandleFunc("/", playHome)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		panic(err)
	}
}
