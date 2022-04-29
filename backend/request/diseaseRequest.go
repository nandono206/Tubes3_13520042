package request



type DiseaseRequest struct {
	Nama 	string `json:"nama" validate:"required"`
	// Sequence string `json:"none"`
	Dna_seq  string `json:"dna_seq"`
	
}