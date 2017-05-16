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
