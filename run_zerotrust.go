package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

func main() {
	user := "user"
	ip_and_port := "ip:port"
	command := "command"

	// Connect to the server
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password("ansitest"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", ip_and_port, config)
	if err != nil {
		fmt.Println("Failed to connect to the server:", err)
		return
	}
	defer conn.Close()

	// Create a session
	session, err := conn.NewSession()
	if err != nil {
		fmt.Println("Failed to create a session:", err)
		return
	}
	defer session.Close()

	// Set up password input
	stdin, err := session.StdinPipe()
	if err != nil {
		fmt.Println("Failed to setup stdin pipe:", err)
		return
	}

	// Run the command
	fmt.Println("Running command...")
	session.Start(command)

	// Pass the password to the command
	fmt.Println("Passing passwd")
	io.WriteString(stdin, "ansitest\n")
	//fmt.Println("Exiting in 30 seconds...")
	//time.Sleep(30 * time.Second)
	fmt.Println("Exiting in:")
	for i := 300; i >= 0; i-- {
		fmt.Printf("%d...\n", i)
		time.Sleep(time.Second)
	}

	// Wait for user to press enter
	fmt.Println("Done! Press enter to exit")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
