package utils

import (
	"bytes"
	"errors"
	log "github.com/sirupsen/logrus"
	"net"
	"strings"
	"sync"
)

func Fatal(any ...interface{}) {
	log.Fatal(any)
}

func Error(any ...interface{}) error {
	log.Error("******* ", any)
	return errors.New("*****")
}

func Warn(any ...interface{}) {
	log.Warn("***** ", any)
}

func Info(any ...interface{}) {
	log.Info("*** ", any)
}

func Debug(any ...interface{}) {
	log.Debug(any)
}

func Trace(any ...interface{}) {
	log.Trace(any)
}

var conn net.Conn
var mtx = &sync.Mutex{}
var lineFeed = []byte("\n")

func SetConsoleConnection(connection net.Conn) {
	mtx.Lock()
	defer mtx.Unlock()
	conn = connection
}

func Print(msg string) {
	mtx.Lock()
	defer mtx.Unlock()
	if conn != nil {
		conn.Write([]byte(msg))
	}
}

func Println(msg string) {
	mtx.Lock()
	defer mtx.Unlock()
	if conn != nil {
		conn.Write([]byte(msg))
		conn.Write(lineFeed)
	}
}

func Read() (string, error) {
	if conn == nil {
		return "", nil
	}
	line := make([]byte, 4096)
	n, e := conn.Read(line)
	if e != nil {
		e = Error("Failed to read line:", e)
		return "", e
	}
	inputLine := strings.TrimSpace(string(line[0:n]))
	return inputLine, nil
}

func NewError(msg ...string) error {
	if msg == nil || len(msg) == 0 {
		return nil
	}
	buff := bytes.Buffer{}
	for _, str := range msg {
		buff.WriteString(str)
		buff.WriteString(" ")
	}
	return errors.New(buff.String())
}
