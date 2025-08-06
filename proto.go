package main

import (
	"bytes"
	"fmt"

	"github.com/tidwall/resp"
)

const (
	CommandSET    = "set"
	CommandGET    = "get"
	CommandHELLO  = "hello"
	CommandClient = "client"
	CommandEXISTS = "exists"
	CommandDEL    = "del"
)

type Command interface {
}

type SetCommand struct {
	key, val []byte
	expire   int64
	nx       bool
	xx       bool
}

type ClientCommand struct {
	value string
}

type HelloCommand struct {
	value string
}

type GetCommand struct {
	key []byte
}

type ExistsCommand struct {
	key []byte
}

type DelCommand struct {
	key []byte
}

func respWriteMap(m map[string]string) []byte {
	buf := &bytes.Buffer{}
	buf.WriteString("%" + fmt.Sprintf("%d\r\n", len(m)))
	rw := resp.NewWriter(buf)
	for k, v := range m {
		rw.WriteString(k)
		rw.WriteString(":" + v)
	}

	return buf.Bytes()
}
