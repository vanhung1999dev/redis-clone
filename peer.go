package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"github.com/tidwall/resp"
)

type Peer struct {
	conn  net.Conn
	msgCh chan Message
	delCh chan *Peer
}

func (p *Peer) Send(msg []byte) (int, error) {
	return p.conn.Write(msg)
}

func NewPeer(conn net.Conn, msgCh chan Message, delCh chan *Peer) *Peer {
	return &Peer{
		conn:  conn,
		msgCh: msgCh,
		delCh: delCh,
	}
}

func (p *Peer) readLoop() error {
	rd := resp.NewReader(p.conn)
	for {
		v, _, err := rd.ReadValue()
		if err == io.EOF {
			p.delCh <- p
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if v.Type() != resp.Array {
			continue
		}

		arr := v.Array()
		if len(arr) == 0 {
			continue
		}

		rawCMD := arr[0]
		cmdStr := strings.ToLower(rawCMD.String())

		// helper to send error reply to client
		sendErr := func(msg string) {
			errResp := fmt.Sprintf("-ERR %s\r\n", msg)
			_, _ = p.conn.Write([]byte(errResp))
		}

		var cmd Command

		switch cmdStr {
		case CommandClient:
			if len(arr) < 2 {
				sendErr("wrong number of arguments for 'client' command")
				continue
			}
			cmd = ClientCommand{
				value: arr[1].String(),
			}

		case CommandGET:
			if len(arr) < 2 {
				sendErr("wrong number of arguments for 'get' command")
				continue
			}
			cmd = GetCommand{
				key: arr[1].Bytes(),
			}

		case CommandSET:
			if len(arr) < 3 {
				sendErr("wrong number of arguments for 'set' command")
				continue
			}
			cmd = SetCommand{
				key: arr[1].Bytes(),
				val: arr[2].Bytes(),
			}

		case CommandHELLO:
			if len(arr) < 2 {
				sendErr("wrong number of arguments for 'hello' command")
				continue
			}
			cmd = HelloCommand{
				value: arr[1].String(),
			}

		case CommandEXISTS:
			if len(arr) < 2 {
				sendErr("wrong number of arguments for 'exists' command")
				continue
			}
			cmd = ExistsCommand{
				key: arr[1].Bytes(),
			}

		case CommandDEL:
			if len(arr) < 2 {
				sendErr("wrong number of arguments for 'del' command")
				continue
			}
			cmd = DelCommand{
				key: arr[1].Bytes(),
			}

		default:
			sendErr(fmt.Sprintf("unknown command '%s'", cmdStr))
			continue
		}

		p.msgCh <- Message{
			cmd:  cmd,
			peer: p,
		}
	}
	return nil
}
