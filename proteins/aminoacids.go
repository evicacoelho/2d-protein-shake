package aminoacids

type Aminoacid struct {
	Letter    string
	Code      string
	Charge    int
	chemName  string
	chemClass string
}

var alanine = Aminoacid{
	Letter:    "A",
	Code:      "Ala",
	Charge:    -1,
	chemName:  "2-Aminopropanoic acid",
	chemClass: "aliphatic",
}

var isoleucine = Aminoacid{
	Letter:    "I",
	Code:      "Ile",
	Charge:    -1,
	chemName:  "(2S,3S)-2-amino-3-methylpentanoic acid",
	chemClass: "aliphatic",
}

var leucine = Aminoacid{
	Letter:    "L",
	Code:      "Leu",
	Charge:    -1,
	chemName:  "(S)-2-amino-4-methyl-pentanoic acid",
	chemClass: "aliphatic",
}

var methionine = Aminoacid{
	Letter:    "M",
	Code:      "Met",
	Charge:    -1,
	chemName:  "2-amino-4-(methylthio)butanoic acid",
	chemClass: "aliphatic",
}

var valine = Aminoacid{}

var Aminoacids [20]Aminoacid
