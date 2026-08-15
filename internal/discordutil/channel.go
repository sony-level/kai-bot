// Package discordutil provides small Discord helpers shared across modules.
package discordutil

import (
	"sort"

	"github.com/bwmarrin/discordgo"
)

// FindAnnounceChannel returns a channel Kai can write to in the given guild:
// the guild's system channel if set, otherwise the first text channel
// (ordered by position). It returns an empty string if none is found.
func FindAnnounceChannel(s *discordgo.Session, guildID string) (string, error) {
	guild, err := s.Guild(guildID)
	if err == nil && guild.SystemChannelID != "" {
		return guild.SystemChannelID, nil
	}

	channels, err := s.GuildChannels(guildID)
	if err != nil {
		return "", err
	}

	sort.Slice(channels, func(i, j int) bool {
		return channels[i].Position < channels[j].Position
	})

	for _, c := range channels {
		if c.Type == discordgo.ChannelTypeGuildText {
			return c.ID, nil
		}
	}

	return "", nil
}
