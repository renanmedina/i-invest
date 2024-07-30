import json 
import os

def write(tokenData, filepath=None):
  if filepath == None:
    filepath = os.getenv('B3_TOKEN_CACHEFILE_PATH')
    
  file = open(filepath, 'w')
  file.write(json.dumps(tokenData, indent=2))
  file.close
