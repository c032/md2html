package main

const htmlTemplate = `
<!doctype html>
<html>
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	{{if .Title}}
	<title>{{.Title}}</title>
	{{end}}
	<style>
	:root {
		--phi: 1.618033988749895;

		--base-font-size: 16px;

		--h3: calc(var(--base-font-size));
		--h2: calc(var(--base-font-size) * var(--phi));
		--h1: calc(var(--base-font-size) * var(--phi) * var(--phi));
	}

	* {
		box-sizing: border-box;
	}

	html, body {
		font-size: var(--base-font-size);
		margin: 0;
		padding: 0;
	}

	body {
		background: #fff;
		color: #333;
		font-family: 'Fira Sans', Ubuntu, sans-serif;
		line-height: 1.6rem;
	}

	main {
		background: #fff;
		margin: 0 auto;
		max-width: 1300px;
		padding: 5rem calc(2% * var(--phi));
		width: 70%;
	}

	h1,
	h2,
	h3,
	h4,
	h5,
	h6 {
		color: #596dff;
		margin: 0;
		padding: 1.25em 0 0.25em 0;
	}

	hr {
		border: none;
		margin: 3rem 0;

		border-bottom: 1px solid #ddd;
	}

	hr + h1,
	hr + h2,
	hr + h3,
	hr + h4,
	hr + h5,
	hr + h6 {
		padding-top: 0;
	}

	h1 {
		font-size: var(--h1);
		font-weight: normal;
	}

	main > h1:first-child {
		padding-top: 0;
	}

	h2 {
		font-size: var(--h2);
		font-weight: normal;
	}

	h3 {
		font-size: var(--h3);
		font-weight: bold;
	}

	p,
	ul {
		margin: 1.5rem 0;
	}

	ul {
		list-style-position: outside;
		margin: 0;
		padding: 0;

		margin-left: 2em;
	}

	ul li + li {
		margin-top: 1em;
	}

	ul li {
		padding-left: 2em;
	}
	</style>
</head>
<body>
	<main>
		{{.Content}}
	</main>
</body>
</html>
`
