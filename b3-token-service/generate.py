from seleniumwire import webdriver
from selenium.webdriver import ChromeOptions
from automations import token_extraction

options = ChromeOptions()
# options.add_argument("--window-size=100,100")
# options.add_argument("--headless")
driver = webdriver.Chrome(options=options)
task = token_extraction.ExtractionTask(driver)
tokenData = task.run()
print("Token extracted:")
print(tokenData)

