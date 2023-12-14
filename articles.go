package main

type Article struct {
	Title    string
	Body     string
	ImageUrl string
}

var Articles = []Article{{
	Title:    "Стаття 1",
	Body:     "Це тіло статті 1",
	ImageUrl: "https://picsum.photos/200/300",
}, {
	Title:    "Стаття 2",
	Body:     "Це тіло статті 2",
	ImageUrl: "https://picsum.photos/300/200",
}, {
	Title:    "Стаття 3",
	Body:     "Це тіло статті 3",
	ImageUrl: "https://picsum.photos/200/200",
}}
