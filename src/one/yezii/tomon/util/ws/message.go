package ws

type MessageForSend struct {
	Op int         `json:"op"`
	D  interface{} `json:"d"`
}
type tokenMap struct {
	Token string `json:"token"`
}

func HeartbeatMessageForSend() MessageForSend {
	return MessageForSend{
		Op: 1,
	}
}
func AuthMessageForSend(token string) MessageForSend {
	return MessageForSend{
		Op: 2,
		D:  tokenMap{Token: token},
	}
}
