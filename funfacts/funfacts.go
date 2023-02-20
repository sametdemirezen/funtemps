package funfacts


/**
  Implementer funfacts-funksjon:
    GetFunFacts(about string) []string
      hvor about kan ha en av tre testverdier, -
        sun, luna eller terra

  Sett inn alle Funfucts i en struktur
  type FunFacts struct {
      Sun []string
      Luna []string
      Terra []string
  }
*/

type FunFacts struct {
	Terra []string
	Sun   []string
	Luna  []string
}
func GetFunFacts(about string) []string {
	funFacts := FunFacts{
		Terra: []string{"Høyeste temperatur målt på Jordens overflate.", "Laveste temperatur målt på Jordens overflate.", "Temperatur i Jordens indre kjerne"},
		Sun:   []string{"Temperatur i Solens kjerne.", "Temperatur på ytre lag av Solen."},
		Luna:  []string{"Temperatur på Månens overflate om natten.", "Temperatur på Månens overflate om dagen."},
	}

	if about == "terra" {
		return funFacts.Terra
	}
	if about == "sun" {
		return funFacts.Sun
	}
	if about == "luna" {
		return funFacts.Luna
	} else {
		return []string{}
	}
}


