package mlphone

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMLPhone(t *testing.T) {
	phone := New()
	testString := []string{
		"ഫെയ്‌സ്ബുക്ക്",
		"മെസഞ്ചര്‍",
		"ആപ്ലിക്കേഷനെ",
		"ഇന്‍സ്റ്റഗ്രാമുമായി",
		"ബന്ധിപ്പിച്ചു.",
		"ബുധനാഴ്ച",
		"ഫെയ്‌സ്ബുക്കാണ്",
		"ഇക്കാര്യം",
		"പ്രഖ്യാപിച്ചത്.",
		"മെസഞ്ചറിലെ",
		"ആകര്‍ഷകമായ",
		"ഫീച്ചറുകള്‍",
		"ലഭ്യമാകുന്നതോടൊപ്പം",
		"ഇന്‍സ്റ്റാഗ്രാം",
		"ഉപയോക്താക്കളുമായി",
		"മെസഞ്ചര്‍",
		"ഉപയോക്താക്കള്‍ക്കും",
		"തിരിച്ചും",
		"ചാറ്റ്",
		"ചെയ്യാം.",
	}

	for _, v := range testString {
		out1, out2, out3 := phone.Encode(v)
		require.NotEmpty(t, out1)
		require.NotEmpty(t, out2)
		require.NotEmpty(t, out3)
	}
}
