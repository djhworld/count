package main

import "testing"
import "reflect"

var EMPTY interface{} = new(interface{})

func TestShouldInitialiseWithEmptyMap(t *testing.T) {
	c := NewCounter()

	if len(c.uniqueItems) != 0 {
		t.Error("Expected empty map")
	}
}

func TestShouldCountUniqueItems(t *testing.T) {
	expected := make(map[string]int)
	expected["hello"] = 2
	expected["goodbye"] = 1
	expected["daniel"] = 3

	c := NewCounter()
	c.Add("daniel")
	c.Add("hello")
	c.Add("daniel")
	c.Add("goodbye")
	c.Add("hello")
	c.Add("daniel")

	if reflect.DeepEqual(c.uniqueItems, expected) != true {
		t.Error("Expected actual to be the same as expected, got", c.uniqueItems)
	}
}

func TestShouldRender(t *testing.T) {
	c := NewCounter()
	var expected map[string]interface{} = make(map[string]interface{})
	expected["1\tdaniel\n"] = EMPTY
	expected["2\ttest\n"] = EMPTY

	var actual map[string]interface{} = make(map[string]interface{})
	mockWriter := &MockWriter{
		mockWrite: func(p []byte) (n int, err error) {
			actual[string(p)] = EMPTY
			return 0, nil
		},
	}

	c.Add("daniel")
	c.Add("test")
	c.Add("test")

	c.Render(mockWriter)

	if reflect.DeepEqual(expected, actual) != true {
		t.Error("Expected actual to be the same as expected, expected", expected, "got", actual)
	}
}

func TestShouldRenderNothingOnEmptyMap(t *testing.T) {
	var actual string = ""
	mockWriter := &MockWriter{
		mockWrite: func(p []byte) (n int, err error) {
			actual = actual + string(p)
			return 0, nil
		},
	}

	c := NewCounter()
	expected := ""

	c.Render(mockWriter)

	if expected != actual {
		t.Error("Expected actual to be the same as expected, expected", expected, "got", actual)
	}
}

type MockWriter struct {
	mockWrite func(p []byte) (n int, err error)
}

func (w *MockWriter) Write(p []byte) (n int, err error) {
	if w.mockWrite != nil {
		return w.mockWrite(p)
	}
	return 0, nil
}
