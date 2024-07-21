from anticaptchaofficial.recaptchav2proxyless import *
import os

CAPTCHA_SITE_KEY = os.getenv('B3_CAPTCHA_SITE_KEY')
ANTI_CAPTCHA_API_KEY = os.getenv('ANTI_CAPTCHA_API_KEY')

class AntiCaptchaService:
  def __init__(self, website, apiKey=ANTI_CAPTCHA_API_KEY, captchaSiteKey=CAPTCHA_SITE_KEY):
    self._solver = recaptchaV2Proxyless()
    self._solver.set_verbose(1)
    self._solver.set_key(apiKey)
    self._solver.set_website_url(website)
    self._solver.set_website_key(captchaSiteKey)

  def solve(self):
    g_response = self._solver.solve_and_return_solution()
    if g_response != 0:
      return g_response
    else:
      print("task finished with error "+self._solver.error_code)
      return None