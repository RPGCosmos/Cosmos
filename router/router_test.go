package router

import (
	"testing"
	"reflect"
	"fmt"
)

func sayHi(inp Input) error {
	fmt.Println()

	return nil
}

type routes struct {
	routeName string
	caller    Action
	result    error
}

var tests = []routes{
	{"1", sayHi, nil},
	{"HI", sayHi, nil},
	{"WORLD", sayHi, nil},
	{"!!", sayHi, nil},
}

func TestRouter_Register(t *testing.T) {
	router := NewRouter()

	for _, pair := range tests {
		fmt.Println(pair)

		err := router.Register(pair.routeName, pair.caller)
		if err != nil {
			t.Error(err)
		}

		exists := false
		calls := false
		for name, caller := range router.routes {
			if name == pair.routeName {
				exists = true
			}
			if reflect.ValueOf(caller) == reflect.ValueOf(pair.caller) {
				calls = true
			}
		}

		if exists == false {
			t.Error("Route not found in routes.")
		}
		if calls == false {
			t.Error("Route function not registered.")
		}
	}
}

func TestRouter_Unregister(t *testing.T) {
	router := NewRouter()

	for _, pair := range tests {
		fmt.Println(pair)

		router.Register(pair.routeName, pair.caller)

		err := router.Unregister(pair.routeName)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestRouter_Handle(t *testing.T) {
	router := NewRouter()

	router.Register("test", sayHi)

	router.Handle("test", Input{args: []string{"hello"}})
}
