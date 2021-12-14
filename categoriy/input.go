package spotify

type SeveralBrowseCategoriesInput struct {
	Token   string
	Country string
	Limit   int
	Locale  string
	Offset  int
}

type SingleBrowseCategoriesInput struct {
	Token      string
	CategoryId string
	Country    string
	Locale     string
}
