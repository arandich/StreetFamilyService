package model

type Catalog struct {
	ID          int32  `json:"id"`
	Category    int32  `json:"category"`
	Name        string `json:"name"`
	Value       int32  `json:"value"`
	Src         string `json:"src"`
	Weight      int32  `json:"weight"`
	Description string `json:"description"`
}
