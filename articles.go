package main

type Article struct {
	Title    string
	Body     string
	ImageUrl string
}

var Articles = []Article{
	{
		Title:    "Article 1",
		Body:     "This is the body of article 1",
		ImageUrl: "https://picsum.photos/200/300",
	},
}
