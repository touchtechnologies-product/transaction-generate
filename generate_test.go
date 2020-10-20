package transaction_generate

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateDigit(t *testing.T) {
	expect := "AS001-0001"
	idGenerate := Generate("AS", "1", "1")
	require.Equal(t, expect, idGenerate)
}

func TestGenerateTenDigit(t *testing.T) {
	expect := "AS001-0010"
	idGenerate := Generate("AS", "1", "10")
	require.Equal(t, expect, idGenerate)
}

func TestGenerateThousandsDigit(t *testing.T) {
	expect := "AS001-1000"
	idGenerate := Generate("AS", "1", "1000")
	require.Equal(t, expect, idGenerate)
}

func TestGenerateTenThousandsDigit(t *testing.T) {
	expect := "AS001-10000"
	idGenerate := Generate("AS", "1", "10000")
	require.Equal(t, expect, idGenerate)
}

func TestGenerateMillionsDigit(t *testing.T) {
	expect := "AS001-99999999"
	idGenerate := Generate("AS", "1", "99999999")
	require.Equal(t, expect, idGenerate)
}
