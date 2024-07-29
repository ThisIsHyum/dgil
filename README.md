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

Create new object inviteLogger with session

```go
session, err := discordgo.New("Bot " + "your token here")
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
