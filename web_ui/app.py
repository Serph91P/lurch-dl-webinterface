from flask import Flask, request, render_template
import subprocess

app = Flask(__name__)

@app.route('/', methods=['GET', 'POST'])
def index():
    if request.method == 'POST':
        url = request.form.get('url')
        if url:
            command = [
                "/usr/local/bin/lurch-dl", "--url", url,
                # Füge hier weitere statische Parameter hinzu, falls benötigt
            ]
            try:
                subprocess.run(command, check=True)
                return "Download gestartet für URL: " + url
            except subprocess.CalledProcessError as e:
                return "Fehler beim Ausführen des Befehls: " + str(e)
    return render_template('index.html')

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
