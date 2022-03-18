package events

import (
	"github.com/DisgoOrg/disgo/discord"
	"github.com/DisgoOrg/snowflake"
)

type WebhooksUpdateEvent struct {
	*GenericEvent
	GuildId   snowflake.Snowflake
	ChannelID snowflake.Snowflake
}

// Guild returns the Guild the webhook was updated in.
// This will only check cached guilds!
func (e *WebhooksUpdateEvent) Guild() (discord.Guild, bool) {
	return e.Bot().Caches().Guilds().Get(e.GuildId)
}

func (e *WebhooksUpdateEvent) Channel() (discord.GuildMessageChannel, bool) {
	return e.Bot().Caches().Channels().GetGuildMessageChannel(e.ChannelID)
}
