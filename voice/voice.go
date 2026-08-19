package voice

import (
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

// stupid name
// i keep getting 401's :sob:
func (v *Voice) DoConnect(s *discordgo.Session, i *discordgo.InteractionCreate, callerID string) (err error) {
	out, _ := s.Client.Get("https://discord.com/api/v9/guilds/" + i.GuildID + "/voice-states/" + i.Member.User.ID)
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: fmt.Sprintf("getting channel `%v` in guild `%v` : `%v`", i.GuildID, i.Member.User.ID, out),
		}})
	/*chs, err := s.GuildChannels(i.GuildID)

	for _, ch := range chs {
		ch.Members
	}*/
	return nil

	/*fmt.Printf("starting check")
	if s.State == nil {
		fmt.Printf("What The Fuck")
		return nil
	}
	vc, err := s.State.VoiceState(i.GuildID, i.Member.User.ID)
	if err != nil || vc == nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "no",
			}})
		return err
	}

	vc_id := vc.ChannelID

	if vc_id == "" {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "you're not in a voice channel dumbass",
			}})
	} else {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: fmt.Sprintf("found you in <#%v>", vc_id),
			}})
	}
	return nil*/
	/*chs, err := s.GuildChannels(v.Config.DiscordToken)
	if err != nil {
		return fmt.Errorf("getting all channels: %w", err)
	}

	// yeah this is stupid but counterpoint i don't know go
	filtered_chs := make([]*discordgo.Channel,0,len(chs))
	for _, ch := range chs {
		if ch.Type == discordgo.ChannelTypeGuildVoice {
			filtered_chs = append(filtered_chs, ch)
		}
	}

	for _, ch := range filtered_chs {
		if ch.
	}*/
}

func (v *Voice) Connect(id int, muted bool, deafened bool) (err error) {
	conn, err := v.Session.ChannelVoiceJoin(v.Config.DiscordGuildID, fmt.Sprint(id), muted, deafened)
	if err != nil {
		v.Connection = conn
		return nil
	} else {
		return fmt.Errorf("connecting to voice channel: %w", err)
	}
}
