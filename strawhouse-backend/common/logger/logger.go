package logger

import (
	"fmt"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"reflect"
	"strings"
)

type Logger struct {
}

func (r *Logger) LogEvent(event fxevent.Event) {
	eventElem := reflect.TypeOf(event).Elem()
	if eventElem.NumField() < 2 {
		if eventElem.Name() == "Stopping" {
			fmt.Println()
			return
		}

		fmt.Printf("%s\n", eventElem.Name())
		return
	}

	constructorField := eventElem.Field(0)
	constructor := reflect.ValueOf(event).Elem().FieldByIndex([]int{constructorField.Index[0]}).String()
	if strings.HasPrefix(constructor, "go.uber.org/") {
		return
	}
	constructor = strings.TrimPrefix(constructor, "github.com/strawstacks/strawhouse/strawhouse-backend/")

	if eventElem.NumField() > 3 {
		output := eventElem.Field(3)
		outputValue := reflect.ValueOf(event).Elem().FieldByIndex([]int{output.Index[0]})
		if outputValue.Kind() != reflect.Slice {
			fmt.Printf("%-18s %s\n", eventElem.Name(), constructor)
			return
		}
		outputSlice := outputValue.Slice(0, 1).Index(0).Interface()
		fmt.Printf("%-18s %-36s %s\n", eventElem.Name(), constructor, outputSlice)
	}
}

func Init() fx.Option {
	return fx.WithLogger(func() fxevent.Logger {
		return new(Logger)
	})
}
