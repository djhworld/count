package main

import (
	"fmt"
	"io"
)

type RenderOptions struct {
	uniqueOnly bool
}

type Counter struct {
	renderOptions RenderOptions
	uniqueItems   map[string]int
}

func NewCounter(renderOptions RenderOptions) *Counter {
	c := new(Counter)
	c.uniqueItems = make(map[string]int)
	c.renderOptions = renderOptions
	return c
}

func (c *Counter) Add(item string) {
	if _, ok := c.uniqueItems[item]; !ok {
		c.uniqueItems[item] = 1
	} else {
		c.uniqueItems[item] += 1
	}
}

func (c *Counter) Render(writer io.Writer) {
	for k, v := range c.uniqueItems {
		if c.renderOptions.uniqueOnly {
			fmt.Fprintf(writer, "%s\n", k)
		} else {
			fmt.Fprintf(writer, "%d\t%s\n", v, k)
		}
	}
}
