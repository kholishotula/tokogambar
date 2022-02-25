package entity

type SimilarImage struct {
	Filename        string  `json:"filename"`
	SimilarityScore float64 `json:"similarity_score"`
}
