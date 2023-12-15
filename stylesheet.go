package main

const stylesheet = `
body {
    font-family: Arial, sans-serif;
    margin: 0;
    padding: 0;
}

article {
    display: flex;
    align-items: center;
    flex-direction: row;
	with: 100%;
    max-width: 900px;
    margin: 20px;
}

article:nth-child(odd) {
    flex-direction: row-reverse;
}

.article-img {
    max-width: 100%;
    height: auto;
    margin: 20px;
    border-radius: 8px;
}

h2 {
    font-size: 24px;
    padding: 0 20px;
}

p {
    font-size: 16px;
    padding: 0 20px;
}

nav {
    background-color: #333;
    color: #fff;
    position: fixed;
    width: 100%;
    top: 0;
}

ul {
    list-style-type: none;
    margin: 0;
    padding: 0;
    overflow: hidden;
    max-width: 800px;
}

li {
    float: left;
}

a {
    display: block;
    color: white;
    text-align: center;
    padding: 14px 16px;
    text-decoration: none;
}

a:hover {
    background-color: #ddd;
    color: black;
}

.content {
    padding: 80px;
    display: flex;
    align-items: center;
    flex-direction: column;
    background-color: #eee;
}

.card-container {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
	margin: 20px;
	flex: 0 0 25%;
}

.card {
    background-color: #eee;
    border-radius: 8px;
    padding: 20px;
    margin: 20px;
    margin-top: 50px;
}

.card-img {
	max-width: 500px;
	max-height: 500px;
}

.glossary-gif {
    display: flex;
    justify-content: center;
    align-items: center;
    max-width: 100%;
    max-height: 100%;
    align-self: center;
    margin: auto;
    margin-top: 100px;
}

/* Media Query for screens smaller than 600px */
@media only screen and (max-width: 600px) {
    .content {
        padding: 10px;
    }

    .card {
        flex: 0 0 100%;
    }

	article {
		max-width: 100%;
		flex-direction: column;
	}

	article:nth-child(odd) {
		flex-direction: column;
	}
}
`
