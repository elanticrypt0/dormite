package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app_port := "3005"

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Dormite",
	})

	app.Static("/", "./public")

	app.Get("dormite/:t", func(c *fiber.Ctx) error {
		return wait4Sleep(c)
	})

	log.Fatal(app.Listen(":" + app_port))

}

func wait4Sleep(c *fiber.Ctx) error {
	if strings.Contains(c.IP(), "127.0.0.1") || strings.Contains(c.IP(), "192.168") {
		time_query := c.Params("t")
		msg := ""
		if time_query != "c" {
			fmt.Println("###########")
			if time_query != "0" {
				t, _ := strconv.Atoi(time_query)
				msg = fmt.Sprintf("\nMe voy a dormir en : %d minutos", t)
				dormite(t)
			}
		} else {
			msg = "\nNo me voy a dormir nada!"
			noteduemas()
		}
		fmt.Println(msg)
		return c.SendString(msg)
	} else {
		return c.SendString("No encontrado")
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
