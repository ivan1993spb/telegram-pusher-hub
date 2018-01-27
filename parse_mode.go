package main

type ParseMode uint8

const (
	ParseModeEmpty = iota
	ParseModeMarkdown
	ParseModeHTML
)

const (
	ParseModeEmptyString    = ""
	ParseModeMarkdownString = "Markdown"
	ParseModeHTMLString     = "HTML"
)

func GetParseMode(parseMode string) ParseMode {
	if parseMode == ParseModeEmptyString {
		return ParseModeEmpty
	}
	if parseMode == ParseModeMarkdownString {
		return ParseModeMarkdown
	}
	if parseMode == ParseModeHTMLString {
		return ParseModeHTML
	}
	return ParseModeEmpty
}

func (pm ParseMode) String() string {
	if pm == ParseModeEmpty {
		return ""
	}
	if pm == ParseModeMarkdown {
		return "Markdown"
	}
	if pm == ParseModeHTML {
		return "HTML"
	}
	return ""
}
