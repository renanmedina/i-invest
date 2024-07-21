import time
from automations.steps import FillUsernameStep, FillPasswordStep, CheckCaptchaStep, SubmitLoginStep, ExtractTokenFromTokenRequest

LOGIN_URL = 'https://www.investidor.b3.com.br/login'
TOKEN_URL = 'https://b3investidor.b2clogin.com/b3Investidor.onmicrosoft.com/oauth2/v2.0/token?'
WAIT_PAGE_LOAD_SECONDS = 4

class ExtractionTask:
  def __init__(self, webdriver, ) -> None:
    self._seleniumDriver = webdriver
    self._steps = [
      { "sleep": WAIT_PAGE_LOAD_SECONDS, "run": FillUsernameStep },
      { "sleep": 5, "run": FillPasswordStep},
      { "sleep": 7, "run": CheckCaptchaStep },
      { "sleep": 20, "run": SubmitLoginStep },
      { "sleep": 10, "run": ExtractTokenFromTokenRequest },
    ]
    self._outputs = []

  def run(self, username, password):
    currentStepIndex = 0
    self._seleniumDriver.get(LOGIN_URL)

    while (currentStepIndex < len(self._steps)):
      currentStep = self._steps[currentStepIndex]
      stepName = currentStep["run"].__name__
      print(f'Waiting {currentStep["sleep"]}s to run step {stepName}')
      time.sleep(currentStep["sleep"])
      print(f'Running step {stepName} code')
      output = currentStep["run"](self._seleniumDriver, context={"user_cpf": username, "user_password": password})
      print(f'Step {stepName} executed')
      # self._outputs[currentStepIndex] = output
      currentStepIndex += 1
    return output