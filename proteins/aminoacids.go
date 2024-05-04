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

var valine = Aminoacid{
	Letter:    "V",
	Code:      "Val",
	Charge:    -1,
	chemName:  "(2S)-2-amino-3-methylbutanoic acid",
	chemClass: "aliphatic",
}

var cysteine = Aminoacid{
	Letter:    "C",
	Code:      "Cys",
	Charge:    0,
	chemName:  "2-amino-3-mercaptopropanoic acid",
	chemClass: "sulfur-containing",
}

var phenylalanine = Aminoacid{
	Letter:    "F",
	Code:      "Phe",
	Charge:    0,
	chemName:  "2-amino-3-phenylpropanoic acid",
	chemClass: "aromatic",
}

var tyrosine = Aminoacid{
	Letter:    "Y",
	Code:      "Tyr",
	Charge:    0,
	chemName:  "2-amino-3-(4-hydroxyphenyl)propanoic acid",
	chemClass: "aromatic",
}

var tryptophan = Aminoacid{
	Letter:    "W",
	Code:      "Trp",
	Charge:    0,
	chemName:  "2-amino-3-(1H-indol-3-yl)propanoic acid",
	chemClass: "aromatic",
}

var proline = Aminoacid{
	Letter:    "P",
	Code:      "Pro",
	Charge:    0,
	chemName:  "(2S)-pyrrolidine-2-carboxylic acid",
	chemClass: "cyclic",
}

var serine = Aminoacid{
	Letter:    "S",
	Code:      "Ser",
	Charge:    0,
	chemName:  "2-amino-3-hydroxypropanoic acid",
	chemClass: "hydroxyl-containing",
}

var threonine = Aminoacid{
	Letter:    "T",
	Code:      "Thr",
	Charge:    0,
	chemName:  "(2S,3R)-2-amino-3-hydroxybutanoic acid",
	chemClass: "hydroxyl-containing",
}

var asparagine = Aminoacid{
	Letter:    "N",
	Code:      "Asn",
	Charge:    0,
	chemName:  "2-amino-3-carbamoylpropanoic acid",
	chemClass: "amide-containing",
}

var glutamine = Aminoacid{
	Letter:    "Q",
	Code:      "Gln",
	Charge:    0,
	chemName:  "2-amino-4-carbamoylbutanoic acid",
	chemClass: "amide-containing",
}

var histidine = Aminoacid{
	Letter:    "H",
	Code:      "His",
	Charge:    1,
	chemName:  "(2S)-2-amino-3-(1H-imidazol-5-yl)propanoic acid",
	chemClass: "basic",
}

var arginine = Aminoacid{
	Letter:    "R",
	Code:      "Arg",
	Charge:    1,
	chemName:  "(2S)-2-amino-5-(diaminomethylideneamino)pentanoic acid",
	chemClass: "basic",
}

var lysine = Aminoacid{
	Letter:    "K",
	Code:      "Lys",
	Charge:    1,
	chemName:  "(2S)-2,6-diaminohexanoic acid",
	chemClass: "basic",
}

var glycine = Aminoacid{
	Letter:    "G",
	Code:      "Gly",
	Charge:    0,
	chemName:  "2-aminoacetic acid",
	chemClass: "aliphatic",
}

var asparticAcid = Aminoacid{
	Letter:    "D",
	Code:      "Asp",
	Charge:    -1,
	chemName:  "2-aminobutanedioic acid",
	chemClass: "acidic",
}

var glutamicAcid = Aminoacid{
	Letter:    "E",
	Code:      "Glu",
	Charge:    -1,
	chemName:  "2-aminopentanedioic acid",
	chemClass: "acidic",
}
