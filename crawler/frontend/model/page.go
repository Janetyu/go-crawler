package model

type SearchResult struct {
	Hits int64
	Start int
	Items []interface{}
	//Items []types.Item
	Query string
	PrevFrom int
	NextFrom int
}
