// Package module defines the Module contract used by Kai features.
package module

import "github.com/bwmarrin/discordgo"

// Module is a Kai feature that can register Discord event handlers.
type Module interface {
	Name() string
	Register(s *discordgo.Session)
}
