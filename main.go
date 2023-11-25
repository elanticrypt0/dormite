package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

var rCounter = 0

func main() {
	app_port := "3004"
	sleepHandler := http.HandlerFunc(wait4Sleep)
	http.Handle("/", sleepHandler)

	fmt.Println("Escuchando puerto ", app_port)
	http.ListenAndServe(":"+app_port, nil)

}

func wait4Sleep(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.RemoteAddr, "192.168") {
		fmt.Println("Request accepted")
		fmt.Printf("\nRequests Number: %d \n", rCounter)
		rCounter++
		time_query := r.URL.Query().Get("t")
		msg := ""
		if time_query != "c" {
			fmt.Println("###########")
			fmt.Println("New Request")
			if time_query != "-1" && time_query != "0" {
				t, _ := strconv.Atoi(time_query)
				msg = fmt.Sprintf("\nMe voy a dormir en : %d minutos", t)
				dormite(t)
			}
		} else {
			msg = "\nNo me voy a dormir nada!"
			noteduemas()
		}
		fmt.Println(msg)
		w.WriteHeader(200)
		w.Write([]byte(msg))
	} else {
		w.WriteHeader(500)
	}
}

func dormite(t int) {

	if t > 0 {
		var cmd *exec.Cmd

		if runtime.GOOS == "windows" {
			args := fmt.Sprintf(" shutdown /s /t %d", t*60)
			cmd = exec.Command("cmd.exe", "/C", args)
		} else {
			time := fmt.Sprintf("+%d", t)
			cmd = exec.Command("shutdown", "-h", time)
		}
		// fmt.Printf("%v", cmd)
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func noteduemas() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd.exe", "/a")
	} else {
		cmd = exec.Command("shutdown", "-c")
	}
	// fmt.Printf("%v", cmd)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
