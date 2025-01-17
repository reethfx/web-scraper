package crawler

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

type SensitiveData struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

func ParseToJSON(body io.Reader) ([]SensitiveData, error) {
	scanner := bufio.NewScanner(body)
	scanner.Split(bufio.ScanLines)

	var results []SensitiveData

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		// Data scrapping method
		if emails := findEmails(line); len(emails) > 0 {
			for _, email := range emails {
				results = append(results, SensitiveData{
					Type:    "Email",
					Content: email,
				})
			}
		}

		if keys := findAPIKeys(line); len(keys) > 0 {
			for _, key := range keys {
				results = append(results, SensitiveData{
					Type:    "API Key",
					Content: key,
				})
			}
		}

		if paths := findSensitivePaths(line); len(paths) > 0 {
			for _, path := range paths {
				results = append(results, SensitiveData{
					Type:    "Sensitive Path",
					Content: path,
				})
			}
		}
	}

	return results, nil
}

// Mail regex
func findEmails(text string) []string {
	regex := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	return regex.FindAllString(text, -1)
}

// API keys regex
func findAPIKeys(text string) []string {
	regex := regexp.MustCompile(`(?i)(api[_-]?key|token|bearer)\s*[:=]\s*["']?([a-zA-Z0-9_\-]+)["']?`)
	matches := regex.FindAllStringSubmatch(text, -1)

	var keys []string
	for _, match := range matches {
		if len(match) > 2 {
			keys = append(keys, match[2])
		}
	}
	return keys
}

// Sensible paths regex
func findSensitivePaths(text string) []string {
	regex := regexp.MustCompile(`(?i)/?(config|backup|admin|db)/?`)
	return regex.FindAllString(text, -1)
}
