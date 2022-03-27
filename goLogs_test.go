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

	logs.Colors = true

	logs.Trace("this is a trace")
	logs.Info("this is a info")
	logs.InfoJson(map[string]int{"foo": 1, "bar": 2, "baz": 3})
	logs.Warning("this is a warning")
	logs.Error("this is a error")
	logs.Custom("ACTION", "this is a custom")
}
