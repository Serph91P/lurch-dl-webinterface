from quart import Quart, render_template, request, websocket
import asyncio
import subprocess

app = Quart(__name__)

async def run_lurch_dl(url):
    command = ["/usr/local/bin/lurch-dl", "--url", url]
    proc = await asyncio.create_subprocess_exec(*command, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    stdout, stderr = await proc.communicate()
    if proc.returncode == 0:
        return stdout.decode()
    else:
        return stderr.decode()

@app.route('/', methods=['GET', 'POST'])
async def index():
    if request.method == 'POST':
        form = await request.form
        url = form.get('url')
        if url:
            output = await run_lurch_dl(url)
            return await render_template('index.html', output=output, url=url)
    return await render_template('index.html', url='')

if __name__ == '__main__':
    app.run()
