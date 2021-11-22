package cmd

import (
	"fmt"
	"log"
	"math/big"

	"github.com/malukimuthusi/asymmetric/generatekeys"
	"github.com/spf13/cobra"
)

// command for generating public and private keys for alice
func GenerateCMD() *cobra.Command {
	generate := &cobra.Command{
		Use:   "generate",
		Short: "Generate Public and private Keys",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Generate n , e, d keys")

			p, q := generatekeys.GeneratePrimeNumbers(*big.NewInt(100))

			fmt.Printf("Values: P=%d, Q=%d\n", p, q)

			// generate N
			n := generatekeys.Product(p, q)

			fmt.Printf("Value of N is %d\n", n)

			// generate E

			e, err := generatekeys.E(n, p, q)
			if err != nil {
				log.Fatalf("program failed, err: %v", err)
			}

			fmt.Printf("Value of E Key is %d\n", e)

			phi := generatekeys.Phi(p, q)
			d, err := generatekeys.D(e, phi)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Value of D is %d\n", d)

		},
	}

	return generate
}
