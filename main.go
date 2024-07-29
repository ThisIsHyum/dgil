package dgil

import (
	"github.com/bwmarrin/discordgo"
)

type InviteLogger struct {
	session       *discordgo.Session
	cachedInvites map[string][]*discordgo.Invite
}

func NewInviteLogger(session *discordgo.Session) InviteLogger {
	inviteLogger := InviteLogger{
		session:       session,
		cachedInvites: map[string][]*discordgo.Invite{},
	}

	inviteLogger.AddGuildsToCache()
	session.AddHandler(inviteLogger.CreateInviteHandler)
	return inviteLogger
}

func (il *InviteLogger) AddGuildsToCache() {
	guilds := il.session.State.Guilds
	for _, guild := range guilds {
		invites, _ := il.session.GuildInvites(guild.ID)
		il.cachedInvites[guild.ID] = invites
	}
}
func (il *InviteLogger) GetInvite(member *discordgo.Member) *discordgo.Invite {
	cachedInvites := il.cachedInvites[member.GuildID]
	invites, _ := il.session.GuildInvites(member.GuildID)
	for _, invite := range invites {
		for _, cachedInvite := range cachedInvites {
			if (cachedInvite.Code == invite.Code) && (invite.Uses > cachedInvite.Uses) {
				il.cachedInvites[member.GuildID] = invites
				return invite
			}
		}
	}
	return &discordgo.Invite{}
}

func (il *InviteLogger) CreateInviteHandler(s *discordgo.Session, i *discordgo.InviteCreate) {
	invites, _ := il.session.GuildInvites(i.GuildID)
	il.cachedInvites[i.GuildID] = invites
}
