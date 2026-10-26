package svgterm

import "strings"

// Palette is named after Xcode's syntax roles rather than ANSI colours, since
// that is where the values come from: the portfolio shows its Swift snippets in
// Xcode's Default theme, and the terminal on the profile should look like it
// came off the same machine.
type Palette struct {
	Name       string
	Background string
	TitleBar   string
	Border     string
	Foreground string
	Dim        string // comments

	Keyword  string // the prompt
	Function string // commands
	String   string // URLs and paths
	Member   string // stack, project and repo names
	Type     string // the role
	Call     string // latency and views
	Number   string
	Flag     string // preprocessor colour: flags, and the stale marker

	// Xcode has no green or red to borrow, so these are the system colours the
	// site uses for the same meaning.
	Success string
	Failure string
}

// Xcode's "Default (Dark)".
var dark = Palette{
	Name:       "dark",
	Background: "#1f1f24",
	TitleBar:   "#2c2c31",
	Border:     "#3a3a40",
	Foreground: "#dfdfe0",
	Dim:        "#7f8c98",
	Keyword:    "#ff7ab2",
	Function:   "#67b7a4",
	String:     "#ff8170",
	Member:     "#78c2b3",
	Type:       "#dabaff",
	Call:       "#b281eb",
	Number:     "#d9c97c",
	Flag:       "#ffa14f",
	Success:    "#30d158",
	Failure:    "#ff453a",
}

// Xcode's "Default (Light)". The greens and reds are Apple's accessible
// variants, the plain ones being too faint on white.
var light = Palette{
	Name:       "light",
	Background: "#ffffff",
	TitleBar:   "#ececee",
	Border:     "#c6c6c8",
	Foreground: "#262626",
	Dim:        "#5d6c79",
	Keyword:    "#9b2393",
	Function:   "#326d74",
	String:     "#c41a16",
	Member:     "#0f68a0",
	Type:       "#3900a0",
	Call:       "#6c36a9",
	Number:     "#1c00cf",
	Flag:       "#643820",
	Success:    "#248a3d",
	Failure:    "#d70015",
}

// ThemeByName resolves the ?theme= parameter, falling back to dark.
func ThemeByName(name string) Palette {
	if strings.EqualFold(strings.TrimSpace(name), "light") {
		return light
	}
	return dark
}
