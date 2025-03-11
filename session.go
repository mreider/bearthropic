package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/matthewreider/bthropic/claude"
)

type Session struct {
	claude *claude.Client
	// noteID removed
	reader *bufio.Reader
}

func NewSession() *Session {
	// Load config
	data, err := os.ReadFile(getConfigPath())
	if err != nil {
		panic(err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		panic(err)
	}

	return &Session{
		claude: claude.NewClient(config.ClaudeAPIKey),
		reader: bufio.NewReader(os.Stdin),
	}
}

func (s *Session) Start() error {
	fmt.Println("Starting new session with Claude.")

	var question string
	var response string

	for {
		// Get the initial question or modification prompt
		if question == "" {
			fmt.Print("Enter your question: ")
		} else {
			fmt.Print("Enter modifications (or 'end' to finish): ")
		}
		question, _ = s.reader.ReadString('\n')
		question = strings.TrimSpace(question)

		// Check if the user wants to end the session
		if strings.ToLower(question) == "end" {
			break
		}

		// Re-initialize s.claude at the beginning of each loop
		data, err := os.ReadFile(getConfigPath())
		if err != nil {
			return err
		}

		var config Config
		if err := json.Unmarshal(data, &config); err != nil {
			return err
		}

		s.claude = claude.NewClient(config.ClaudeAPIKey)

		// Show "working" animation
		stopChan := make(chan bool)
		go func() {
			chars := []string{"/", "-", "\\", "|"}
			for i := 0; ; i++ {
				fmt.Printf("\rWorking: %s", chars[i%len(chars)])
				time.Sleep(100 * time.Millisecond)

				// Check if the stop signal has been received
				select {
				case <-stopChan:
					return
				default:
					// Continue animation
				}
			}
		}()

		response, err = s.claude.CreateMessage(question)
		// Stop the animation
		stopChan <- true
		close(stopChan)
		fmt.Print("\r")

		if err != nil {
			return err
		}

		// Process the response to fix date tag formatting.
		response = processResponse(response)

		// Always create a new note; previous note trashing is removed.
		if err := s.createBearNote(response); err != nil {
			return err
		}

		fmt.Println("Note created. Review it; Ctrl-C or type 'end' to finish, or type changes:")
		question = "" // Reset question to prompt for modifications
	}

	return nil
}

func processResponse(resp string) string {
	// Remove <sup> tags if found.
	resp = strings.ReplaceAll(resp, "<sup>", "")
	resp = strings.ReplaceAll(resp, "</sup>", "")
	// Split into lines and if the first line contains an inline date tag, move the date to a new line.
	lines := strings.Split(resp, "\n")
	if len(lines) > 0 {
		// Look for " #20" in the first line as a crude indicator.
		if idx := strings.Index(lines[0], " #20"); idx != -1 {
			titlePart := strings.TrimSpace(lines[0][:idx])
			datePart := strings.TrimSpace(lines[0][idx+1:])
			// If datePart starts with '#' and appears of correct length (e.g., "#2022-12-20")
			if strings.HasPrefix(datePart, "#") && len(datePart) >= 11 {
				// Replace first line with title and then insert date on next line.
				lines[0] = titlePart
				newLines := []string{lines[0], datePart}
				if len(lines) > 1 {
					newLines = append(newLines, lines[1:]...)
				}
				resp = strings.Join(newLines, "\n")
			}
		}
	}
	return resp
}

func (s *Session) createBearNote(response string) error {
	// Write the response to clipboard.
	if err := clipboard.WriteAll(response); err != nil {
		return fmt.Errorf("failed to write to clipboard: %w", err)
	}

	// Use Bear's create action.
	bearURLCreate := "bear://x-callback-url/create?clipboard=yes"
	if err := exec.Command("open", bearURLCreate).Run(); err != nil {
		return fmt.Errorf("failed to create Bear note: %w", err)
	}

	return nil
}

func openURL(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
	return cmd.Run()
}
