package main

import (
	"bufio"
	"fmt"
	"gopkg.in/urfave/cli"
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
	results := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var item string = scanner.Text()
		if _, ok := results[item]; !ok {
			results[item] = 1
		} else {
			results[item] += 1
		}
	}

	for k, v := range results {
		fmt.Printf("%s\t%d\n", k, v)
	}

	return nil
}
