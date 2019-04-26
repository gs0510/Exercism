package grep

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//Search looks for pattern in the files.
func Search(pattern string, flags []string, files []string) []string {
	result := []string{}
	flagMap := make(map[string]bool)
	multipleFiles := len(files) > 1
	for _, arg := range flags {
		flagMap[arg] = true
	}
	for _, fileName := range files {
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for n := 1; scanner.Scan(); n++ {
			line := scanner.Text()

			matched := line
			if flagMap["-i"] {
				pattern = strings.ToLower(pattern)
				matched = strings.ToLower(line)
			}
			match := strings.Contains(matched, pattern)

			if flagMap["-x"] {
				match = matched == pattern
			}

			if _, ok := flagMap["-v"]; (!ok && match) || (ok && !match) {
				if _, ok := flagMap["-n"]; ok {
					line = fmt.Sprintf("%d:%s", n, line)
				}
				if multipleFiles {
					line = fmt.Sprintf("%s:%s", fileName, line)
				}
				if _, ok := flagMap["-l"]; ok {
					result = append(result, fileName)
					break
				}
				result = append(result, line)
			}

		}
	}
	return result
}
