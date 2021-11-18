package cmd

import (
	"fmt"

	"github.com/malukimuthusi/asymmetric/crypt"
	"github.com/spf13/cobra"
)

var e int
var n int
var d int
var msg int

var rootCmd = &cobra.Command{
	Use:   "asymmetric",
	Short: "Asymmetric encrypt and decrypt messages",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Encrypting the message: %d\n", msg)

		cypher := crypt.Encrypt(msg, uint64(e), uint64(n))

		fmt.Printf("Generated the cypher: %d\n", cypher)

		fmt.Println("Decrypting")

		m := crypt.Decrypt(cypher, uint64(d), uint64(n))

		fmt.Printf("Decrypted message: %d\n", m)
	},
}

func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&msg, "msg", "m", 2, "message")
	rootCmd.PersistentFlags().IntVarP(&e, "e", "e", 5, "value of e")
	rootCmd.PersistentFlags().IntVarP(&n, "n", "n", 14, "value of n")

	rootCmd.PersistentFlags().IntVarP(&d, "d", "d", 11, "value of d")

	rootCmd.AddCommand(GenerateCMD())

}
