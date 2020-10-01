package mlphone

import (
	"regexp"
	"strings"
)

var vowels = map[string]string{
	"അ": "A",
	"ആ": "A",
	"ഇ": "I",
	"ഈ": "I",
	"ഉ": "U",
	"ഊ": "U",
	"ഋ": "R",
	"എ": "E",
	"ഏ": "E",
	"ഐ": "AI",
	"ഒ": "O",
	"ഓ": "O",
	"ഔ": "O",
}

var consonants = map[string]string{
	"ക": "K",
	"ഖ": "K",
	"ഗ": "K",
	"ഘ": "K",
	"ങ": "NG",
	"ച": "C",
	"ഛ": "C",
	"ജ": "J",
	"ഝ": "J",
	"ഞ": "NJ",
	"ട": "T",
	"ഠ": "T",
	"ഡ": "T",
	"ഢ": "T",
	"ണ": "N1",
	"ത": "0",
	"ഥ": "0",
	"ദ": "0",
	"ധ": "0",
	"ന": "N",
	"പ": "P",
	"ഫ": "F",
	"ബ": "B",
	"ഭ": "B",
	"മ": "M",
	"യ": "Y",
	"ര": "R",
	"ല": "L",
	"വ": "V",
	"ശ": "S1",
	"ഷ": "S1",
	"സ": "S",
	"ഹ": "H",
	"ള": "L1",
	"ഴ": "Z",
	"റ": "R1",
}

var chillus = map[string]string{
	"ൽ": "L",
	"ൾ": "L1",
	"ൺ": "N1",
	"ൻ": "N",
	"ർ": "R1",
	"ൿ": "K",
}

var compounds = map[string]string{
	"ക്ക": "K2",
	"ഗ്ഗ": "K",
	"ങ്ങ": "NG",
	"ച്ച": "C2",
	"ജ്ജ": "J",
	"ഞ്ഞ": "NJ",
	"ട്ട": "T2",
	"ണ്ണ": "N2",
	"ത്ത": "0",
	"ദ്ദ": "D",
	"ദ്ധ": "D",
	"ന്ന": "NN",
	"ന്ത": "N0",
	"ങ്ക": "NK",
	"ണ്ട": "N1T",
	"ബ്ബ": "B",
	"പ്പ": "P2",
	"മ്മ": "M2",
	"യ്യ": "Y",
	"ല്ല": "L2",
	"വ്വ": "V",
	"ശ്ശ": "S1",
	"സ്സ": "S",
	"ള്ള": "L12",
	"ഞ്ച": "NC",
	"ക്ഷ": "KS1",
	"മ്പ": "MP",
	"റ്റ": "T",
	"ന്റ": "NT",
	"്രി": "R",
	"്രു": "R",
}

var modifiers = map[string]string{
	"ാ": "",
	"ഃ": "",
	"്": "",
	"ൃ": "R",
	"ം": "3",
	"ി": "4",
	"ീ": "4",
	"ു": "5",
	"ൂ": "5",
	"െ": "6",
	"േ": "6",
	"ൈ": "7",
	"ൊ": "8",
	"ോ": "8",
	"ൌ": "9",
	"ൗ": "9",
}

var (
	regexKey0, _         = regexp.Compile(`[1,2,4-9]`)
	regexKey1, _         = regexp.Compile(`[2,4-9]`)
	regexNonMalayalam, _ = regexp.Compile(`[\P{Malayalam}]`)
	regexAlphaNum, _     = regexp.Compile(`[^0-9A-Z]`)
)

// MLPhone is the Malayalam-phone tokenizer.
type MLPhone struct {
	modCompounds  *regexp.Regexp
	modConsonants *regexp.Regexp
	modVowels     *regexp.Regexp
}

// New returns a new instance of the KNPhone tokenizer.
func New() *MLPhone {
	var (
		glyphs []string
		mods   []string
		kn     = &MLPhone{}
	)

	// modifiers.
	for k := range modifiers {
		mods = append(mods, k)
	}

	// compounds.
	for k := range compounds {
		glyphs = append(glyphs, k)
	}
	kn.modCompounds, _ = regexp.Compile(`((` + strings.Join(glyphs, "|") + `)(` + strings.Join(mods, "|") + `))`)

	// consonants.
	glyphs = []string{}
	for k := range consonants {
		glyphs = append(glyphs, k)
	}
	kn.modConsonants, _ = regexp.Compile(`((` + strings.Join(glyphs, "|") + `)(` + strings.Join(mods, "|") + `))`)

	// vowels.
	glyphs = []string{}
	for k := range vowels {
		glyphs = append(glyphs, k)
	}
	kn.modVowels, _ = regexp.Compile(`((` + strings.Join(glyphs, "|") + `)(` + strings.Join(mods, "|") + `))`)

	return kn
}

// Encode encodes a unicode Malayalm string to its Roman MLPhone hash.
// Ideally, words should be encoded one at a time, and not as phrases
// or sentences.
func (ml *MLPhone) Encode(input string) (string, string, string) {
	// key2 accounts for hard and modified sounds.
	key2 := ml.process(input)

	// key1 loses numeric modifiers that denote phonetic modifiers.
	key1 := regexKey1.ReplaceAllString(key2, "")

	// key0 loses numeric modifiers that denote hard sounds, doubled sounds,
	// and phonetic modifiers.
	key0 := regexKey0.ReplaceAllString(key2, "")

	return key0, key1, key2
}

func (ml *MLPhone) process(input string) string {
	// Remove all non-malayalam characters.
	input = regexNonMalayalam.ReplaceAllString(strings.Trim(input, ""), "")

	// All character replacements are grouped between { and } to maintain
	// separatability till the final step.

	// Replace and group modified compounds.
	input = ml.replaceModifiedGlyphs(input, compounds, ml.modCompounds)

	// Replace and group unmodified compounds.
	for k, v := range compounds {
		input = strings.ReplaceAll(input, k, `{`+v+`}`)
	}

	// Replace and group modified consonants and vowels.
	input = ml.replaceModifiedGlyphs(input, consonants, ml.modConsonants)
	input = ml.replaceModifiedGlyphs(input, vowels, ml.modVowels)

	// Replace and group unmodified consonants.
	for k, v := range consonants {
		input = strings.ReplaceAll(input, k, `{`+v+`}`)
	}

	// Replace and group unmodified vowels.
	for k, v := range vowels {
		input = strings.ReplaceAll(input, k, `{`+v+`}`)
	}

	// Replace all modifiers.
	for k, v := range modifiers {
		input = strings.ReplaceAll(input, k, v)
	}

	// Remove non alpha numeric characters (losing the bracket grouping).
	return regexAlphaNum.ReplaceAllString(input, "")
}

func (ml *MLPhone) replaceModifiedGlyphs(input string, glyphs map[string]string, r *regexp.Regexp) string {
	for _, matches := range r.FindAllStringSubmatch(input, -1) {
		for _, m := range matches {
			if rep, ok := glyphs[m]; ok {
				input = strings.ReplaceAll(input, m, rep)
			}
		}
	}
	return input
}
