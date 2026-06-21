package main

import (
	"flag"
	"log"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/VincentZyuApps/htop-go/internal/dashboard"
	"github.com/VincentZyuApps/htop-go/internal/ui"
	appversion "github.com/VincentZyuApps/htop-go/internal/version"
)

func main() {
	styleFlag := flag.String("style", string(dashboard.StyleClassic), "UI style: classic or overview")
	tableFlag := flag.Bool("table", false, "enable lower process table and full process sampling")
	taskFlag := flag.Bool("task", false, "show task metrics")
	loadFlag := flag.Bool("load", false, "show load metrics")
	interval := 2 * time.Second
	flag.DurationVar(&interval, "t", interval, "refresh interval, e.g. 500ms, 2s, 5s")
	flag.DurationVar(&interval, "interval", interval, "refresh interval, e.g. 500ms, 2s, 5s")
	flag.Parse()

	style := dashboard.Style(*styleFlag)
	switch style {
	case dashboard.StyleClassic, dashboard.StyleOverview:
	default:
		log.New(os.Stderr, "", 0).Fatalf("invalid --style %q; valid values: classic, overview", *styleFlag)
	}

	p := tea.NewProgram(
		ui.NewModel(appversion.Number, style, *tableFlag, *taskFlag, *loadFlag, interval),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		log.New(os.Stderr, "", 0).Fatalf("htop-go: %v", err)
	}
}
