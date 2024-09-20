package main

/*
run 命令：运行一个容器
ps 命令：列出正在运行的容器
images 命令：列出本地镜像

使用方法：
运行容器：go run docker.go run [镜像名]
2. 列出容器：go run docker.go ps
3. 列出镜像：go run docker.go images

这个实现实际上是调用了系统中已安装的 Docker 命令。

要真正实现一个完整的 Docker 功能，需要更复杂的代码来处理容器化、网络、存储等方面的问题。

这个示例主要是为了展示如何使用 Golang 创建一个基本的命令行界面来模拟 Docker 的一些核心命令。
*/
import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("使用方法: go run docker.go [命令]")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "run":
		runContainer()
	case "ps":
		listContainers()
	case "images":
		listImages()
	default:
		fmt.Printf("未知命令: %s\n", command)
		os.Exit(1)
	}
}

func runContainer() {
	if len(os.Args) < 3 {
		fmt.Println("使用方法: go run docker.go run [镜像名]")
		os.Exit(1)
	}

	imageName := os.Args[2]
	cmd := exec.Command("docker", "run", imageName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("运行容器失败: %v\n", err)
		os.Exit(1)
	}
}

func listContainers() {
	cmd := exec.Command("docker", "ps")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("列出容器失败: %v\n", err)
		os.Exit(1)
	}
}

func listImages() {
	cmd := exec.Command("docker", "images")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("列出镜像失败: %v\n", err)
		os.Exit(1)
	}
}