package generatekeys_test

import (
	"testing"

	"github.com/malukimuthusi/asymmetric/generatekeys"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePrimeNumbers(t *testing.T) {
	tests := []struct {
		name string
		n    uint64
	}{
		{
			name: "happy case",
			n:    100,
		},
		{
			name: "generate large prime numbers",
			n:    1000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			n1, n2 := generatekeys.GeneratePrimeNumbers(tt.n)
			assert.NotEqualValues(t, n1, n2)
			assert.True(t, n1 > 1)
			assert.True(t, n2 > 1)
		})
	}

}
