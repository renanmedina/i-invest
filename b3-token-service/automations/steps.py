from selenium.webdriver.common.by import By
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from seleniumwire.utils import decode
import json

def wait(driver, time):
  return WebDriverWait(driver, time)

def FillUsernameStep(driver, context={}):
  cpfInput = wait(driver, 10).until(EC.element_to_be_clickable((By.ID, 'investidor')))
  cpfInput.send_keys(context['user_cpf']) # fake cpf
  submitButton = wait(driver, 10).until(EC.element_to_be_clickable((By.CLASS_NAME, 'b3i-signin__login__action__entrar')))
  submitButton.click()

def FillPasswordStep(driver, context={}):
  passwordInput = driver.find_element(By.ID, 'PASS_INPUT')
  passwordInput.send_keys(context['user_password'])

def CheckCaptchaStep(driver, context={}):
  # print(driver.current_url)
  # anticaptcha = AntiCaptchaService(driver.current_url)
  # captcha_response = anticaptcha.solve()

  # driver.execute_script('var el=document.getElementById("g-recaptcha-response");el.style.display="";')
  # driver.execute_script(f'document.getElementById("g-recaptcha-response").innerHTML = "{captcha_response}";')
  # driver.execute_script('var el=document.getElementById("g-recaptcha-response");el.style.display="none";')
  captcha_frame = driver.find_element(By.XPATH, "//iframe[starts-with(@name, 'a-') and starts-with(@src, 'https://www.google.com/recaptcha')]")
  driver.switch_to.frame(captcha_frame)
  recaptcha = driver.find_element(By.ID, "recaptcha-anchor-label")
  recaptcha.click()
  driver.switch_to.parent_frame()

def SubmitLoginStep(driver, context={}):
  submitButton = driver.find_element(By.ID, 'Btn_CONTINUE')
  submitButton.click()

def ExtractTokenFromTokenRequest(driver, context={}):
  requests = driver.requests
  for request in requests:
    if request.url == "https://b3investidor.b2clogin.com/b3Investidor.onmicrosoft.com/oauth2/v2.0/token?p=B2C_1A_SIGN_IN":
      response = request.response
      decoded_body = decode(response.body, response.headers.get('Content-Encoding', 'identity'))
      token_info_string = decoded_body.decode('utf8').replace("'", '"')
      token_info = json.loads(token_info_string)
      return token_info
      # break
