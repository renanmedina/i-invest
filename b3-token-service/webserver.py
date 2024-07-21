from http.server import BaseHTTPRequestHandler, HTTPServer
from seleniumwire import webdriver
from automations import token_extraction
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
      self.send_response(200, json.dumps(tokenData, indent=2))
    else:
      self.send_response(404)

    return

if __name__ == "__main__":
  webServer = HTTPServer((hostName, serverPort), WebserverHandler)
  print("Server started http://%s:%s" % (hostName, serverPort))

  try:
    webServer.serve_forever()
  except KeyboardInterrupt:
    pass

  webServer.server_close()
  print("Server stopped.")