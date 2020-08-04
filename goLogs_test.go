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
	logs.Warning("this is a warning")
	logs.Error("this is a error")
	logs.Custom("ACTION", "this is a custom")
}
