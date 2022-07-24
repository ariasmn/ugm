package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/ariasmn/ugm/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

var supportedOS = map[string]bool{
	"linux": true,
	"freebsd": true,
	"openbsd": true,
	"netbsd": true,
}

func main() {
	if !supportedOS[runtime.GOOS] {
		fmt.Println("Current OS not supported. Refer to the documentation for more information.")
		os.Exit(0)
	}

	p := tea.NewProgram(tui.InitialModel(), tea.WithAltScreen())
	
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
