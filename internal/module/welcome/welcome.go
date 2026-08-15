// Package welcome sends a greeting message when a new member joins the server.
package welcome

import (
	"fmt"
	"math/rand/v2"
	"strings"

	"kai-bot/internal/discordutil"

	"github.com/bwmarrin/discordgo"
)

// defaultMessages is the built-in list of welcome messages, picked at random.
const defaultMessages = `Welcome to {server}, {user}! We're glad to have you here 🎉 | Hey {user}, welcome aboard! Make yourself at home 👋 | A new member has arrived! Welcome, {user} 🚀 | Welcome, {user}! We hope you enjoy your time with us 😊 | Glad you're here, {user}! Welcome to the community 🎊 | Hello {user}! Welcome to {server} 👋 | Welcome aboard, {user}! Your adventure starts here 🚀 | Everyone say hello to {user}, our newest member! 🎉 | Great to see you, {user}! Welcome to the server 😄 | Welcome, {user}! Don't forget to check out the server rules 📜 | Hey {user}! We've been waiting for you. Welcome! ✨ | Look who just joined! Welcome to {server}, {user} 👀 | Welcome to the family, {user}! We hope you'll feel at home here ❤️ | A wild {user} appeared! Welcome to the server 🎮 | The community just got better! Welcome, {user} 🌟 | New member unlocked: {user}! Welcome to {server} 🔓 | Welcome, {user}! Feel free to explore and join the conversation 💬 | Hello and welcome, {user}! We're happy to have you with us 🤝 | Make some noise for {user}, our newest member! 🎉 | Welcome to {server}, {user}! Please read the rules and enjoy your stay 📚 | Hey {user}, you made it! Welcome aboard 🥳 | Another awesome person joined us! Welcome, {user} ⭐ | Welcome, {user}! Introduce yourself when you're ready 👋 | The gates are open! Welcome to {server}, {user} 🏰 | Kai welcomes you to {server}, {user}! Have a great time 🤖`

// Config holds the welcome module settings.
type Config struct{}

// Module sends a welcome message when a new member joins.
type Module struct {
	messages []string
}

// New creates a welcome module with the built-in messages.
func New(cfg Config) *Module {
	return &Module{messages: splitMessages(defaultMessages)}
}

// Name returns the module name.
func (m *Module) Name() string { return "welcome" }

// Register attaches the guild member add handler.
func (m *Module) Register(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
		if e.Member == nil || e.Member.User == nil || isBot(e.Member.User) {
			return
		}

		guild, err := s.Guild(e.Member.GuildID)
		if err != nil {
			fmt.Printf("welcome: failed to get guild: %v\n", err)
			return
		}
		if guild == nil {
			fmt.Printf("welcome: guild %s not found\n", e.Member.GuildID)
			return
		}

		channelID, err := discordutil.FindAnnounceChannel(s, e.Member.GuildID)
		if err != nil {
			fmt.Printf("welcome: failed to find channel: %v\n", err)
			return
		}
		if channelID == "" {
			fmt.Printf("welcome: no writable text channel found for guild %s\n", e.Member.GuildID)
			return
		}

		content := formatGreeting(e.Member.User, guild, m.messages)
		if _, err := s.ChannelMessageSend(channelID, content); err != nil {
			fmt.Printf("welcome: failed to send message: %v\n", err)
			return
		}

		fmt.Printf("welcome: greeted %s\n", e.Member.User.Username)
	})
}

// isBot reports whether the given user is a bot account.
func isBot(u *discordgo.User) bool {
	return u != nil && u.Bot
}

// formatGreeting builds the welcome message content, mentioning the user.
func formatGreeting(user *discordgo.User, guild *discordgo.Guild, messages []string) string {
	msg := pickMessage(messages)
	msg = strings.ReplaceAll(msg, "{user}", user.Username)
	msg = strings.ReplaceAll(msg, "{server}", guild.Name)
	return fmt.Sprintf("👋 <@%s> — %s", user.ID, msg)
}

// splitMessages splits a pipe-separated string into individual messages.
func splitMessages(input string) []string {
	parts := strings.Split(input, "|")
	var out []string
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

// pickMessage returns a random message from the list.
func pickMessage(messages []string) string {
	return messages[rand.IntN(len(messages))]
}
