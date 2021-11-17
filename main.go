package main

// MALUKI MUTHUSI
// P15/81741/2017

import (
	"github.com/malukimuthusi/asymmetric-algorithm/cmd"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()

	cmd.Execute()
}
