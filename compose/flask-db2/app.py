from flask import Flask
import mysql.connector
import os

app = Flask(__name__)

# Adatbázis kapcsolati adatok környezeti változókból
DB_HOST = os.getenv("DB_HOST", "mydb.myflasknet")
DB_USER = os.getenv("DB_USER", "flaskuser")
DB_PASSWORD = os.getenv("DB_PASSWORD", "flaskpass")
DB_NAME = os.getenv("DB_NAME", "flaskdb")

def connect_db():
    return mysql.connector.connect(
        host=DB_HOST,
        user=DB_USER,
        password=DB_PASSWORD,
        database=DB_NAME
    )

@app.route('/')
def index():
    db = connect_db()
    cursor = db.cursor()

    # Tábla létrehozása (ha nem létezik)
    cursor.execute("""
        CREATE TABLE IF NOT EXISTS users (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(100)
        )
    """)

    # Alapértelmezett adatok beszúrása
    cursor.execute("INSERT INTO users (name) VALUES ('Alice'), ('Bob')")
    db.commit()

    # Adatok lekérdezése
    cursor.execute("SELECT * FROM users")
    users = cursor.fetchall()

    cursor.close()
    db.close()

    return "<br>".join([f"ID: {user[0]}, Name: {user[1]}" for user in users])

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, debug=True)

