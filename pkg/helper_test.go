package pkg

import (
	"testing"
)

func TestVerifyEmailAddress(t *testing.T) {
	testCases := []struct {
		email    string
		expected bool
	}{
		{"user@example.com", true},
		{"user@@example.com", false},
		{"user@example", false},
		{"user@-example.com", false},
		{"user@", false},
		{"", false},
		{"user@.sub.domain.com", true},
		{"user@sub-domain.co.uk", true},
	}

	for _, tc := range testCases {
		t.Run(tc.email, func(t *testing.T) {
			err := VerifyEmailAddress(tc.email)
			if err != nil && tc.expected {
				t.Errorf("VerifyEmailAddress(%q) expected error, got no error", tc.email)
			}
		})
	}
}

func TestEntryPassword(t *testing.T) {
	slat, err := GenerateSalt()
	if err != nil {
		t.Errorf("generateSalt() expected no error, got %v", err)
	}
	expect1 := EncryptPassword(slat, "xianren68")
	expect2 := EncryptPassword(slat, "xianren68")
	if expect1 != expect2 {
		t.Errorf("EncryptPassword(%q) expected %q, got %q", slat, expect1, expect2)
	}

}
