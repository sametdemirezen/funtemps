package funfacts
type FunFacts struct {
	Terra []string
	Sun   []string
	Luna  []string
}
func GetFunFacts(str string) []string {
	funFacts := FunFacts{
		Terra: []string{
		"Høyeste temperatur målt på Jordens overflate.", 
		"Laveste temperatur målt på Jordens overflate.", 
		"Temperatur i Jordens indre kjerne"},

		Sun:   []string{
		"Temperatur i Solens kjerne.", 
		"Temperatur på ytre lag av Solen."},
		
		Luna:  []string{
		"Temperatur på Månens overflate om natten.", 
		"Temperatur på Månens overflate om dagen."},
	}

	if str == "terra" {
		return funFacts.Terra
	}
	if str == "sun" {
		return funFacts.Sun
	}
	if str == "luna" {
		return funFacts.Luna
	} else {
		return []string{}
	}
}


