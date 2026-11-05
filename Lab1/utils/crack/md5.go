package crack

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
)

// CrackMD5 reads wordlistPath line by line, hashes each candidate with MD5,
// and checks it against targetHash. Every attempt is printed to the screen
// AND written into logFile. Returns the password if found, or "" if not.
func CrackMD5(targetHash string, wordlistPath string, logFile *os.File) string {
	file, err := os.Open(wordlistPath)
	if err != nil {
		fmt.Println("Error opening wordlist:", err)
		return ""
	}

	scanner := bufio.NewScanner(file)
	attempt := 0

	for scanner.Scan() {
		candidate := scanner.Text()
		attempt++

		sum := md5.Sum([]byte(candidate))
		hash := fmt.Sprintf("%x", sum)

		line := fmt.Sprintf("[%d] trying %q -> %s\n", attempt, candidate, hash)
		fmt.Print(line)
		logFile.WriteString(line)

		if hash == targetHash {
			result := fmt.Sprintf("\nMATCH FOUND after %d attempts: %q -> %s\n", attempt, candidate, hash)
			fmt.Print(result)
			logFile.WriteString(result)
			file.Close()
			return candidate
		}
	}

	file.Close()
	fmt.Println("\nNo match found in wordlist.")
	return ""
}
