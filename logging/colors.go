package logging

type TextModifier string

// Text reset / modifiers
const (
	Reset         TextModifier = "\033[0m" // reset all styles
	Bold          TextModifier = "\033[1m"
	Dim           TextModifier = "\033[2m"
	Italic        TextModifier = "\033[3m"
	Underline     TextModifier = "\033[4m"
	Blink         TextModifier = "\033[5m"
	Reverse       TextModifier = "\033[7m"
	Hidden        TextModifier = "\033[8m"
	Strikethrough TextModifier = "\033[9m"
)

// Standard foreground colors
const (
	Black   TextModifier = "\033[30m"
	Red     TextModifier = "\033[31m"
	Green   TextModifier = "\033[32m"
	Yellow  TextModifier = "\033[33m"
	Blue    TextModifier = "\033[34m"
	Magenta TextModifier = "\033[35m"
	Cyan    TextModifier = "\033[36m"
	White   TextModifier = "\033[37m"
	// Bright variants
	BrightBlack   TextModifier = "\033[90m"
	BrightRed     TextModifier = "\033[91m"
	BrightGreen   TextModifier = "\033[92m"
	BrightYellow  TextModifier = "\033[93m"
	BrightBlue    TextModifier = "\033[94m"
	BrightMagenta TextModifier = "\033[95m"
	BrightCyan    TextModifier = "\033[96m"
	BrightWhite   TextModifier = "\033[97m"
)

// Standard background colors
const (
	BlackBG   TextModifier = "\033[40m"
	RedBG     TextModifier = "\033[41m"
	GreenBG   TextModifier = "\033[42m"
	YellowBG  TextModifier = "\033[43m"
	BlueBG    TextModifier = "\033[44m"
	MagentaBG TextModifier = "\033[45m"
	CyanBG    TextModifier = "\033[46m"
	WhiteBG   TextModifier = "\033[47m"
	// Bright backgrounds
	BrightBlackBG   TextModifier = "\033[100m"
	BrightRedBG     TextModifier = "\033[101m"
	BrightGreenBG   TextModifier = "\033[102m"
	BrightYellowBG  TextModifier = "\033[103m"
	BrightBlueBG    TextModifier = "\033[104m"
	BrightMagentaBG TextModifier = "\033[105m"
	BrightCyanBG    TextModifier = "\033[106m"
	BrightWhiteBG   TextModifier = "\033[107m"
)
