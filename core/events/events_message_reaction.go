package events

import (
	"github.com/DisgoOrg/disgo/discord"
	"github.com/DisgoOrg/snowflake"
)

// GenericReactionEvent is called upon receiving MessageReactionAddEvent or MessageReactionRemoveEvent
type GenericReactionEvent struct {
	*GenericEvent
	UserID    snowflake.Snowflake
	ChannelID snowflake.Snowflake
	MessageID snowflake.Snowflake
	GuildID   *snowflake.Snowflake
	Emoji     discord.ReactionEmoji
}

func (e *GenericReactionEvent) User() (discord.User, bool) {
	return e.Bot().Caches().Users().Get(e.UserID)
}

// MessageReactionAddEvent indicates that a core.User added a discord.MessageReaction to a core.Message in a core.Channel(this+++ requires the discord.GatewayIntentGuildMessageReactions and/or discord.GatewayIntentDirectMessageReactions)
type MessageReactionAddEvent struct {
	*GenericReactionEvent
	Member discord.Member
}

// MessageReactionRemoveEvent indicates that a core.User removed a discord.MessageReaction from a core.Message in a core.GetChannel(requires the discord.GatewayIntentGuildMessageReactions and/or discord.GatewayIntentDirectMessageReactions)
type MessageReactionRemoveEvent struct {
	*GenericReactionEvent
}

// MessageReactionRemoveEmojiEvent indicates someone removed all discord.MessageReaction of a specific core.Emoji from a core.Message in a core.Channel(requires the discord.GatewayIntentGuildMessageReactions and/or discord.GatewayIntentDirectMessageReactions)
type MessageReactionRemoveEmojiEvent struct {
	*GenericEvent
	ChannelID snowflake.Snowflake
	MessageID snowflake.Snowflake
	GuildID   *snowflake.Snowflake
	Emoji     discord.ReactionEmoji
}

// MessageReactionRemoveAllEvent indicates someone removed all discord.MessageReaction(s) from a core.Message in a core.Channel(requires the discord.GatewayIntentGuildMessageReactions and/or discord.GatewayIntentDirectMessageReactions)
type MessageReactionRemoveAllEvent struct {
	*GenericEvent
	ChannelID snowflake.Snowflake
	MessageID snowflake.Snowflake
	GuildID   *snowflake.Snowflake
}
