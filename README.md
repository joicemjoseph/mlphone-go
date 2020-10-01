# MLPhone
![](https://gitlab.com/joice/mlphone-go/badges/master/coverage.svg)
![](https://gitlab.com/joice/mlphone-go/badges/master/pipeline.svg)

Rewriting [MLPhone](https://github.com/knadh/mlphone) to Go by porting [KNPhone](https://github.com/knadh/knphone).

MLphone is a phonetic algorithm for indexing Malayalam words by their pronunciation, like Metaphone for English. The algorithm generates three Romanized phonetic keys (hashes) of varying phonetic affinities for a given Malayalam word.

Full [documentation](http://nadh.in/code/mlphone) 

# Intro
MLphone is a phonetic algorithm for indexing Malayalam words by their pronunciation,
like Metaphone for English. The algorithm generates three Romanized phonetic keys (hashes) of varying
phonetic affinities for a given Malayalam word.

The	algorithm takes into account the context sensitivity of sounds, syntactic and
phonetic gemination, compounding, modifiers, and other known exceptions to produce
Romanized phonetic hashes of increasing phonetic affinity that are very faithful
to the pronunciation of the original Malayalam word.

* key0 = a broad phonetic hash comparable to a Metaphone key that doesn't account for hard sounds (ഷ, ണ..) or phonetic modifiers

* key1 = is a slightly more inclusive hash that accounts for hard sounds

* key2 = highly inclusive and narrow hash that accounts for hard sounds and phonetic modifiers


MLphone was created to aid spelling tolerant Malayalam word search, but may 
be useful in tasks like spell checking, word suggestion etc.

# Examples

|Word|key0|key1|key2|Transliteration|Metaphone|
|----|----|----|----|---------------|---------|
|നീലക്കുയില്‍|NLKYL|NLKYL|N4LK25Y4L|Neelakkuyil‍|NLKYL|
|മൃഗം|MRK3|MRK3|MRK3|Mrugam|MRKM|
|മ്രിഗം|MRK3|MRK3|MRK3|Mrigam|MRKM|
|ഉത്സവം|U0SV3|U0SV3|U0SV3|Uthsavam|U0SFM|
|ഉല്‍സവം|U0SV3|U0SV3|U0SV3|Ul‍savam|ULSFM|
|വാഹനം|VHN3|VHN3|VHN3|Vaahanam|FHNM|
|വിഹനനം|VHNN3|VHNN3|V4HNN3|Vihananam|FHNNM|
|രാഷ്ട്രീയം|RSTRY3|RS1TRY3|RS1TR4Y3|Raashtreeyam|RXTRYM|
|കണ്ണകി|KNK|KNK|KN2K4|Kannaki|KNK|
|കന്യക|KNYK|KNYK|KNYK|Kanyaka|KNYK|
|മനം|MN3|MN3|MN3|Manam|MNM|
|മണം|MN3|MN13|MN13|Manam|MNM|
|വിഭക്ത്യാഭാസം|VBK0YBS3|VBK0YBS3|V4BK0YBS3|Vibhakthyaabhaasam|FBHK0YBHSM|
|വലയം|VLY3|VLY3|VLY3|Valayam|FLYM|
|വളയം|VLY3|VL1Y3|VL1Y3|Valayam|FLYM|
|രഥം|R03|R03|R03|Ratham|R0M|
|രദം|R03|R03|R03|Radam|RTM|
|രത്തം|R03|R03|R03|Rattham|RTM|
|രധം|R03|R03|R03|Radham|RTHM|


# Usage
```go
import 'gitlab.com/joice/mlphone-go'
...
...
...

phone := mlphone.New()
val1, val2, val3 := phone.Encode("മണം")
...
...
...
```
# License
GNU GENERAL PUBLIC LICENSE v3
