package models

type RankBook struct {
	// rank title of books
	RankTitle string
	// book's name and url about 10
	BookInfos []BookInfo
}

type BookInfo struct {
	BookName string
	BookURL  string
}
