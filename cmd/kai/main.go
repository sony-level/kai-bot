// Package main is the entry point for the Kai Discord bot.
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"kai-bot/internal/config"
	"kai-bot/internal/discordutil"
	"kai-bot/internal/module/welcome"

	"github.com/bwmarrin/discordgo"
)

// main loads the configuration, connects Kai to Discord, registers its
// modules, and blocks until an interrupt or termination signal is received.
func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(1) // this line
	}

	dg, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		fmt.Println("error creating Discord session:", err)
		os.Exit(1)
	}

	dg.Identify.Intents = discordgo.IntentsGuildMembers | discordgo.IntentsGuilds

	welcome.New(welcome.Config{}).Register(dg)

	dg.AddHandler(func(s *discordgo.Session, e *discordgo.GuildCreate) {
		if e.Unavailable {
			return
		}
		announceOnline(s, e.Guild.ID)
	})

	if err := dg.Open(); err != nil {
		fmt.Println("error opening connection:", err)
		os.Exit(1)
	}
	defer dg.Close()

	fmt.Println("Kai is running. Press CTRL+C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc
}

// announceOnline sends an "online" message in the first writable text
// channel it finds in the given guild, so Kai announces itself in every
// server it is installed on without requiring a fixed channel ID.
func announceOnline(s *discordgo.Session, guildID string) {
	channelID, err := discordutil.FindAnnounceChannel(s, guildID)
	if err != nil {
		fmt.Println("error finding announce channel:", err)
		return
	}
	if channelID == "" {
		fmt.Println("no writable text channel found for guild", guildID)
		return
	}

	if _, err := s.ChannelMessageSend(channelID, "✅ Kai is now online!"); err != nil {
		fmt.Println("error sending online message:", err)
	}
}
