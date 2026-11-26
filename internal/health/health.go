package health

import (
	"strings"
)

var typeVerbs = []string{"add", "update", "fix", "delete", "docs", "refactor"}
var topics = []string{"readme", "branch", "merge", "conflict", "workflow", "reflection", "feature", "analyzer"}

func inList(words string, list []string) bool {
	for _, v := range list {
		if v == words {
			return true
		}
	}
	return false
}

func ScoreMessage(msg string) int {
	words := strings.Fields(strings.ToLower(msg))
	score := 0
	if len(words) >= 3 {
		score++
	}
	if len(words) > 0 && inList(strings.TrimSuffix(words[0], ":"), typeVerbs) {
		score++
	}
	for _, w := range words {
		if strings.Contains(w, ".") || inList(w, topics) {
			score++
			break
		}
	}
	return score
}

func Label(avg float64) string {
	if avg < 1.5 {
		return "messy"
	} else if avg >= 1.5 && avg < 2.5 {
		return "acceptable"
	} else {
		return "tidy"
	}
}
