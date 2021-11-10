package cmd

import (
	"fmt"

	"github.com/malukimuthusi/asymmetric-algorithm/asymmetric"
	"github.com/spf13/cobra"
)

var alicePrivateKey int
var alicePublicKey int
var bobPrivateKey int
var bobPublicKey int
var msg int

var rootCmd = &cobra.Command{
	Use:   "asymmetric-algorithm",
	Short: "Asymmetric algorithm, encrypt and decrypt messages",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Alice Encrypting the message: %d\n", msg)

		cypher := asymmetric.Encrypt(msg, uint64(alicePrivateKey), uint64(bobPublicKey))

		fmt.Printf("Generated the cypher: %d\n", cypher)

		fmt.Println("Bob will decrypt the cypher")

		m := asymmetric.Decrypt(cypher, uint64(bobPrivateKey), uint64(alicePublicKey))

		fmt.Printf("Bob decrypted and obtained the message: %d", m)
	},
}

func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&msg, "msg", "m", 2, "message")
	rootCmd.PersistentFlags().IntVarP(&alicePrivateKey, "ap", "a", 5, "Alice private key")
	rootCmd.PersistentFlags().IntVarP(&alicePublicKey, "ab", "b", 14, "Alice Public key")

	rootCmd.PersistentFlags().IntVarP(&bobPrivateKey, "bp", "c", 11, "Bob Private key")
	rootCmd.PersistentFlags().IntVarP(&bobPublicKey, "bb", "d", 14, "bob Public key")

}
