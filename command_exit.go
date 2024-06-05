package main

import "os"

func commandExit(cnf *config, args ...string) error {
	os.Exit(1)
	return nil
}
