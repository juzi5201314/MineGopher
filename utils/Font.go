package utils

import "strings"

const (
	AnsiPre   = "\u001b["
	AnsiReset = AnsiPre + "0m"

	AnsiBold       = AnsiPre + "1m"
	AnsiItalic     = AnsiPre + "3m"
	AnsiUnderlined = AnsiPre + "4m"

	AnsiBlack   = AnsiPre + "30m"
	AnsiRed     = AnsiPre + "31m"
	AnsiGreen   = AnsiPre + "32m"
	AnsiYellow  = AnsiPre + "33m"
	AnsiBlue    = AnsiPre + "34m"
	AnsiMagenta = AnsiPre + "35m"
	AnsiCyan    = AnsiPre + "36m"
	AnsiWhite   = AnsiPre + "37m"
	AnsiGray    = AnsiPre + "30;1m"

	AnsiBrightRed     = AnsiPre + "31;1m"
	AnsiBrightGreen   = AnsiPre + "32;1m"
	AnsiBrightYellow  = AnsiPre + "33;1m"
	AnsiBrightBlue    = AnsiPre + "34;1m"
	AnsiBrightMagenta = AnsiPre + "35;1m"
	AnsiBrightCyan    = AnsiPre + "36;1m"
	AnsiBrightWhite   = AnsiPre + "37;1m"

	Pre = "§"

	Black      = Pre + "0"
	Blue       = Pre + "1"
	Green      = Pre + "2"
	Cyan       = Pre + "3"
	Red        = Pre + "4"
	Magenta    = Pre + "5"
	Orange     = Pre + "6"
	BrightGray = Pre + "7"
	Gray       = Pre + "8"
	BrightBlue = Pre + "9"

	BrightGreen   = Pre + "a"
	BrightCyan    = Pre + "b"
	BrightRed     = Pre + "c"
	BrightMagenta = Pre + "d"
	Yellow        = Pre + "e"
	White         = Pre + "f"

	Obfuscated    = Pre + "k"
	Bold          = Pre + "l"
	StrikeThrough = Pre + "m"
	Underlined    = Pre + "n"
	Italic        = Pre + "o"

	Reset = Pre + "r"
)

type ColorString string

var colorConvert = map[string]string{
	Black:         AnsiBlack,
	Blue:          AnsiBlue,
	Green:         AnsiGreen,
	Cyan:          AnsiCyan,
	Red:           AnsiRed,
	Magenta:       AnsiMagenta,
	Orange:        AnsiYellow,
	BrightGray:    AnsiWhite,
	Gray:          AnsiGray,
	BrightBlue:    AnsiBrightBlue,
	BrightGreen:   AnsiBrightGreen,
	BrightCyan:    AnsiBrightCyan,
	BrightRed:     AnsiBrightRed,
	BrightMagenta: AnsiBrightMagenta,
	Yellow:        AnsiBrightYellow,
	White:         AnsiBrightWhite,
	Bold:          AnsiBold,
	Underlined:    AnsiUnderlined,
	Italic:        AnsiItalic,
	Reset:         AnsiReset,
	StrikeThrough: AnsiUnderlined,
	Obfuscated:    AnsiUnderlined,
}

func (str ColorString) ToANSI() string {
	text := string(str)
	for toConvert, convertValue := range colorConvert {
		text = strings.Replace(text, toConvert, convertValue, -1)
	}
	return text
}

func (str ColorString) ToMinecraft() string {
	text := string(str)
	for convertValue, toConvert := range colorConvert {
		text = strings.Replace(text, toConvert, convertValue, -1)
	}
	return text
}

func (str ColorString) StripMinecraft() string {
	text := string(str)
	for toConvert := range colorConvert {
		text = strings.Replace(text, toConvert, "", -1)
	}
	return text
}

func (str ColorString) StripANSI() string {
	text := string(str)
	for _, toConvert := range colorConvert {
		text = strings.Replace(text, toConvert, "", -1)
	}
	return text
}

func (str ColorString) StripAll() string {
	text := string(str)
	for mcpeColor, ansiColor := range colorConvert {
		text = strings.Replace(text, mcpeColor, "", -1)
		text = strings.Replace(text, ansiColor, "", -1)
	}
	return text
}
