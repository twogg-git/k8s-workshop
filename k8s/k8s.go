package main

import (
	"fmt"
	"net"
	"net/http"
)

var (
	version  = "1.0"
	bannedIp = "9"
	html     = `<!DOCTYPE html><html><body><center>
		<img src="https://raw.githubusercontent.com/twogg-git/k8s-intro/master/kubernetes_katacoda.png">
		<h1 style="color:green">Playing with Kubernetes</h1>
		<h2 style="color:blue">Your server IP:` + getServerIP() + `</h2>
		<h3 style="color:blue">Version: twogghub/k8s-workshop:` + version + `</h3>	
		<h3 style="color:red">Banned IP ending number: ` + bannedIp + `</h3>	
		</center></body></html>`
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
	fmt.Fprintf(w, html)
}

func playWithIP(w http.ResponseWriter, r *http.Request) {
	ip := getServerIP()
	if ip[len(ip)-1:] == bannedIp {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		fmt.Fprintf(w, "Play cool, I'm alive!!!")
	}
}

func main() {
	http.HandleFunc("/", playHome)
	http.HandleFunc("/status", playWithIP)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
