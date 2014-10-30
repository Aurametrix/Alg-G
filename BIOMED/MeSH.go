// import "github.com/gnewton/jianGoMeSHi" 
// package jianGoMeSHi
// Go library to read [MEDLINE/PubMed Medical Subject Headings (MeSH)](http://www.nlm.nih.gov/mesh/) XML

const DESCRIPTOR_RECORD = "DescriptorRecord"

type Abbreviation  struct{
	Abbreviation string `xml:",chardata"`
}

//type ActiveMeSHYearList struct{
//	Year int
//}

type AllowableQualifier struct{
	QualifierReferredTo *QualifierReferredTo
	Abbreviation Abbreviation
}

type AllowableQualifiersList struct{
	AllowableQualifier []AllowableQualifier
}

type Concept struct{
	CASN1Name string `json:",omitempty"`
	ConceptName string `xml:">String"`
	ConceptRelationList ConceptRelationList `json:",omitempty"`
	ConceptUI string
	PreferredConceptYN string `xml:",attr"`
	RegistryNumber string `json:",omitempty"`
	RelatedRegistryNumberList RelatedRegistryNumberList  `json:",omitempty"`
	ScopeNote string `json:",omitempty"`
	SemanticTypeList SemanticTypeList `json:",omitempty"`
	TermList TermList `json:",omitempty"`
	TranslatorsEnglishScopeNote string `json:",omitempty"`
	TranslatorsScopeNote string `json:",omitempty"`
}

type ConceptList struct{
	Concept []Concept `json:",omitempty"`
}

type ConceptRelation struct{
	Concept1UI string
	Concept2UI string
	RelationName string `xml:",attr"`
}

type ConceptRelationList struct{
	ConceptRelation []ConceptRelation `json:",omitempty"`}


type Date struct{
	Year string `json:",omitempty"`
	Month string `json:",omitempty"`
	Day string `json:",omitempty"`
}


type DescriptorRecord struct {
	ActiveMeSHYearList []string `xml:">Year"`
	AllowableQualifiersList AllowableQualifiersList
	Annotation string `json:",omitempty"`
	ConceptList ConceptList
	ConsiderAlso string `json:",omitempty"`
	DateCreated Date `json:",omitempty"`
	DateEstablished Date `json:",omitempty"`
	DateRevised Date `json:",omitempty"`
	DescriptorName string `xml:"DescriptorName>String"`
	DescriptorUI string
	EntryCombinationList EntryCombinationList
	HistoryNote string `json:",omitempty"`
	OnlineNote string `json:",omitempty"`
	PharmacologicalActionList PharmacologicalActionList
	PreviousIndexingList PreviousIndexingList
	PublicMeSHNote string `json:",omitempty"`
	RecordOriginatorsList RecordOriginatorsList
	RunningHead string `json:",omitempty"`
	SeeRelatedList SeeRelatedList `json:",omitempty"`
	TreeNumberList TreeNumberList
}

type DescriptorRecordSet struct{
	DescriptorRecord []DescriptorRecord
	LanguageCode string `xml:",attr"`
}

type DescriptorReferredTo struct{
	DescriptorUI string
	DescriptorName string `xml:"DescriptorName>String"`
	DescriptorRecord *DescriptorRecord  `xml:"-" json:",omitempty"`
	Url string `json:",omitempty"`
}

type ECIN struct{
	DescriptorReferredTo *DescriptorReferredTo `json:",omitempty"`
}

type ECOUT struct{
	DescriptorReferredTo *DescriptorReferredTo `json:",omitempty"`
}

type EntryCombination struct{
	ECIN ECIN `json:",omitempty"`
	ECOUT ECOUT `json:",omitempty"`
}

type EntryCombinationList struct{
	EntryCombination []EntryCombination `json:",omitempty"`
}

type PharmacologicalAction struct{
	DescriptorReferredTo *DescriptorReferredTo `json:",omitempty"`
	PharmacologicalActionSubstanceList PharmacologicalActionSubstanceList `json:",omitempty"`
}

type PharmacologicalActionSubstanceList struct{
	Substance []Substance `json:",omitempty"`
}

type PharmacologicalActionList struct{
	PharmacologicalAction []PharmacologicalAction `json:",omitempty"`
}

type  PreviousIndexingList struct{
	PreviousIndexing []string `json:",omitempty"`
}

type QualifierReferredTo struct{
	QualifierUI string
	QualifierName string `xml:"QualifierName>String"`
	QualifierRecord *QualifierRecord `xml:"-" json:",omitempty"`
	Url string `json:",omitempty"`
}



type RecordOriginator struct{
	RecordOriginator string `json:",omitempty"`
	RecordMaintainer string `json:",omitempty"`
	RecordAuthorizer string `json:",omitempty"`
}

type RecordOriginatorsList struct{
	RecordOriginator []RecordOriginator `json:",omitempty"`
}

type RelatedRegistryNumber struct{
	RelatedRegistryNumber string  `xml:",chardata"`
}

type RelatedRegistryNumberList struct{
	RelatedRegistryNumber []RelatedRegistryNumber `json:",omitempty"`
}

type SeeRelatedDescriptor struct{
	DescriptorReferredTo *DescriptorReferredTo `json:",omitempty"`
}

type SeeRelatedList struct{
	SeeRelatedDescriptor []SeeRelatedDescriptor  `json:",omitempty"`
}

type SemanticType struct{
	SemanticTypeUI string `json:",omitempty"`
	SemanticTypeName string `json:",omitempty"`
}

type SemanticTypeList struct{
	SemanticType []SemanticType `json:",omitempty"`
}

type Term struct{
	Abbreviation string `json:",omitempty"`
	ConceptPreferredTermYN string `xml:",attr"`
	DateCreated Date `json:",omitempty"`
	EntryVersion string `json:",omitempty"`
	IsPermutedTermYN string `xml:",attr" json:",omitempty"`
	LexicalTag string `xml:",attr" json:",omitempty"`
	PrintFlagYN string `xml:",attr" json:",omitempty"`
	RecordPreferredTermYN string `xml:",attr" json:",omitempty"`
	SortVersion string `json:",omitempty"`
	String string `json:",omitempty"`
	TermNote string `json:",omitempty"`
	TermUI string `json:",omitempty"`
	ThesaurusIDlist ThesaurusIDlist `json:",omitempty"`
}

type TermList struct{
	Term []Term `json:",omitempty"`
}

type ThesaurusIDlist struct{
	ThesaurusID []string `json:",omitempty"`
}

type TreeNumber struct{
	TreeNumber string `xml:",chardata"`
	Url string `json:",omitempty"`
}

type TreeNumberList struct{
	TreeNumber []TreeNumber
}
