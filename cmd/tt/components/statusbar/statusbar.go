package statusbar

import (
	"github.com/alswl/go-toodledo/cmd/tt/styles"
	"github.com/charmbracelet/lipgloss"
	"github.com/knipferrc/teacup/statusbar"
)

type Model struct {
	statusbar.Bubble
}

func NewDefault() Model {
	sb := statusbar.New(
		statusbar.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Light: styles.White, Dark: styles.White},
			Background: lipgloss.AdaptiveColor{Light: styles.Pink, Dark: styles.Pink},
		},
		statusbar.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Light: styles.White, Dark: styles.White},
			Background: lipgloss.AdaptiveColor{Light: styles.DarkGray, Dark: styles.DarkGray},
		},
		statusbar.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Light: styles.White, Dark: styles.White},
			Background: lipgloss.AdaptiveColor{Light: styles.DarkPurple, Dark: styles.DarkPurple},
		},
		statusbar.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Light: styles.White, Dark: styles.White},
			Background: lipgloss.AdaptiveColor{Light: styles.DarkPurple, Dark: styles.DarkPurple},
		},
	)
	return Model{sb}
}
