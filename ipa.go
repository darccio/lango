package main

type Class int
const (
	Labial Class = iota
	Coronal
	Guttural
)

type Position int
const (
	Front Position = iota
	Back
)

type Intensity int
const (
	Acute Intensity = iota
	Grave
)

type Articulator int
const (
	/* Active */
	LowerLip Articulator = iota /* Labial */
	Laminal /* Tongue blade */
	Apical /* Tongue tip */
	Subapical /* Underside of tongue */
	Dorsal /* Tongue body */
	Radical /* Tongue root */
	Laryngeal /* Larynx */

	/* Passive */
	UpperLip
	UpperTeeth
	AlveolarRidge
	Postalveolar
	HardPalate
	SoftPalate
	Uvula
	Pharynx
	Epiglottis
	Glottis
)

type Place struct {
	Major Class
	Intensity Intensity
	Active Articulator
	Passive Articulator
}
var (
	Bilabial = &Place{Labial, Grave, LowerLip, UpperLip}
	Labiodental = &Place{Labial, Grave, LowerLip, UpperTeeth}
	Linguolabial = &Place{Coronal, Grave, Laminal, UpperLip}
	Interdental = &Place{Coronal, Acute, Laminal, UpperTeeth}
	Alveolar = &Place{Coronal, Acute, Laminal, AlveolarRidge} /* Denti-alveolar, laminal alveolar */
	PalatoAlveolar = &Place{Coronal, Acute, Laminal, Postalveolar}
	Dental = &Place{Coronal, Acute, Apical, UpperTeeth}
	Retroflex = &Place{Coronal, Acute, Apical, HardPalate}
	AlveoloPalatal = &Place{Guttural, Acute, Dorsal, Postalveolar}
	Palatal = &Place{Guttural, Acute, Dorsal, HardPalate}
	Velar = &Place{Guttural, Grave, Dorsal, SoftPalate}
	Uvular = &Place{Guttural, Grave, Dorsal, Uvula}
	Pharyngeal = &Place{Guttural, Grave, Radical, Pharynx}
	Epiglottal = &Place{Guttural, Grave, Laryngeal, Epiglottis}
	Glottal = &Place{Guttural, Grave, Laryngeal, Glottis}
)

func (p *Place) Position() (position Position) {
	position = Front
	if p.Major == Guttural {
		position = Back
	}
	return position
}

type Manner int
const (
	/* Pulmonic */
	Nasal Manner = iota
	Stop
	SibilantFricative
	NonSibilantFricative
	Approximant
	Flap
	Trill
	LateralFricative
	LateralApproximant
	LateralFlap
	Affricate

	/* Co-articulated consonants */
	Occlusive
	Continuant

	/* Non-pulmonic */
	Click
	Implosive
	Ejective
)

type Phoneme struct {
	Place      *Place
	Manner     Manner
	Voiced     bool
	Labialized bool
	Nasalized  bool
}

func NewPhoneme(place *Place, manner Manner, voiced bool) *Phoneme {
	return &Phoneme{place, manner, voiced, false, false}
}

func NewLabializedPhoneme(place *Place, manner Manner, voiced bool) *Phoneme {
	return &Phoneme{place, manner, voiced, true, false}
}

func NewNasalizedPhoneme(place *Place, manner Manner, voiced bool) *Phoneme {
	return &Phoneme{place, manner, voiced, false, true}
}

func (p *Phoneme) Pulmonic() bool {
	switch p.Manner {
		case Click, Implosive, Ejective:
			return false
	}
	return true
}

func (p *Phoneme) Occlusive() bool {
	// Continuant nasal phonemes don't exist in real languages.
	switch p.Manner {
		case Occlusive, Nasal, Stop, Affricate, Implosive, Ejective, Click:
			return true
	}
	return false
}

func (p *Phoneme) Oral() bool {
	return p.Manner != Nasal
}

func (p *Phoneme) Central() bool {
	switch p.Manner {
		case SibilantFricative, NonSibilantFricative, Trill, Flap, Approximant:
			return true
	}
	return false
}

func (p *Phoneme) Lateral() bool {
	switch p.Manner {
		case LateralFricative, LateralApproximant, LateralFlap:
			return true
	}
	return false
}

func (p *Phoneme) NoCentralLateralDichotomy() bool {
	return !p.Central() && !p.Lateral()
}

// TODO reflection test
var (
	BilNas = NewPhoneme(Bilabial, Nasal, false)
	BilNasV = NewPhoneme(Bilabial, Nasal, true)
	LabNas = NewPhoneme(Labiodental, Nasal, false)
	LabNasV = NewPhoneme(Labiodental, Nasal, true)
	DenNas = NewPhoneme(Dental, Nasal, false)
	DenNasV = NewPhoneme(Dental, Nasal, true)
	AlvNas = NewPhoneme(Alveolar, Nasal, false)
	AlvNasV = NewPhoneme(Alveolar, Nasal, true)
	PaaNasV = NewPhoneme(PalatoAlveolar, Nasal, true)
	RetNas = NewPhoneme(Retroflex, Nasal, false)
	RetNasV = NewPhoneme(Retroflex, Nasal, true)
	AlpNasV = NewPhoneme(AlveoloPalatal, Nasal, true)
	PalNas = NewPhoneme(Palatal, Nasal, false)
	PalNasV = NewPhoneme(Palatal, Nasal, true)
	VelNas = NewPhoneme(Velar, Nasal, false)
	VelNasV = NewPhoneme(Velar, Nasal, true)
	UvuNas = NewPhoneme(Uvular, Nasal, false)
	UvuNasV = NewPhoneme(Uvular, Nasal, true)
	// Nasal palatal approximant
	PalNAppV = NewNasalizedPhoneme(Palatal, Approximant, true)
	// Labio-velar approximant
	VelLAppV = NewLabializedPhoneme(Velar, Approximant, true)
	// Nasal labio-velar approximant
	VelNLAppV = &Phoneme{Velar, Approximant, true, true, true}
	// Voiceless nasal glottal approximant
	GloNApp = NewNasalizedPhoneme(Glottal, Approximant, false)
	GloNStop = NewNasalizedPhoneme(Glottal, Stop, false)
	GloStop = NewPhoneme(Glottal, Stop, false)
)

var Phonemes = map[string]*Phoneme {
	// Nasal
	"m̥": BilNas,
	"m": BilNasV,
	"ɱ̊": LabNas,
	"ɱ": LabNasV,
	"n̪̊": DenNas,
	"n̪": DenNasV,
	"n̥": AlvNas,
	"n": AlvNasV,
	"n̠": PaaNasV,
	"ɳ̊": RetNas,
	"ɳ": RetNasV,
	"ɲ̟": AlpNasV,
	"ɲ̥": PalNas,
	"ɲ": PalNasV,
	"ŋ̊": VelNas,
	"ŋ": VelNasV,
	"ɴ̥": UvuNas,
	"ɴ": UvuNasV,
	// Approximant
	"w": VelLAppV,
	// Stop
	"ʔ": GloStop,
	// Nasalized
	"ʔ̃": GloNStop,
	// Nasal glides
	"ȷ̃": PalNAppV,
	"w̃": VelNLAppV,
	"h̃": GloNApp,
	"": nil,
}

var Equivalences = map[string][]string {
	"n̪": { "n" },
	"n̪̊": { "n̥" },
	"m̪": { "ɱ" },
	"ɳ̥": { "ɳ̊" },
	"n̠ʲ": { "ɲ̟" },
	"ȵ": { "ɲ̟" },
	"": { "" },
}
