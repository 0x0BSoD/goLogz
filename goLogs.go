package goLogz

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorWhite  = "\033[37m"
)

type GoLogz struct {
	loggers map[string]*log.Logger
	Colors  bool
}

type ParameterItem struct {
	Level     string
	OutHandle string
	LineNum   bool
}

func Init(params []ParameterItem) (GoLogz, error) {

	items := make(map[string]*log.Logger)

	//default config
	items["Trace"] = log.New(os.Stdout, "[TRACE] ", log.Ldate|log.Ltime|log.LUTC)
	items["Info"] = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.LUTC)
	items["Warning"] = log.New(os.Stdout, "[WARNING] ", log.Ldate|log.Ltime|log.LUTC)
	items["Error"] = log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime|log.LUTC)

	defFlags := log.Ldate | log.Ltime | log.LUTC | log.Lshortfile
	noLineNumFlags := log.Ldate | log.Ltime | log.LUTC

	for _, p := range params {
		flags := defFlags
		if !p.LineNum {
			flags = noLineNumFlags
		}

		switch p.Level {
		case "Trace":
			handle, err := whatIoChecker(p.OutHandle)
			if err != nil {
				return GoLogz{}, err
			}
			items["Trace"] = log.New(handle, "[TRACE] ", flags)
		case "Info":
			handle, err := whatIoChecker(p.OutHandle)
			if err != nil {
				return GoLogz{}, err
			}
			items["Info"] = log.New(handle, "[INFO] ", flags)
		case "Warning":
			handle, err := whatIoChecker(p.OutHandle)
			if err != nil {
				return GoLogz{}, err
			}
			items["Warning"] = log.New(handle, "[WARNING] ", flags)
		case "Error":
			handle, err := whatIoChecker(p.OutHandle)
			if err != nil {
				return GoLogz{}, err
			}
			items["Error"] = log.New(handle, "[ERROR] ", flags)
		default:
			handle, err := whatIoChecker(p.OutHandle)
			if err != nil {
				return GoLogz{}, err
			}
			items[p.Level] = log.New(handle, fmt.Sprintf("[%s] ", p.Level), flags)
		}
	}

	return GoLogz{loggers: items}, nil
}

func (g *GoLogz) Trace(msg ...interface{}) {
	if g.Colors {
		g.loggers["Trace"].Println(colorWhite, msg, colorReset)
	} else {
		g.loggers["Trace"].Println(msg...)
	}
}
func (g *GoLogz) TraceJson(msg interface{}) {
	b, err := json.MarshalIndent(msg, "", "  ")
	if err == nil {
		msg = string(b)
	}
	if g.Colors {
		g.loggers["Trace"].Println(colorWhite, "🡦\n", msg, colorReset)
	} else {
		g.loggers["Trace"].Println(msg, "🡦\n")
	}
}

func (g *GoLogz) Info(msg ...interface{}) {
	if g.Colors {
		g.loggers["Info"].Println(colorBlue, msg, colorReset)
	} else {
		g.loggers["Info"].Println(msg...)
	}
}
func (g *GoLogz) InfoJson(msg interface{}) {
	b, err := json.MarshalIndent(msg, "", "  ")
	if err == nil {
		msg = string(b)
	}
	if g.Colors {
		g.loggers["Info"].Println(colorBlue, "🡦\n", msg, colorReset)
	} else {
		g.loggers["Info"].Println("🡦\n", msg)
	}
}

func (g *GoLogz) Warning(msg ...interface{}) {
	if g.Colors {
		g.loggers["Warning"].Println(colorYellow, msg, colorReset)
	} else {
		g.loggers["Warning"].Println(msg...)
	}
}

func (g *GoLogz) Error(msg ...interface{}) {
	if g.Colors {
		g.loggers["Error"].Println(colorRed, msg, colorReset)
	} else {
		g.loggers["Error"].Println(msg...)
	}
}

func (g *GoLogz) Custom(name string, msg ...interface{}) {
	if logger, ok := g.loggers[name]; ok {
		if g.Colors {
			logger.Println(colorPurple, msg, colorReset)
		} else {
			logger.Println(msg...)
		}

	}
}
