package cmd

import (
	"fmt"
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

			p, q := generatekeys.GeneratePrimeNumbers(*big.NewInt(1000))

			fmt.Printf("Values: P=%d, Q=%d\n", p.Int64(), q.Int64())

			// generate N
			n := generatekeys.Product(p, q)

			fmt.Printf("Value of N = %d\n", n.Int64())

			// generate E

			e := generatekeys.E(n, p, q)

			fmt.Printf("Value of E Key = %d\n", e.Int64())

			phi := generatekeys.Phi(p, q)
			fmt.Printf("Value of PHI = %d\n", phi.Int64())

			d := generatekeys.D(e, phi)

			fmt.Printf("Value of D = %d\n", d.Int64())

		},
	}

	return generate
}
