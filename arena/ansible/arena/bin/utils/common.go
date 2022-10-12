package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// WriteFile  the rendered templates to files
func WriteFile(ansibleHome, fileName string) (file *os.File) {
	// make sure templates are written to dir inventory/group_vars/ below the repo
	file, err := os.OpenFile(ansibleHome+"/inventory/group_vars/"+fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		os.Exit(1)
	}
	return file
}

func WriteGlobalFile(ansibleHome, fileName string) (file *os.File) {
	// make sure templates are written to dir inventory/group_vars/ below the repo
	file, err := os.OpenFile(ansibleHome+"/inventory/group_vars/all/"+fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		os.Exit(1)
	}
	return file
}

func ReadFile(fileName string) (lines []string) {
	file, err1 := os.Open(fileName)
	if err1 != nil {
		fmt.Println("Fail to open the log file, err:", err1)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') //mention this is the character
		if err == io.EOF {
			if len(line) != 0 {
				lines = append(lines, strings.TrimSuffix(line, "\n"))
			}
			break
		}
		if err != nil {
			fmt.Println("Fail to read file: ", err)
			os.Exit(1)
		}
		lines = append(lines, strings.TrimSuffix(line, "\n"))
	}
	return
}

func DebugByFile(str string) {
	file, err := os.OpenFile("/tmp/debug.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(str)
}
