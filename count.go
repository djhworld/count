package main

import (
	"bufio"
	"os"

	"github.com/urfave/cli"
)

const UNIQUE_ONLY = "unique"
const ERR = 255

func openFile(filename string) (*os.File, error) {
	if filename == "" || filename == "-" {
		return os.Stdin, nil
	}

	return os.Open(filename)
}

func countInput(c *cli.Context) error {
	file, err := openFile(c.Args().Get(0))
	if err != nil {
		return cli.NewExitError(err, ERR)
	}
	defer file.Close()

	counter := NewCounter(RenderOptions{uniqueOnly: c.Bool(UNIQUE_ONLY)})
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		counter.Add(scanner.Text())
	}

	counter.Render(os.Stdout)
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "count"
	app.Usage = "Count unique entries in line delimited input from stdin"
	app.Usage = "Count unique lines in line delimited input"
	app.UsageText = "count <input_file>\n\n\t if input_file is a single dash (`-`) or absent, the standard input is read."
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Daniel Harper",
			Email: "@djhworld",
		},
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "unique, u",
			Usage: "only print unique lines",
		},
	}
	app.Action = countInput
	app.Run(os.Args)
}
