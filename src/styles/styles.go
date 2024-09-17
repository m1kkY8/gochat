package styles

import (
	"math/rand/v2"
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	Border      lipgloss.Style
	SenderStyle lipgloss.Style
}

func GenerateRandomANSIColor() string {
	return strconv.Itoa(rand.IntN(256))
}

func DefaultStyle(userColor string) *Styles {
	border := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("63"))

	senderStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(userColor))

	return &Styles{
		Border:      border,
		SenderStyle: senderStyle,
	}
}
