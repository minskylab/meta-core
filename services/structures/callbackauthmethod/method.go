package callbackauthmethod

type CallbackAuthMethod string

const (
	NONE       CallbackAuthMethod = "NONE"
	BASIC_AUTH                    = "BASIC_AUTH"
)

func (c CallbackAuthMethod) String() string {
	switch c {
	case NONE:
		return string(c)
	case BASIC_AUTH:
		return string(c)
	default:
		return "undefined"
	}
}
