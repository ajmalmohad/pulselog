package main

import (
	"pulselog/auth/config"
)

func main() {
	config.LoadEnvironmentVars()
	config.ConnectDatabase()
}