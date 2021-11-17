package cmd

import (
	"fmt"
	"log"

	"github.com/malukimuthusi/asymmetric/generatekeys"
	"github.com/spf13/cobra"
)

// command for generating public and private keys for alice
func GenerateCMD() *cobra.Command {
	generate := &cobra.Command{
		Use:   "generate",
		Short: "Generate Public and private Keys",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Generate Private and Public Keys for alice")

			p, q := generatekeys.GeneratePrimeNumbers()

			fmt.Printf("Values: P=%d, Q=%d\n", p, q)

			// generate Alice's public key
			n := generatekeys.Product(p, q)

			fmt.Printf("Value of N is %d\n", n)

			// generate Alice Private Key

			e, err := generatekeys.E(n, p, q)
			if err != nil {
				log.Fatalf("program failed, err: %v", err)
			}

			fmt.Printf("Alice's Private Key is %d\n", e)

			phi := generatekeys.Phi(p, q)
			d, err := generatekeys.D(n, e, phi)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Alice's Public Key is %d\n", d)

		},
	}

	return generate
}
