package logrus_caller_hook

import (
	"bytes"
	"testing"

	"github.com/bobcyw/logrus"
)

func BenchmarkNew(b *testing.B) {
	log := logrus.New()
	log.Level = logrus.DebugLevel
	log.Hooks.Add(New())
	buff := bytes.Buffer{}
	buff.Grow(172000000)
	log.Out = &buff
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Println("hello")
		}
	})
}

func BenchmarkNewCustom(b *testing.B) {
	log := logrus.New()
	log.Level = logrus.DebugLevel
	log.Hooks.Add(NewCustom("callee", true))
	buff := bytes.Buffer{}
	buff.Grow(172000000)
	log.Out = &buff
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Println("hello")
		}
	})
}
