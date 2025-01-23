package main

import (
	"fmt"

	"gitlab.com/scabala/zelligo"
)

type Plugin struct {
	tabs     map[string]bool
	testRuns uint32
	config   map[string]string
	modeLog  map[string]int
}

// Pipe implements zelligo.ZellijPlugin.
func (p *Plugin) Pipe(message zelligo.PipeMessage) (bool, error) {
	fmt.Println("unimplemented Pipe")
	return false, nil
}

// Render implements zelligo.ZellijPlugin.
func (p *Plugin) Render(rows uint32, cols uint32) error {
	fmt.Println("unimplemented Render")
	return nil
}

// Update implements zelligo.ZellijPlugin.
func (p *Plugin) Update(event zelligo.Event) (bool, error) {
	fmt.Println("unimplemented Update")
	return false, nil
}

func (p *Plugin) Load(configuration map[string]string) error {
	fmt.Printf("unimplemented Load %+v\n", configuration)
	return nil
}

func init() {
	zelligo.RegisterPlugin(&Plugin{})
}

func main() {}
