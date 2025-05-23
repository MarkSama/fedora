from flask import Flask
import os

app = Flask(__name__)

@app.route('/')
def hello():
    return os.getenv("WELCOME_MESSAGE", "Hello, World!")

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=5000)

