package cmd

import (
	"crypto/sha1"
	"fmt"
	"math/big"

	"github.com/malukimuthusi/asymmetric/crypt"
	"github.com/spf13/cobra"
)

var e int64
var n int64
var d int64
var msg int64

var rootCmd = &cobra.Command{
	Use:   "asymmetric",
	Short: "Asymmetric encrypt and decrypt messages",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Keys are N = %d, D = %d, E = %d\n", n, d, e)
		fmt.Printf("Encrypting the message: %d\n", msg)

		c := crypt.Encrypt(*big.NewInt(msg), *big.NewInt(e), *big.NewInt(n))

		h := sha1.New()
		h.Write([]byte(c.String()))

		bs := h.Sum([]byte(c.String()))
		fmt.Printf("The hash of the Cypher = %x\n", bs)

		fmt.Printf("Generated the cypher: %d\n", c.Int64())

		fmt.Println("Decrypting....")

		t := crypt.Decrypt(c, *big.NewInt(d), *big.NewInt(n))

		fmt.Printf("Original message: %d\n", t.Int64())
	},
}

func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().Int64VarP(&msg, "msg", "m", 2, "message")
	rootCmd.PersistentFlags().Int64VarP(&e, "e", "e", 5, "value of e")
	rootCmd.PersistentFlags().Int64VarP(&n, "n", "n", 14, "value of n")

	rootCmd.PersistentFlags().Int64VarP(&d, "d", "d", 11, "value of d")

	rootCmd.AddCommand(GenerateCMD())

}
