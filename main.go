package main

import (
	"github.com/Miciah/clam-scanner/cmd"
	"github.com/golang/glog"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		glog.Exitln(err)
	}
}
