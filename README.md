# phuslog_logger

Wrapper Phuslog - Fastest structured logging.

[![Go Report Card](https://goreportcard.com/badge/github.com/prongbang/phuslog_logger)](https://goreportcard.com/report/github.com/prongbang/phuslog_logger)

### Install

```
go get github.com/prongbang/phuslog_logger
```

### How to use

```golang
logs := phuslog_logger.NewPhuslogLogger()
```

- Data

```golang
d := &data{Key: "Key", Value: "Value"}
```

- Info

```golang
logs.Info("test", "data", d) 
```

```golang
2024-06-22T20:26:02.298+07:00 INF 35 phuslog_logger/phuslog_logger_test.go:20 > test data={"key":"Key","value":"Value"}
```

- Warn

```golang
logs.Warn("test", "data", d) 
```

```golang
2024-06-22T20:26:02.298+07:00 WRN 35 phuslog_logger/phuslog_logger_test.go:21 > test data={"key":"Key","value":"Value"}
```

- Error

```golang
logs.Error("test", "data", d) 
```

```golang
2024-06-22T20:26:02.298+07:00 ERR 35 phuslog_logger/phuslog_logger_test.go:23 > test data={"key":"Key","value":"Value"}
```

- Fatal

```golang
logs.Fatal("test", "data", d) 
```

```golang
2024-06-22T20:29:37.596+07:00 FTL 35 phuslog_logger/phuslog_logger_test.go:24 > test data={"key":"Key","value":"Value"}
```
