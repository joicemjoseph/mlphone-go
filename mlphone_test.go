package mlphone

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testVal struct {
	word     string
	expected expected
}

type expected struct {
	val1, val2, val3 string
}

func TestMLPhone(t *testing.T) {
	phone := New()
	testStrings := []testVal{
		{
			word: "നീലക്കുയില്‍",
			expected: expected{"NLKYL",
				"NLKYL",
				"N4LK25Y4L",
			},
		},
		{
			word: "മൃഗം",
			expected: expected{
				"MRK3",
				"MRK3",
				"MRK3",
			},
		},
		{
			word: "മ്രിഗം",
			expected: expected{
				"MRK3",
				"MRK3",
				"MRK3",
			},
		},
		{
			word: "ഉത്സവം",
			expected: expected{
				"U0SV3",
				"U0SV3",
				"U0SV3",
			},
		},
		{
			word: "ഉല്‍സവം",
			expected: expected{
				"U0SV3",
				"U0SV3",
				"U0SV3",
			},
		}, {
			word: "വാഹനം",
			expected: expected{
				"VHN3",
				"VHN3",
				"VHN3",
			},
		},
		{
			word: "വിഹനനം",
			expected: expected{
				"VHNN3",
				"VHNN3",
				"V4HNN3",
			},
		},
		{
			word: "രാഷ്ട്രീയം",
			expected: expected{
				"RSTRY3",
				"RS1TRY3",
				"RS1TR4Y3",
			},
		},
		{
			word: "കണ്ണകി",
			expected: expected{
				"KNK",
				"KNK",
				"KN2K4",
			},
		},
		{
			word: "കന്യക",
			expected: expected{
				"KNYK",
				"KNYK",
				"KNYK",
			},
		},
		{
			word: "മനം",
			expected: expected{
				"MN3",
				"MN3",
				"MN3",
			},
		},
		{
			word: "മണം",
			expected: expected{
				"MN3",
				"MN13",
				"MN13",
			},
		},
		{
			word: "വിഭക്ത്യാഭാസം",
			expected: expected{
				"VBK0YBS3",
				"VBK0YBS3",
				"V4BK0YBS3",
			},
		},
		{
			word: "വലയം",
			expected: expected{
				"VLY3",
				"VLY3",
				"VLY3",
			},
		},
		{
			word: "വളയം",
			expected: expected{
				"VLY3",
				"VL1Y3",
				"VL1Y3",
			},
		},
		{
			word: "രഥം",
			expected: expected{
				"R03",
				"R03",
				"R03",
			},
		},
		{
			word: "രദം",
			expected: expected{
				"R03",
				"R03",
				"R03",
			},
		},
		{
			word: "രത്തം",
			expected: expected{
				"R03",
				"R03",
				"R03",
			},
		},
		{
			word: "രധം",
			expected: expected{
				"R03",
				"R03",
				"R03",
			},
		},
	}
	for _, v := range testStrings {
		out1, out2, out3 := phone.Encode(v.word)
		require.Equal(t, v.expected.val1, out1)
		require.Equal(t, v.expected.val2, out2)
		require.Equal(t, v.expected.val3, out3)
	}
}
