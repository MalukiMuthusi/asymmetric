package cmd

import (
	"fmt"

	"github.com/malukimuthusi/asymmetric-algorithm/generatekeys"
	"github.com/spf13/cobra"
)

func GenerateCMD() *cobra.Command {
	generate := &cobra.Command{
		Use:   "generate",
		Short: "Generate Public and private Keys",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Generate Private Keys for alice")

			p, q := generatekeys.GeneratePrimeNumbers()

			// generate Alice's public key
			pb := generatekeys.Product(p, q)

			fmt.Printf("Alice's Public Key is %d\n", pb)

			// generate Alice Private Key
			
		},
	}

	return generate
}
