// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateUserInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

type GraphData struct {
	Datasets []*GraphDataset `json:"datasets"`
}

type GraphDataset struct {
	Date  int `json:"date"`
	Value int `json:"value"`
}

type HistoricalData struct {
	Translations *GraphData `json:"translations"`
}

type Mutation struct {
}

type Query struct {
}

type Stats struct {
	Users         int             `json:"users"`
	Organizations int             `json:"organizations"`
	Projects      int             `json:"projects"`
	Translations  int             `json:"translations"`
	Historical    *HistoricalData `json:"historical"`
}

type User struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Role    *UserRole `json:"role"`
	Country *string   `json:"country,omitempty"`
}

type UserInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserRole struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
