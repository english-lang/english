package tag

// WordType is the type of english word
type WordType string

const (

	// LeftRoundBracket Tag
	LeftRoundBracket = WordType("(")

	// RightRoundBracket Tag
	RightRoundBracket = WordType(")")

	// Comma Tag
	Comma = WordType(",")

	// Colon Tag
	Colon = WordType(":")

	// Period Tag
	Period = WordType(".")

	// OpeningQutation Tag
	OpeningQutation = WordType("``")
	// ClosingQutation Tag
	ClosingQutation = WordType("''")

	// NumberSign Tag
	NumberSign = WordType("#")
	// Currency Tag
	Currency = WordType("$")

	// Conjunction Tag
	// bring the different elements of (a complex activity or organization) into a relationship that will ensure efficiency or harmony.
	Conjunction = WordType("CC")

	// CardinalNumber Tag
	// a number denoting quantity (one, two, three, etc.), as opposed to an ordinal number (first, second, third, etc.).
	CardinalNumber = WordType("CD")

	// Determiner tag
	// a modifying word that determines the kind of reference a noun or noun group has, for example a , the , every.
	Determiner = WordType("DT")

	// ExistentialThere Tag
	ExistentialThere = WordType("EX")

	// ForeignWord Tag
	ForeignWord = WordType("FW")
	// ConjunctionSubordinatingOrPreposition Tag

	ConjunctionSubordinatingOrPreposition = WordType("IN")
	// Adjective Tag
	// a word or phrase naming an attribute, added to or grammatically related to a noun to modify or describe it.
	Adjective = WordType("JJ")

	// ComparativeAdjective Tag
	ComparativeAdjective = WordType("JJR")

	// SuperlativeAdjective Tag
	SuperlativeAdjective = WordType("JJS")

	// ListItemMarker Tag
	ListItemMarker = WordType("LS")

	// VerbModalAuxiliary Tag
	VerbModalAuxiliary = WordType("MD")

	// NounSingularOrMass Tag
	NounSingularOrMass = WordType("NN")

	// NounProperSingular Tag
	NounProperSingular = WordType("NNP")

	// NounProperPlural Tag
	NounProperPlural = WordType("NNPS")

	// NounPlural Tag
	NounPlural = WordType("NP")

	// Predeterminer Tag
	// a word or phrase that occurs before a determiner, typically quantifying the noun phrase, for example both or a lot of.
	Predeterminer = WordType("PDT")

	// PossessiveEnding Tag
	PossessiveEnding = WordType("POS")

	// PersonalPronoun Tag
	// each of the pronouns in English ( I , you , he , she , it , we , they , me , him , her , us , and them ) comprising a set that shows contrasts of person, gender, number, and case.
	PersonalPronoun = WordType("PRP")

	// PossessivePronoun Tag
	// a pronoun indicating possession, for example mine , yours , hers , theirs.
	PossessivePronoun = WordType("PRP$")

	Adverb            = WordType("RB")
	ComparativeAdverb = WordType("RBR")
	SuperlativeAdverb = WordType("RBS")
	ParticleAdverb    = WordType("RP")

	// Symbol Tag
	// a thing that represents or stands for something else, especially a material object representing something abstract.
	Symbol = WordType("SYM")

	// InfinitivalTo Tag
	InfinitivalTo = WordType("TO")

	// BaseFormVerb Tag
	// a word used to describe an action, state, or occurrence, and forming the main part of the predicate of a sentence, such as hear , become , happen.
	BaseFormVerb = WordType("VB")

	PastTenseVerb                     = WordType("VBD")
	GerundOrPresentParticipleVerb     = WordType("VBG")
	PastParticipleVerb                = WordType("VBN")
	NonThirdPersonSingularpresentVerb = WordType("VBP")
	ThirdPersonSingularpresentVerb    = WordType("VBZ")
	WHDeterminer                      = WordType("WDT")
	PersonalWHPronoun                 = WordType("WP")
	PossessiveWHPronoun               = WordType("WP")
	WHAdverb                          = WordType("WRB")
)
