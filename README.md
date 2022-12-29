# Intro [![GoDoc](https://godoc.org/github.com/omerye/gojdb?status.svg)](https://godoc.org/github.com/omerye/gojdb/jdwp)

gojdb is a debugger for Java implemented in Go.

# TODO

* [ ] Add a reference to the original project (inside LIECENCE) (plus the commit on copy)
* [ ] Another abstraction layer for remote debugging.
* [ ] Make this layer callable by an RPC method.
* [ ] Wrap `adb` functionallity (using package `github.com/zach-klippenstein/goadb`)
* [ ] Interactive command line interface layer.
* [ ] Try extending `jdwp.Value` to make it easier to cast.
* [ ] Handle single step.
* [ ] Backtrace.
* [ ] Link to source path
* [ ] Break on app startup
* [ ] Change methods to return maps instead of slice for better efficiency (Example - conn.GetClasses should return `map[ClassID]ClassInfo` instead of `[]ClassInfo`)
* [ ] Trace
* [ ] Operations on demand (example - trace/untrace on location etc..)
* Get some ideas from [Delve](https://github.com/go-delve/delve)
* Get some ideas from JDB
* Probably shouldn't use any of the available prompt library as each one of them looks trash. Just `readlines` looks fine.
