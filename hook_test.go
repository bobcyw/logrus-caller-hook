package logrus_caller_hook

import (
	"bytes"
	"fmt"
	"runtime"
	"testing"
	"time"

	"github.com/bobcyw/logrus"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	log := logrus.New()
	log.Level = logrus.DebugLevel
	log.Hooks.Add(New())
	buff := bytes.Buffer{}
	log.Out = &buff
	log.Println("check log")
	assert.Equal(t, fmt.Sprintf("time=\"%s\" level=info msg=\"check log\" caller=\"hook_test.go:20\" \n", time.Now().Format(time.RFC3339)), buff.String())
	buff.Reset()
	log.WithField("type", "check").Println("check field")
	assert.Equal(t, fmt.Sprintf("time=\"%s\" level=info msg=\"check field\" caller=\"hook_test.go:23\" type=check \n", time.Now().Format(time.RFC3339)), buff.String())
}

func TestNewCustom(t *testing.T) {
	log := logrus.New()
	log.Level = logrus.DebugLevel
	log.Hooks.Add(NewCustom("callee", true))
	buff := bytes.Buffer{}
	log.Out = &buff
	log.Println("check log")
	assert.Equal(t, fmt.Sprintf("time=\"%s\" level=info msg=\"check log\" callee=\"%s:33\" \n", time.Now().Format(time.RFC3339), currentFile()), buff.String())
	buff.Reset()
	log.WithField("type", "check").Println("check field")
	assert.Equal(t, fmt.Sprintf("time=\"%s\" level=info msg=\"check field\" callee=\"%s:36\" type=check \n", time.Now().Format(time.RFC3339), currentFile()), buff.String())
}

func currentFile() string {
	_, filename, _, _ := runtime.Caller(1)
	return filename
}

