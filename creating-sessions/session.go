package main

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.New()
		c = &http.Cookie{Name: "session", Value: sID.String()}
	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	var u user
	if s, ok := dbSessions[c.Value]; ok {
		u = dbUsers[s.un]
	}
	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s := dbSessions[c.Value]
	_, ok := dbUsers[s.un]
	return ok
}

func cleanSessions() {
	fmt.Println("Before clean")
	showSessions()
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second*30) || v.un == "" {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("After clean")
	showSessions()
}

func showSessions() {
	fmt.Println("--------------------")
	for k, v := range dbSessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}
