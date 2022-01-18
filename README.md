# mini-redash-go

## Usage

```
$ go run main.go connect --file sample/sample.sql
{"level":"info","ts":1642549944.589237,"caller":"cmd/connect.go:76","msg":"database connected"}
+----+-------+
| ID | NAME  |
+----+-------+
|  1 | test2 |
|  2 | test4 |
|  3 | test1 |
|  4 | test3 |
+----+-------+
```
