package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"git-basics/internal/health"
)

func main() {
	// 1. просим git выдать историю: каждая строка "дата|сообщение"
	out, err := exec.Command("git", "log", "--pretty=format:%ad|%s", "--date=short").Output()
	if err != nil {
		fmt.Println("не удалось прочитать git log:", err)
		return
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")

	var first, last time.Time
	total := 0

	for _, line := range lines {
		parts := strings.SplitN(line, "|", 2)
		if len(parts) != 2 {
			continue
		}
		date, _ := time.Parse("2006-01-02", parts[0])
		if first.IsZero() || date.Before(first) {
			first = date
		}
		if date.After(last) {
			last = date
		}

		total += health.ScoreMessage(parts[1])
	}

	commits := len(lines)
	days := int(last.Sub(first).Hours() / 24)

	if days < 1 {
		days = 1
	}

	cadence := float64(commits) / float64(days)
	avg := float64(total) / float64(commits)

	fmt.Println("=== Repo Health ===")
	fmt.Printf("Commits: %d\n", commits)
	fmt.Printf("Span: %d days\n", days)
	fmt.Printf("Cadence: %.2f commits/day\n", cadence)
	fmt.Printf("Average message score: %.1f / 3\n", avg)
	fmt.Printf("History: %s\n", health.Label(avg))
}
