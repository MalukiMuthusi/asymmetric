package main

import (
	"github.com/malukimuthusi/asymmetric-algorithm/cmd"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	// generate two prime numbers

	cmd.Execute()
}
