package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	log.SetFlags(log.Lshortfile)

	if len(os.Args) < 2 {
		log.Fatal("need subcommand as first argument")
	}
	_ = doArchive()
	_ = doDebSource()
	_ = doPurge()
	_ = doLint()

	switch os.Args[1] {
	case "install":
		doInstall(os.Args[2:])
	case "test":
		doTest(os.Args[2:])
	default:
		log.Fatal("unknown command ", os.Args[1])
	}
}

func doInstall(cmdline []string) {
	var output string
	flag.StringVar(&output, "o", "", "output binary name")
	flag.CommandLine.Parse(cmdline)

	if output == "" {
		log.Fatal("No output path specified")
	}

	var buildCmd *exec.Cmd
	if runtime.GOOS == "windows" {
		buildCmd = exec.Command("go", "build", "-o", output+".exe", "main.go")
	} else {
		buildCmd = exec.Command("go", "build", "-o", output, "main.go")
	}

	outputBytes, err := buildCmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to build: %v\nOutput: %s", err, string(outputBytes))
	}
}

func doTest(cmdline []string) {
	flag.CommandLine.Parse(cmdline)
	testCmd := exec.Command("go", "test", "./...")
	err := testCmd.Run()
	if err != nil {
		log.Fatalf("Failed to test: %v", err)
	}
}

func doLint() error {
	return nil
}

func doArchive() error {
	return nil
}

func doDebSource() error {
	return nil
}

func doPurge() error {
	return nil
}
