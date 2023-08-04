package main

import (
	"github.com/goraft/raft"
	"net/http"
)

func main() {
	s, _ := raft.NewServer("", "", nil, nil, nil, "")
	command := &raft.DefaultJoinCommand{}
	if _, err := s.Do(command); err != nil {
		http.Error(nil, err.Error(), http.StatusInternalServerError)
		return
	}
}
