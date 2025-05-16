package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func usage() {
	fmt.Print(`Usage:
go run ./build
go run ./build docker
`)
	flag.PrintDefaults()
}

func run(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	var arg string

	flag.Usage = usage
	flag.Parse()
	if flag.NArg() == 0 {
		arg = ""
	} else if flag.NArg() == 1 {
		arg = flag.Arg(0)
	} else {
		usage()
		os.Exit(1)
	}

	data, err := os.ReadFile(".access_token")
	if err != nil {
		panic(err)
	}
	accessToken := strings.TrimSpace(string(data))

	goBuild := exec.Command("go", "build", "-ldflags", "-X main.accessToken="+accessToken, "-o", "build/trmnl-server")

	if arg == "" {
		run(goBuild)
	} else if arg == "docker" {
		goBuild.Env = os.Environ()
		goBuild.Env = append(goBuild.Env, "CC=zig cc -target x86_64-linux")
		goBuild.Env = append(goBuild.Env, "CXX=zig c++ -target x86_64-linux")
		goBuild.Env = append(goBuild.Env, "CGO_ENABLED=1")
		goBuild.Env = append(goBuild.Env, "GOARCH=amd64")
		goBuild.Env = append(goBuild.Env, "GOOS=linux")
		run(goBuild)
		dockerBuild := exec.Command("docker", "build", ".", "-t", "trmnl-server:latest")
		run(dockerBuild)
		dockerSave := exec.Command("docker", "save", "-o", "build/trmnl-server.tar", "trmnl-server:latest")
		run(dockerSave)
	} else {
		usage()
		os.Exit(1)
	}
}
