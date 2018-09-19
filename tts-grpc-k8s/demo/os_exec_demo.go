package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	// lookPath("java")
	// lookPath("go")
	// lookPath("rust")
	// lookPath("/ruby")
	// fmt.Println("-------")
	// command()
	// commandEnv()
	// commandContext()
	// 注意：Output()和CombinedOutput()不能够同时使用，因为command的标准输出只能有一个，同时使用的话便会定义了两个，便会报错
	// combinedOutput()
	// output()
	// run()
	// start()
	stderrPipe()
}

// lookPath 在环境变量中查找可行二进制文件
func lookPath(path string) {
	rpath, err := exec.LookPath(path)
	if err != nil {
		log.Printf("installing %s is in your future", path)
	}
	log.Printf("%s is available at %s\n", path, rpath)
}

func command() {
	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("in all caps:%q\n", out.String())
}

func commandEnv() {
	cmd := exec.Command("go")
	cmd.Env = append(os.Environ(), "FOO=duplicate_value", "FOO=actual_value")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func commandContext() {
	log.Println("start...")
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	if err := exec.CommandContext(ctx, "sleep", "5000").Run(); err != nil {
		log.Fatal(err)
	}
}

func combinedOutput() {
	cmd := exec.Command("sh", "-c", "echo stdout; echo 1>&2 stderr")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)
}

func output() {
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the date is %s\n", out)
}

func run() {
	cmd := exec.Command("sleep", "1")
	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	log.Printf("Command finished with err:%v", err)
}

func start() {
	cmd := exec.Command("sleep", "5")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("command finished with error:%v", err)
}

func stderrPipe() {
	cmd := exec.Command("sh", "-c", "echo stdout; echo 1>&2 stderr")
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	sp, _ := ioutil.ReadAll(stderr)
	fmt.Printf("%s\n", sp)
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
