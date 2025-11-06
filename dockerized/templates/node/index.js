(function () {
	const express = require('express')
	const app = express()

	// Health endpoint
	app.get('/health', (req, res) => res.json({ status: 'ok' }))

	// Plugin route matching the Go plugin: /hello
	app.get('/hello', (req, res) => {
		res.send('Hello from Node plugin!')
	})

	const port = process.env.PORT || 8081
	app.listen(port, () => console.log(`hello-node listening on ${port}`))
})()

