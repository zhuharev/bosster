package main

import "log"

func debug(fmt string, args ...interface{}) {
	if !Debug {
		return
	}
	log.Printf(fmt, args...)
}
