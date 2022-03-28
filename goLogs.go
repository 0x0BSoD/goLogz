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

type LogLevel int

const (
	Trace LogLevel = iota
	Info
	Warning
	Error
)

func (l LogLevel) String() string {
	_strings := []string{
		"Trace",
		"Info",
		"Warning",
		"Error",
	}
	return _strings[l]
}

type GoLogz struct {
	loggers map[string]*log.Logger
	Colors  bool
	Level   LogLevel
	colors  map[string]string
}

type ParameterItem struct {
	Level     string
	OutHandle string
	LineNum   bool
}

func Init(params []ParameterItem) (GoLogz, error) {

	items := make(map[string]*log.Logger)
	colors := make(map[string]string, 4)

	//default config
	items[Trace.String()] = log.New(os.Stdout, "[TRACE] ", log.Ldate|log.Ltime|log.LUTC)
	items[Info.String()] = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.LUTC)
	items[Warning.String()] = log.New(os.Stdout, "[WARNING] ", log.Ldate|log.Ltime|log.LUTC)
	items[Error.String()] = log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime|log.LUTC)

	//default colors
	colors[Trace.String()] = colorWhite
	colors[Info.String()] = colorBlue
	colors[Warning.String()] = colorYellow
	colors[Error.String()] = colorRed

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

	return GoLogz{loggers: items, colors: colors}, nil
}

func (g *GoLogz) log(l LogLevel, prettyJson bool, m ...interface{}) {
	jump := false
	if prettyJson {
		b, err := json.MarshalIndent(m[0], "", "  ")
		if err == nil {
			if g.Colors {
				g.loggers[l.String()].Println(g.colors[l.String()], "🡦\n"+string(b), colorReset)
			} else {
				g.loggers[l.String()].Println("🡦\n" + string(b))
			}
		} else {
			jump = true
		}
	} else {
		jump = true
	}

	if jump {
		if g.Colors {
			g.loggers[l.String()].Println(g.colors[l.String()], m, colorReset)
		} else {
			g.loggers[l.String()].Println(m...)
		}
	}
}

func (g *GoLogz) Trace(msg ...interface{}) {
	if g.Level == Trace {
		g.log(Trace, false, msg...)
	}
}
func (g *GoLogz) TraceJson(msg ...interface{}) {
	if g.Level == Trace {
		g.log(Trace, true, msg...)
	}
}

func (g *GoLogz) Info(msg ...interface{}) {
	if (g.Level == Trace) || (g.Level == Info) {
		g.log(Info, false, msg...)
	}
}
func (g *GoLogz) InfoJson(msg ...interface{}) {
	if (g.Level == Trace) || (g.Level == Info) {
		g.log(Info, true, msg...)
	}
}

func (g *GoLogz) Warning(msg ...interface{}) {
	if (g.Level == Trace) || (g.Level == Info) || (g.Level == Warning) {
		g.log(Warning, false, msg...)
	}
}
func (g *GoLogz) WarningJson(msg ...interface{}) {
	if (g.Level == Trace) || (g.Level == Info) || (g.Level == Warning) {
		g.log(Warning, true, msg...)
	}
}

func (g *GoLogz) Error(msg ...interface{}) {
	if (g.Level == Trace) || (g.Level == Info) || (g.Level == Warning) || (g.Level == Error) {
		g.log(Error, false, msg...)
	}
}
func (g *GoLogz) ErrorJson(msg ...interface{}) {
	if (g.Level == Trace) || (g.Level == Info) || (g.Level == Warning) || (g.Level == Error) {
		g.log(Error, true, msg...)
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
