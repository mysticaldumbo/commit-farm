package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateRandomString(length int) string {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()+=-_\\|}]{[~`"
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(chars[rand.Intn(len(chars))])
	}
	return sb.String()
}

func main() {
	for {
		randString := generateRandomString(20)
		err := ioutil.WriteFile("bot.txt", []byte(randString), 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			os.Exit(1)
		}
		cmd := exec.Command("git", "add", ".")
		cmd.Run()
		cmd = exec.Command("git", "commit", "-m", randString)
		cmd.Run()
		cmd = exec.Command("git", "push", "origin", "main")
		cmd.Run()
		time.Sleep(time.Second)
	}
}
