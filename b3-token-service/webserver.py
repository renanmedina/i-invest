from http.server import BaseHTTPRequestHandler, HTTPServer
from seleniumwire import webdriver
from automations import token_extraction, token_cacher
import json
import os
from dotenv import load_dotenv

load_dotenv()

hostName = "0.0.0.0"
serverPort = int(os.getenv('SERVER_PORT'))

class WebserverHandler(BaseHTTPRequestHandler):
  def do_GET(self):
    if self.path == "/new-token":
      driver = webdriver.Chrome()
      task = token_extraction.ExtractionTask(driver)
      tokenData = task.run(os.getenv('USER_CPF'), os.getenv('USER_PASSWORD'))
      token_cacher.write(tokenData, filepath=os.getenv('B3_TOKEN_CACHEFILE_PATH'))
      self.send_json(200, tokenData)
    else:
      self.send_json(404, {"error": "Path not found"})
    return
  
  def send_json(self, status, jsonData):
    self.send_response(status)
    self.send_header('Content-Type', 'application/json')
    self.end_headers()
    json_string = json.dumps(jsonData, indent=2, ensure_ascii=False)
    self.wfile.write(bytes(json_string, 'utf-8'))

if __name__ == "__main__":
  webServer = HTTPServer((hostName, serverPort), WebserverHandler)
  print("Server started http://%s:%s" % (hostName, serverPort))

  try:
    webServer.serve_forever()
  except KeyboardInterrupt:
    pass

  webServer.server_close()
  print("Server stopped.")