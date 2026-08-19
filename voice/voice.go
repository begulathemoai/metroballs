/// this is my manifesto
// basically you can't get voice support without doing a bunch of bullshit
// cuz this bot is built with discordgo v0.29.0 which is old
// so no DAVE support. but ! there's a fork that adds DAVE support
// and it's based on master. which is new. so either you somehow downgrade the fork to v0.29.0,
// or you deal with the fact that a bunch of method signatures changed, which means you have to
// at least touch up a big part of the bot (and i'm trying to change as little of the internal code as possible);
// or you implement DAVE yourself but you would have to
// play around with the internals of discordgo which means doing a fork of v0.29.0 just for metrobot
//
// i'm shit enough in go to understand what the solution is gonna be

package voice

import (
	"encoding/json"
	"fmt"

	"github.com/begulathemoai/metroballs/config"
	"github.com/bwmarrin/discordgo"
)

type Voice struct {
	Session    *discordgo.Session
	Config     *config.Config
	Connection *discordgo.VoiceConnection
}

func New(sesh *discordgo.Session, conf *config.Config) Voice {
	return Voice{
		Session:    sesh,
		Config:     conf,
		Connection: nil,
	}
}

// there must be a function somewhere in discordgo that does this but i haven't found it
func getVoiceState(s *discordgo.Session, gID string, uID string) (j map[string](any), err error) {
	out, err := s.Request("GET", "https://discord.com/api/v10/guilds/"+gID+"/voice-states/"+uID, nil)
	if err != nil {
		return nil, fmt.Errorf("while requesting voice state: %w", err)
	}
	var how_to_sleep map[string](any) // https://youtu.be/f5wrW7gv7b8
	err = json.Unmarshal(out, &how_to_sleep)
	if err != nil {
		return nil, fmt.Errorf("while unmarshaling received voice state: %w", err)
	}
	return how_to_sleep, nil

}

// stupid name
// i keep getting 401's :sob:
func (v *Voice) DoConnect(s *discordgo.Session, i *discordgo.InteractionCreate, callerID string) (err error) {
	user_voice_state, err := getVoiceState(s, i.GuildID, i.Member.User.ID)
	if err != nil {

		return fmt.Errorf("while getting voice state: %w", err)
	}
	voice_channel_id, ok := user_voice_state["channel_id"].(string)
	if ok {
		err := v.Connect(voice_channel_id, false, true)
		if err != nil {
			return fmt.Errorf("while connecting to voice: %w", err)
		}
	} else {
		return fmt.Errorf("returned \"channel_id\" is not a string")
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "connected (i will keep getting disconnected & reconnected as long as e2ee/dave isn't added)",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	return nil
}

func (v *Voice) Connect(id string, muted bool, deafened bool) (err error) {
	conn, err := v.Session.ChannelVoiceJoin(v.Config.DiscordGuildID, id, muted, deafened)
	if err != nil {
		v.Connection = conn
		return nil
	} else {
		return fmt.Errorf("connecting to voice channel: %w", err)
	}
}
