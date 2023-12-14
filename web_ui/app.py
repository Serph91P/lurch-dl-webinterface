from quart import Quart, websocket, render_template
import asyncio
import json
import subprocess

app = Quart(__name__, static_folder='/webapp/static', static_url_path='/static')
download_process = None  # Globale Variable für den Download-Prozess
last_url = None  # Globale Variable für die zuletzt verwendete URL

@app.route('/')
async def index():
    return await render_template('index.html')

# Funktion, um verfügbare Formate und Kapitel abzurufen
async def get_lurch_dl_data(url):
    format_process = await asyncio.create_subprocess_exec(
        '/usr/local/bin/lurch-dl', '--list-formats', '--json',
        stdout=subprocess.PIPE)
    chapter_process = await asyncio.create_subprocess_exec(
        '/usr/local/bin/lurch-dl', '--url', url, '--list-chapters', '--json',
        stdout=subprocess.PIPE)

    format_output = await format_process.stdout.read()
    chapter_output = await chapter_process.stdout.read()

    formats = json.loads(format_output.decode()) if format_output else []
    chapters = json.loads(chapter_output.decode()) if chapter_output else []

    return formats, chapters

@app.route('/get-options-data')
async def get_options_data():
    if last_url:
        formats, chapters = await get_lurch_dl_data(last_url)
        return json.dumps({'formats': formats, 'chapters': chapters})
    else:
        return json.dumps({'error': 'Keine URL angegeben'})

@app.websocket('/ws')
async def ws():
    global download_process, last_url
    while True:
        data = await websocket.receive()
        message = json.loads(data)

        if message.get("action") in ["start", "continue"]:
            url = message.get("url")
            if url:
                last_url = url  # Aktualisiere die zuletzt verwendete URL
                formats, chapters = await get_lurch_dl_data(url)
            command = ["/usr/local/bin/lurch-dl", "--url", message.get("url"), "--json"]

            # Füge weitere Optionen hinzu
            if message.get("chapter"):
                command += ["--chapter", message["chapter"]]
            if message.get("format"):
                command += ["--format", message["format"]]
            if message.get("output"):
                command += ["--output", message["output"]]
            if message.get("start"):
                command += ["--start", message["start"]]
            if message.get("stop"):
                command += ["--stop", message["stop"]]
            if message.get("overwrite"):
                command += ["--overwrite"]
            if message.get("maxRate"):
                command += ["--max-rate", message["maxRate"]]

            # Starte oder setze den Download fort
            download_process = await asyncio.create_subprocess_exec(
                *command,
                stdout=subprocess.PIPE, stderr=subprocess.PIPE)
            asyncio.create_task(send_output_to_frontend(websocket, download_process))

        elif message.get("action") == "stop":
            if download_process:
                download_process.terminate()
                download_process = None

@app.route('/options')
async def options():
    return await render_template('options.html')

async def send_output_to_frontend(ws, process):
    while True:
        output = await process.stdout.readline()
        if output:
            output_data = output.decode().strip()
            try:
                json_data = json.loads(output_data)
                await ws.send(json.dumps(json_data))
            except json.JSONDecodeError:
                await ws.send(json.dumps({"type": "output", "data": output_data}))
        else:
            break

if __name__ == '__main__':
    app.run()
