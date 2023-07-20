package test

import (
	"testing"

	"github.com/hashicorp/aws-sdk-go-base/v2/diag"
)

type DiagsValidator func(*testing.T, diag.Diagnostics)

type ErrValidator func(error) bool

type DiagValidator func(diag.Diagnostic) bool

func ExpectNoDiags(t *testing.T, diags diag.Diagnostics) {
	expectDiagsCount(t, diags, 0)
}

func ExpectErrDiagValidator(msg string, ev ErrValidator) DiagsValidator {
	return func(t *testing.T, diags diag.Diagnostics) {
		// Check for the correct type of error before checking for single diagnostic
		if !expectDiagsContainsErr(diags, ev) {
			t.Fatalf("expected %s, got %#v", msg, diags)
		}

		expectDiagsCount(t, diags, 1)
	}
}

func ExpectDiagValidator(msg string, dv DiagValidator) DiagsValidator {
	return func(t *testing.T, diags diag.Diagnostics) {
		// Check for the correct type of error before checking for single diagnostic
		if !expectDiagsContainsDiagFunc(diags, dv) {
			t.Fatalf("expected %s, got %#v", msg, diags)
		}

		expectDiagsCount(t, diags, 1)
	}

}

func ExpectWarningDiagValidator(expected diag.Diagnostic) DiagsValidator {
	return func(t *testing.T, diags diag.Diagnostics) {
		// Check for the correct type of error before checking for single diagnostic
		if !expectDiagsContainsDiag(diags, expected) {
			t.Fatalf("expected Diagnostic matching %#v, got %#v", expected, diags)
		}

		expectDiagsCount(t, diags, 1)
	}
}

func expectDiagsCount(t *testing.T, diags diag.Diagnostics, c int) {
	if l := diags.Count(); l != c {
		t.Fatalf("Diagnostics: expected %d element, got %d\n%#v", c, l, diags)
	}
}

func expectDiagsContainsErr(diags diag.Diagnostics, ev ErrValidator) bool {
	for _, d := range diags.Errors() {
		if e, ok := d.(diag.DiagnosticWithErr); ok {
			if ev(e.Err()) {
				return true
			}
		}
	}
	return false
}

func expectDiagsContainsDiag(diags diag.Diagnostics, expected diag.Diagnostic) bool {
	for _, d := range diags {
		if d.Equal(expected) {
			return true
		}
	}
	return false
}

func expectDiagsContainsDiagFunc(diags diag.Diagnostics, dv DiagValidator) bool {
	for _, d := range diags {
		if dv(d) {
			return true
		}
	}
	return false
}
