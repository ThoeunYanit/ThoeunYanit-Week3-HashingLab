package main

import (
	"fmt"
	"os"

	"crackinglab/utils/crack"
)

func main() {
	targetHash := "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"
	wordlistPath := "wordlists/nord_vpn.txt"

	logFile, err := os.Create("verbose.txt")
	if err != nil {
		fmt.Println("Error creating verbose.txt:", err)
		return
	}

	password := crack.CrackSHA1(targetHash, wordlistPath, logFile)
	logFile.Close()

	if password != "" {
		fmt.Println("\nCracked! Password is:", password)
	} else {
		fmt.Println("\nPassword not found.")
	}
}
