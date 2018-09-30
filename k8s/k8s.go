package main

import (
	"fmt"
	"net"
	"net/http"
)

var (
	version  = "1.3-liveness"
	bannedIp = "0.0.0.0"
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
		<img src="https://raw.githubusercontent.com/twogg-git/k8s-workshop/master/src/1.3.png">
		<h1 style="color:green">Playing with Kubernetes</h1>
		<h2 style="color:blue">Your server IP:` + getServerIP() + `</h2>
		<h3 style="color:blue">Version: twogghub/k8s-workshop:` + version + `</h3>	
		<h3 style="color:red">Banned IP: ` + bannedIp + `</h3>	
		</center></body></html>`
	fmt.Fprintf(w, html)
}

func playHealth(w http.ResponseWriter, r *http.Request) {
	ip := getServerIP()
	if ip == bannedIp {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		fmt.Fprintf(w, "I'm alive!.. but also dead for IP: "+bannedIp)
	}
}

func playDead(w http.ResponseWriter, r *http.Request) {
	bannedIp = r.URL.Query().Get("ip")
	fmt.Fprintf(w, "Now playing dead for IP: "+bannedIp)
}

func main() {
	http.HandleFunc("/", playHome)
	http.HandleFunc("/health", playHealth)
	http.HandleFunc("/kill", playDead)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
