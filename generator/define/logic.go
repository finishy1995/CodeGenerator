package define

type LogicKey interface {
	// Exec step and generate output
	Exec(step Step, args ...string)
	// GetKey get key name
	GetKey() string
}

var (
	LogicKeyMap = make(map[string]LogicKey)
)

func Register(keys ...LogicKey) {
	for _, key := range keys {
		if key != nil {
			LogicKeyMap[key.GetKey()] = key
		}
	}
}
