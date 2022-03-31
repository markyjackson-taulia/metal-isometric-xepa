package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func main() {
	// Install XFCE GUI, VNC server, and other necessary packages

	cmd := exec.Command("apk", "add", "--no-cache", "ca-certificates", "curl", "openssl", "sudo", "xvfb", "x11vnc", "xfce4", "xfce4-terminal", "faenza-icon-theme", "bash", "procps", "nano", "git", "pciutils", "gzip", "cpio", "tar", "unzip", "xarchiver")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Set VNC password: ("alpine" but you can set it to whatever)

	cmd = exec.Command("mkdir", "-p", "/root/.vnc")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	cmd = exec.Command("x11vnc", "-storepasswd", "alpine", "/root/.vnc/passwd")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Start GUI and VNC server services

	cmd = exec.Command("export", "DISPLAY=:99")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	cmd = exec.Command("export", "RESOLUTION=1920x1080x24")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	cmd = exec.Command("nohup", "/usr/bin/Xvfb", ":99", "-screen", "0", "$RESOLUTION", "-ac", "+extension", "GLX", "+render", "-noreset")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}