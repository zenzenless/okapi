package rest

type Signer interface {
	Sign([]byte)string
}

type HamcSigner struct {
	SecretKey string
}


// sign