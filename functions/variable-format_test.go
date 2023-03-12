package functions_test

import (
	"testing"

	"github.com/LintaoAmons/string-go/functions"
)

func TestToCamelCase(t *testing.T) {
	input := "param_aaa_bbb_id"
	expectedOutput := "paramAaaBbbId"
	output := functions.ToCamelCase(input)
	if output != expectedOutput {
		t.Errorf("toCamelCase(%q) = %q, expected %q", input, output, expectedOutput)
	}
}
