//go:build all
// +build all

package tests

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

// TestChecklist tracks test results with visual checklist
type TestChecklist struct {
	items []ChecklistItem
}

// ChecklistItem represents an individual test result
type ChecklistItem struct {
	Name     string
	Status   bool // true = passed, false = failed
	Duration time.Duration
}

// NewTestChecklist creates a new test checklist
func NewTestChecklist() *TestChecklist {
	return &TestChecklist{
		items: make([]ChecklistItem, 0),
	}
}

// AddResult adds a test result to the checklist
func (tc *TestChecklist) AddResult(name string, status bool, duration time.Duration) {
	tc.items = append(tc.items, ChecklistItem{
		Name:     name,
		Status:   status,
		Duration: duration,
	})
}

// PrintChecklist prints a formatted checklist with visual indicators
func (tc *TestChecklist) PrintChecklist(t *testing.T) {
	t.Log("\n" + strings.Repeat("═", 80))
	t.Log("📋 TEST CHECKLIST")
	t.Log(strings.Repeat("═", 80))

	passed := 0
	failed := 0
	totalDuration := time.Duration(0)

	// Find the maximum name length for alignment
	maxNameLength := 0
	for _, item := range tc.items {
		if len(item.Name) > maxNameLength {
			maxNameLength = len(item.Name)
		}
	}

	for _, item := range tc.items {
		totalDuration += item.Duration

		var statusSymbol string
		var durationColor string

		if item.Status {
			statusSymbol = "✅"
			durationColor = "\033[32m" // Green for passed
			passed++
		} else {
			statusSymbol = "❌"
			durationColor = "\033[31m" // Red for failed
			failed++
		}

		// Format duration with color coding
		durationStr := fmt.Sprintf("(%v)", item.Duration)
		if item.Duration > time.Second {
			durationStr += " ⏰"
		}

		// Create aligned output
		padding := strings.Repeat(" ", maxNameLength-len(item.Name)+2)
		checklistItem := fmt.Sprintf("   %s %s%s%s%s", statusSymbol, item.Name, padding, durationColor, durationStr)

		// Reset color after duration
		t.Log(checklistItem + "\033[0m")
	}

	t.Log(strings.Repeat("─", 80))

	// Summary with color coding
	if failed > 0 {
		t.Logf("📊 RESULTS: \033[32m%d PASSED\033[0m, \033[31m%d FAILED\033[0m", passed, failed)
		t.Logf("⏱️  TOTAL TIME: %v", totalDuration)
		t.Logf("🚨 FAILURE RATE: \033[31m%.1f%%\033[0m", float64(failed)/float64(len(tc.items))*100)
	} else {
		t.Logf("📊 RESULTS: \033[32m%d PASSED\033[0m, %d FAILED", passed, failed)
		t.Logf("⏱️  TOTAL TIME: %v", totalDuration)
		t.Logf("🎉 SUCCESS RATE: \033[32m100%%\033[0m")
	}

	t.Log(strings.Repeat("═", 80))

	// Highlight failed tests
	if failed > 0 {
		t.Log("\n🚨 FAILED TESTS SUMMARY:")
		for _, item := range tc.items {
			if !item.Status {
				t.Logf("   ❌ \033[31m%s - requires attention\033[0m", item.Name)
			}
		}
	}
}