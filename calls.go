package main

import "github.com/amarnathcjd/gogram/telegram"

var (
	call *Client
)

func setupCallsCore() {
	ntg := NTgCalls()
	call = ntg
}

func joinGroupCall(client *telegram.Client, ntg *Client, chatId interface{}, media MediaDescription) {
	me, _ := client.GetMe() // get the current user for JoinAs

	call, err := client.GetGroupCall(chatId) // get the group call object
	if err != nil {
		panic(err)
	}

	rawChannel, _ := client.GetSendableChannel(chatId)                                         // get the channel object
	jsonParams, err := ntg.CreateCall(rawChannel.(*telegram.InputChannelObj).ChannelID, media) // create call object with media description
	if err != nil {
		panic(err)
	}

	callResRaw, err := client.PhoneJoinGroupCall(
		&telegram.PhoneJoinGroupCallParams{
			Muted:        false,
			VideoStopped: true, // false for video call
			Call:         call,
			Params: &telegram.DataJson{
				Data: jsonParams,
			},
			JoinAs: &telegram.InputPeerUser{
				UserID:     me.ID,
				AccessHash: me.AccessHash,
			},
		},
	)

	if err != nil {
		panic(err)
	}

	callRes := callResRaw.(*telegram.UpdatesObj)
	for _, update := range callRes.Updates {
		switch u := update.(type) {
		case *telegram.UpdateGroupCallConnection: // wait for connection params
			_ = ntg.Connect(rawChannel.(*telegram.InputChannelObj).ChannelID, u.Params.Data)
		}
	}
}
