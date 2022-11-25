package define

type markHandler struct {
	prefix       string
	prefixLength int
	suffix       string
	suffixLength int
}

var (
	TplLineFeed   = "\n"
	OutLineFeed   = "\n"
	TplFileSuffix = ".tpl"

	logicMarkPrefix  = "#{" // logic mark reg format, for example #{Loop 3} => Loop 3
	logicMarkSuffix  = "}"
	logicMarkHandler *markHandler

	DictLengthMark = "*Length"
	DictKeyMark    = "*Key"
)
