package main

import (
	"golang.design/x/clipboard"
	"strings"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	copiedText := clipboard.Read(clipboard.FmtText)

	changelog := string(copiedText)
	changelog = strings.TrimSpace(changelog)
	splitted := strings.Split(changelog, "\n")

	var sb strings.Builder
	for i, s := range splitted {
		s := strings.TrimSpace(s)
		if s != "" {
			if i != 0 {
				sb.WriteString("\n")
			}
			sb.WriteString("* ")
			sb.WriteString(s)
		}
	}

	clipboard.Write(clipboard.FmtText, []byte(sb.String()))

}
