package main

import (
	"fmt"
	"os"

	"crackinglab/utils/crack"
)

func main() {
	targetHash := "6a85dfd77d9cb35770c9dc6728d73d3f"
	wordlistPath := "wordlists/nord_vpn.txt"

	logFile, err := os.Create("verbose.txt")
	if err != nil {
		fmt.Println("Error creating verbose.txt:", err)
		return
	}

	password := crack.CrackMD5(targetHash, wordlistPath, logFile)
	logFile.Close()

	if password != "" {
		fmt.Println("\nCracked! Password is:", password)
	} else {
		fmt.Println("\nPassword not found.")
	}
}
