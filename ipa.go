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
	Continuant
	Occlusive

	/* Non-pulmonic */
	Click
	Implosive
	Ejective
)

type Phoneme struct {
	Place Place
	Manner Manner
	Voiced bool
}

