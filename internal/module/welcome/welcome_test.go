// Package welcome tests the welcome message formatting and bot detection.
package welcome

import (
	"testing"

	"github.com/bwmarrin/discordgo"
)

func TestFormatGreeting(t *testing.T) {
	user := &discordgo.User{ID: "123456", Username: "alice"}
	guild := &discordgo.Guild{Name: "Kai Server"}
	messages := []string{"Hello {user}", "Welcome to {server}, {user}"}
	got := formatGreeting(user, guild, messages)
	want := []string{
		"👋 <@123456> — Hello alice",
		"👋 <@123456> — Welcome to Kai Server, alice",
	}
	found := false
	for _, w := range want {
		if got == w {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("formatGreeting() = %q, want one of %v", got, want)
	}
}

func TestFormatGreetingWithServerOnly(t *testing.T) {
	user := &discordgo.User{ID: "42", Username: "bob"}
	guild := &discordgo.Guild{Name: "Test Guild"}
	messages := []string{"Welcome to {server}, {user}!"}
	got := formatGreeting(user, guild, messages)
	want := "👋 <@42> — Welcome to Test Guild, bob!"
	if got != want {
		t.Errorf("formatGreeting() = %q, want %q", got, want)
	}
}

func TestSplitMessages(t *testing.T) {
	got := splitMessages(" a | b |c  ")
	want := []string{"a", "b", "c"}
	if len(got) != len(want) {
		t.Fatalf("splitMessages() = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("splitMessages()[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestPickMessage(t *testing.T) {
	messages := []string{"a", "b", "c"}
	got := pickMessage(messages)
	found := false
	for _, m := range messages {
		if got == m {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("pickMessage() = %q, want one of %v", got, messages)
	}
}

func TestIsBot(t *testing.T) {
	if !isBot(&discordgo.User{Bot: true}) {
		t.Error("expected bot user to be detected as bot")
	}
	if isBot(&discordgo.User{Bot: false}) {
		t.Error("expected non-bot user not to be detected as bot")
	}
	if isBot(nil) {
		t.Error("expected nil user not to be detected as bot")
	}
}
