package main

import (
	"fmt"
	"io"
)

type Counter struct {
	uniqueItems map[string]int
}

func NewCounter() *Counter {
	c := new(Counter)
	c.uniqueItems = make(map[string]int)
	return c
}

func (c *Counter) Add(item string) {
	if _, ok := c.uniqueItems[item]; !ok {
		c.uniqueItems[item] = 1
	} else {
		c.uniqueItems[item] += 1
	}
}

func (c *Counter) Render(writer io.Writer, uniqueOnly bool) {
	for k, v := range c.uniqueItems {
		if uniqueOnly {
			fmt.Fprintf(writer, "%s\n", k)
		} else {
			fmt.Fprintf(writer, "%d\t%s\n", v, k)
		}
	}
}
