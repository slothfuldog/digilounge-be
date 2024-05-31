package function

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var logFile *os.File

func InitLogFileLin() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println("Working Directory:", wd)

	filePath := "/logs/"
	currentTime := time.Now().Format("2006-01-02")
	fileName := fmt.Sprintf("%slogfile_%s.txt", filePath, currentTime)

	// Create the directory if it doesn't exist (only for relative paths)
	if !filepath.IsAbs(filePath) {
		err := os.MkdirAll(filePath, 0755)
		if err != nil {
			return err
		}
	}

	// Print the resolved filepath for debugging
	resolvedFilePath := filepath.Join(wd, fileName)
	fmt.Println("Resolved Filepath:", resolvedFilePath)

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	logFile = file
	return nil
}

func InitLogFileWin() error {

	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println("Working Directory:", wd)

	// Define the filepath
	filePath := "logs\\" // Using Windows filepath separator
	currentTime := time.Now().Format("2006-01-02")
	fileName := fmt.Sprintf("%slogfile_%s.txt", filePath, currentTime)

	// Print the resolved filepath for debugging
	resolvedFilePath := filepath.Join(wd, fileName)
	fmt.Println("Resolved Filepath:", resolvedFilePath)

	// Create the directory if it doesn't exist
	err = os.MkdirAll(filepath.Join(wd, filePath), 0755)
	if err != nil {
		return err
	}

	// Open or create the file
	file, err := os.OpenFile(resolvedFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	logFile = file

	return nil
}

func GenerateRandomID() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result string
	for i := 0; i < 10; i++ {
		result += string(charset[rand.Intn(len(charset))])
	}

	ms := fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond))

	ms = ms[len(ms)-6:]

	result += time.Now().Format("20060102150405") + ms

	return result
}

func PrintLog(detail string) {
	os := runtime.GOOS

	switch os {
	case "windows":
		if err := InitLogFileWin(); err != nil {
			fmt.Println("Error initializing log file:", err)
			return
		}
	case "linux":
		if err := InitLogFileLin(); err != nil {
			fmt.Println("Error initializing log file:", err)
			return
		}
	}

	if logFile != nil {
		defer logFile.Sync() // Ensure logs are written before program exit
		randomID := GenerateRandomID()
		pc, file, line, _ := runtime.Caller(1)
		funcName := runtime.FuncForPC(pc).Name()

		// Extract only function name without package path
		lastSlashIndex := strings.LastIndex(funcName, "/")
		if lastSlashIndex >= 0 {
			funcName = funcName[lastSlashIndex+1:]
		}

		_, fileName := filepath.Split(file)
		logFile.WriteString(fmt.Sprintf("%s:%d:%s:%s: %s\n", strings.ToUpper(fileName), line, randomID, strings.ToUpper(funcName), detail))
	} else {
		fmt.Println("Log file is not initialized")
	}

}
