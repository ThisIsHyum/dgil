# dgil
**d**iscord**go** **i**nvite **l**ogger

## Installation
``` sh
go get github.com/ThisIsHyum/dgil
```

## Usage
Import packages dgil and discordgo and into your project.

```go
import (
	"github.com/ThisIsHyum/dgil"
	"github.com/bwmarrin/discordgo"
)
```

open session and create object inviteLogger

```go
session, err := discordgo.New("Bot " + "your token here")
sesion.Open()
inviteLogger = dgil.NewInviteLogger(session)
```

Create join handler and handle it
```go
session.AddHandler(JoinHandler)

func JoinHandler(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
	invite := inviteLogger.GetInvite(e.Member)
	/*
	Your code
	*/
}
