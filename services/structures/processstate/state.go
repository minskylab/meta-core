package processstate

type ProcessState string

const (
	PENDING      ProcessState = "PENDING"
	PROVISIONING              = "PROVISIONING"
	DEPLOYING                 = "DEPLOYING"
	RUNNING                   = "RUNNING"
	IDLE                      = "IDLE"
	STOPPING                  = "STOPPING"
	EXIT_TIMEOUT              = "EXIT_TIMEOUT"
	EXIT_SUCCESS              = "EXIT_SUCCESS"
)

func (ps ProcessState) String() string {
	switch ps {
	case PENDING:
		return string(ps)
	case PROVISIONING:
		return string(ps)
	case DEPLOYING:
		return string(ps)
	case RUNNING:
		return string(ps)
	case IDLE:
		return string(ps)
	case STOPPING:
		return string(ps)
	case EXIT_TIMEOUT:
		return string(ps)
	case EXIT_SUCCESS:
		return string(ps)
	default:
		return "undefined"
	}
}
