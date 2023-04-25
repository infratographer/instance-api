//go:generate sqlboiler crdb --add-soft-deletes

// package main is the entry point
package main

import "go.infratographer.com/instance-api/cmd"

func main() {
	cmd.Execute()
}
