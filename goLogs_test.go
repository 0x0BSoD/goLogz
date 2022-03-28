package goLogz

import "testing"

func TestInit(t *testing.T) {
	logs, err := Init([]ParameterItem{
		{
			Level:     "ACTION",
			OutHandle: "STDOUT",
			LineNum:   false,
		},
	})
	if err != nil {
		t.Failed()
		return
	}

	logs.Colors = false
	logs.Level = Error

	logs.Trace("this is a trace")
	logs.TraceJson(map[string]int{"foo": 1, "bar": 2, "baz": 3})
	logs.Info("this is a info")
	logs.InfoJson(map[string]int{"foo": 1, "bar": 2, "baz": 3})
	logs.Warning("this is a warning")
	logs.WarningJson(map[string]int{"foo": 1, "bar": 2, "baz": 3})
	logs.Error("this is a error")
	logs.ErrorJson(map[string]int{"foo": 1, "bar": 2, "baz": 3})
	logs.Custom("ACTION", "this is a custom")
}
