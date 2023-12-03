from quart import Quart, websocket, render_template
import asyncio
import json
import subprocess

app = Quart(__name__, static_folder='/webapp/static', static_url_path='/static')
download_process = None  # Globale Variable für den Download-Prozess

@app.route('/')
async def index():
    return await render_template('index.html')

@app.websocket('/ws')
async def ws():
    global download_process
    while True:
        data = await websocket.receive()
        message = json.loads(data)

        if message.get("action") == "start":
            url = message.get("url")
            # Starte den Download-Prozess mit --json Flag
            download_process = await asyncio.create_subprocess_exec(
                "/usr/local/bin/lurch-dl", "--url", url, "--json",
                stdout=subprocess.PIPE, stderr=subprocess.PIPE)

            asyncio.create_task(send_output_to_frontend(websocket, download_process))

        elif message.get("action") == "stop":
            if download_process:
                # Stoppe den Download-Prozess
                download_process.terminate()
                download_process = None

async def send_output_to_frontend(ws, process):
    while True:
        output = await process.stdout.readline()
        if output:
            output_data = output.decode().strip()
            try:
                # Versuche, die Ausgabe als JSON zu parsen
                json_data = json.loads(output_data)
                await ws.send(json.dumps(json_data))
            except json.JSONDecodeError:
                # Senden von Rohdaten, falls das Parsen fehlschlägt
                await ws.send(json.dumps({"type": "output", "data": output_data}))
        else:
            break

if __name__ == '__main__':
    app.run()
