from flask import Flask, jsonify
import os

app = Flask(__name__)


@app.route('/health')
def health():
	return jsonify(status='ok')


@app.route('/bye')
def bye():
	return jsonify(message='Goodbye from Python plugin!')


if __name__ == '__main__':
	port = int(os.getenv('PORT', 8082))
	app.run(host='0.0.0.0', port=port)

