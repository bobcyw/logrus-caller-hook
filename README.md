# logrus-caller-hook [![Build Status](https://travis-ci.org/bobcyw/logrus-logstash-hook.svg?branch=master)](https://travis-ci.org/bobcyw/logrus-logstash-hook)
solve logrus's caller problem

By default:
```go
package main

import (
	"github.com/bobcyw/logrus"
	"github.com/bobcyw/logrus-caller-hook"
)

func main() {
	log := logrus.New()
	log.Hooks.Add(logrus_caller_hook.New())
	log.Println("hi")
}
```
The output looks like this ```time="2017-05-17T06:41:10+08:00" level=info msg=hi caller="main.go:11"```


By custom:
```go
package main

import (
	"github.com/bobcyw/logrus"
	"github.com/bobcyw/logrus-caller-hook"
)

func main() {
	log := logrus.New()
	log.Hooks.Add(logrus_caller_hook.NewCustom("callee", true))
	log.Println("hi")
}
```
The output is ```time="2017-05-17T06:43:11+08:00" level=info msg=hi callee="/Users/caoyawen/xxx/src/github.com/bobcyw/logrus-caller-hook/examples/custom/main.go:11"```
