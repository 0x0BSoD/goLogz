# goLogz

Default levels:
  * TRACE
  * INFO
  * WARNING
  * ERROR
  
 Adding new level:
 ```go
logs, err := Init([]ParameterItem{
		{
			Level:     "ACTION",
			OutHandle: "STDOUT",
			LineNum: false,
		},
	})
```

**OutHandle** - can be a path to file or one of `STDOUT | STDERR | DISCARD`

Example:
```go
logs, err := Init([]ParameterItem{
    {
        Level:     "ACTION",
        OutHandle: "STDOUT",
        LineNum: false,
    },
})
if err != nil {
    panic(err)
}

logs.colors = true

logs.Trace("this is a trace")
logs.Info("this is a info")
logs.Warning("this is a warning")
logs.Error("this is a error")
logs.Custom("ACTION", "this is a custom")
```
Output:
```
[TRACE] 2020/08/03 17:48:06 goLogs.go:89: this is a trace  
[INFO] 2020/08/03 17:48:06 goLogs.go:97:  this is a info 
[WARNING] 2020/08/03 17:48:06 goLogs.go:105:  this is a warning 
[ERROR] 2020/08/03 17:48:06 goLogs.go:113:  this is a error
[ACTION] 2020/08/03 17:48:06  this is a custom 
```