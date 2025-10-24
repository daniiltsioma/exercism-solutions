package secret

var actions = map[uint]string{
	1: "wink",
	2: "double blink",
	4: "close your eyes",
	8: "jump",
}

func Handshake(code uint) []string {
	handshake := []string{}

	var order []uint
	if code & 16 == 16 {
		order = []uint{8,4,2,1}
	} else {
		order = []uint{1,2,4,8}
	}

	for _, v := range order {
		if code & v == v {
			handshake = append(handshake, actions[v])
		}
	}

	return handshake
}
