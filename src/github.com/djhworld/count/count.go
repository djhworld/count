package main

import (
	"bufio"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "count"
	app.Usage = "Count unique entries in line delimited input from stdin"
	app.Action = countInput
	app.Run(os.Args)
}

func countInput(c *cli.Context) error {
	counter := NewCounter()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		counter.Add(scanner.Text())
	}

	counter.Render(os.Stdout)
	return nil
}
