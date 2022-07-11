package sqrt

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMany(t *testing.T) {
	// Read test cases from "sqrt_case.csv" and check them
	dat, err := os.ReadFile("./sqrt_cases.csv")
    if err != nil {
		t.Fatalf("File Not Found:  %s", err)
		os.Exit(1)
	}

	r := csv.NewReader(strings.NewReader(string(dat)))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("BAD %s", err)
		}

		// Converting Input String to Float32
		input,err := strconv.ParseFloat(record[0], 32)
		require.NoError(t, err)

		// Testing Square Root
		output, err := Sqrt(input)
		require.NoError(t, err)

		// Converting Expected Output String to Float32
		x_output,err := strconv.ParseFloat(record[1], 32)
		require.NoError(t, err)

		// Comparing Results
		require.InDelta(t, x_output, output, 0.001)
	}
}
