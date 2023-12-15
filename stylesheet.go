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
    max-width: 900px;
    margin: 20px;
}

article:nth-child(odd) {
    flex-direction: row-reverse;
}

img {
    max-width: 100%;
    height: auto;
    margin-right: 20px;
    border-radius: 8px;
}

h2 {
    font-size: 24px;
    padding-left: 20px;
    padding-right: 20px;
}

p {
    font-size: 16px;
    padding-left: 20px;
    padding-right: 20px;
}

nav {
    background-color: #333;
    color: #fff;
    position: fixed;
    width: 100%;
}

ul {
    list-style-type: none;
    margin: 0;
    padding: 0;
    overflow: hidden;
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
`
