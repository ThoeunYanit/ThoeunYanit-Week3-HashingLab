package main

import (
	"fmt"
	"os"

	"crackinglab/utils/crack"
)

func main() {
	targetHash := "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"
	wordlistPath := "wordlists/nord_vpn.txt"

	logFile, err := os.Create("verbose.txt")
	if err != nil {
		fmt.Println("Error creating verbose.txt:", err)
		return
	}

	password := crack.CrackSHA512(targetHash, wordlistPath, logFile)
	logFile.Close()

	if password != "" {
		fmt.Println("\nCracked! Password is:", password)
	} else {
		fmt.Println("\nPassword not found.")
	}
}
