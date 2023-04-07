package types

type ChannelSub interface {
	GetArg() map[string]string
	ReceiveData([]byte)
}

