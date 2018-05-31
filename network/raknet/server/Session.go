package server

import "time"

type Session struct {
	LastUpdate time.Time
}

func (session *Session) Tick(t int64) {

}

func (session *Session) Close() {

}
