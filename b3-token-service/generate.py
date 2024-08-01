from seleniumwire import webdriver
from automations import token_extraction, token_cacher
from dotenv import load_dotenv
import os

load_dotenv()

driver = webdriver.Chrome()
task = token_extraction.ExtractionTask(driver)
tokenData = task.run(os.getenv('USER_CPF'), os.getenv('USER_PASSWORD'))
token_cacher.write(tokenData, filepath=os.getenv('B3_TOKEN_CACHEFILE_PATH'))
