package main

import (
	"fmt"
	"net"
	"net/http"
)

var (
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

func getHTML(status string) string {
	message := "DevOps, Mr. Cat speaking, how can I help you?"
	image := "1.3.0"
	color := "orange"
	if status == "alive" {
		message = "High five budy, I'm healthy as ever!"
		image = "1.3.1"
		color = "green"
	}
	if status == "dead" {
		message = "I don't want to die Mr. Stark..."
		image = "1.3.2"
		color = "red"
	}
	return `<!DOCTYPE html><html><body><center>
	<img src="https://raw.githubusercontent.com/twogg-git/k8s-workshop/1.3-liveness/src/` + image + `.png">
	<h1 style="color:` + color + `">` + message + `</h3>	
	<h2 style="color:green">Playing with Kubernetes</h1>
	<h2 style="color:blue">Server IP ` + getServerIP() + `</h2>
	<h3 style="color:blue">Version twogghub/k8s-workshop:1.3-liveness</h3>	
	</center></body></html>`
}

func playHome(w http.ResponseWriter, r *http.Request) {
	if bannedIp == getServerIP() {
		fmt.Fprintf(w, getHTML("dead"))
	} else {
		fmt.Fprintf(w, getHTML("home"))
	}
}

func playHealth(w http.ResponseWriter, r *http.Request) {
	if getServerIP() == bannedIp {
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		fmt.Fprintf(w, getHTML("alive"))
	}
}

func playKillMe(w http.ResponseWriter, r *http.Request) {
	bannedIp = getServerIP()
	fmt.Fprintf(w, getHTML("dead"))
}

func main() {
	http.HandleFunc("/", playHome)
	http.HandleFunc("/health", playHealth)
	http.HandleFunc("/killme", playKillMe)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
