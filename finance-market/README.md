<p align="center">
  <a href="https://brapi.dev/#gh-dark-mode-only">
    <img src="./public/logotype.svg" width="256" />
  </a>
  <a href="https://brapi.dev/#gh-light-mode-only">
    <img src="./public/favicon.svg" width="96" />
  </a>
</p>

<p align="center">
    <a href="https://github.com/Alissonsleal/brapi/issues"><img alt="GitHub issues" src="https://img.shields.io/github/issues/Alissonsleal/brapi?color=blueviolet&style=flat-square"></a>
    <a href="https://github.com/Alissonsleal/brapi/network"><img alt="GitHub forks" src="https://img.shields.io/github/forks/Alissonsleal/brapi?color=blueviolet&style=flat-square"></a>
    <a href="https://github.com/Alissonsleal/brapi/stargazers"><img alt="GitHub stars" src="https://img.shields.io/github/stars/Alissonsleal/brapi?color=blueviolet&style=flat-square"></a>
    <a href="https://github.com/Alissonsleal/brapi/blob/master/LICENSE"><img alt="GitHub license" src="https://img.shields.io/github/license/Alissonsleal/brapi?color=blueviolet&style=flat-square"></a>
    <a href="https://github.com/Alissonsleal/"><img alt="Alisson Leal GitHub Profile" src="https://img.shields.io/badge/made%20by-Alisson%20Leal-blueviolet?style=flat-square&logo=appveyor"></a>
</p>

# brapi

Na brapi, você tem acesso à cotação em tempo real das ações da Bovespa e criptomoedas com um delay de até 15 minutos. Você tem acesso à uma API que mostra todos os dados necessarios para você desenvolver a sua própria aplicação relacionada ao mercado de ações brasileiro, cripto ou moedas. Ajudamos desenvolvedores a construir o futuro das fintechs democratizando o acesso aos dados do mercado financeiro brasileiro e de criptomoedas.

- Saiba mais: [https://brapi.dev](https://brapi.dev)

- Documentação: [https://brapi.dev/docs](https://brapi.dev/docs)

- Collection do Postman: [https://app.getpostman.com/run-collection/da5f72c67bf46c6c4a5f](https://app.getpostman.com/run-collection/da5f72c67bf46c6c4a5f)

- SwaggerAPI: [https://app.swaggerhub.com/apis-docs/Alissonsleal/brapi](https://app.swaggerhub.com/apis-docs/Alissonsleal/brapi)

## Recursos

### Acesso em tempo real

Providenciamos dados do mercado de ação brasileiro em tempo real e totalmente grátis.

### Requisições Ilimitadas

Nossa API não tem limites por enquanto, use a vontade.

### Suportamos Criptomoedas

Você pode buscar infomações de qualquer criptomoeda em qualquer moeda

```json
// https://brapi.dev/api/v2/crypto?coin=BTC&currency=BRL
{
  "coins": [
    {
      "currency": "BRL",
      "currencyRateFromUSD": 5.2429,
      "coinName": "Bitcoin USD",
      "coin": "BTC",
      "regularMarketChange": -1323.7298561629998,
      "regularMarketPrice": 245026.23419429996,
      "regularMarketChangePercent": -0.5373331,
      "regularMarketDayLow": 239498.84395450001,
      "regularMarketDayHigh": 248097.83958829998,
      "regularMarketDayRange": "239498.84395450001 - 248097.83958829998",
      "regularMarketVolume": 157595524783.9232,
      "marketCap": 4603348733480.141,
      "regularMarketTime": 1629063662,
      "coinImageUrl": "https://s.yimg.com/uc/fin/img/reports-thumbnails/1.png"
    }
  ]
}
```

### Suportamos Moedas

Você pode buscar infomações e converter várias moedas

```json
// https://brapi.dev/api/v2/currency?currency=USD-BRL
{
  "currency": [
    {
      "fromCurrency": "USD",
      "toCurrency": "BRL",
      "name": "Dólar Americano/Real Brasileiro",
      "high": "5.3469",
      "low": "5.2517",
      "bidVariation": "0.0355",
      "percentageChange": "0.68",
      "bidPrice": "5.2885",
      "askPrice": "5.2895",
      "updatedAtTimestamp": "1631912338",
      "updatedAtDate": "2021-09-17 17:58:58"
    }
  ]
}
```

### De Desenvolvedores para Desenvolvedores

Desenvolvemos a API mais fácil de começar a utilizar, sem cadastros, sem cartão de crédito e sem qualquer tipo de autenticação.

```json
// GET https://brapi.dev/api/quote/COGN3
{
  "results": {
    "symbol": "COGN3",
    "shortName": "COGNA ON    ON      NM",
    "longName": "Cogna Educação S.A.",
    "currency": "BRL",
    "regularMarketPrice": 4.63,
    "regularMarketDayHigh": 4.7,
    "regularMarketDayLow": 4.58,
    "regularMarketDayRange": "4.58 - 4.7",
    "regularMarketChange": 0.010000229,
    "regularMarketChangePercent": 0.21645518,
    "regularMarketTime": "2021-02-05T21:06:00.000Z",
    "marketCap": 8615874560,
    "regularMarketVolume": 29814400,
    "regularMarketPreviousClose": 4.62,
    "regularMarketOpen": 4.65,
    "averageDailyVolume10Day": 42852887,
    "averageDailyVolume3Month": 54733323,
    "fiftyTwoWeekLowChange": 1.0600002,
    "fiftyTwoWeekLowChangePercent": 0.2969188,
    "fiftyTwoWeekRange": "3.57 - 12.08",
    "fiftyTwoWeekHighChange": -7.45,
    "fiftyTwoWeekHighChangePercent": -0.61672187,
    "fiftyTwoWeekLow": 3.57,
    "fiftyTwoWeekHigh": 12.08,
    "twoHundredDayAverage": 5.499161,
    "twoHundredDayAverageChange": -0.86916065,
    "twoHundredDayAverageChangePercent": -0.15805332
  },
  "requestedAt": "2021-02-06T21:45:13.131Z"
}
```

<br />

## Feito com:

- [x] Node.js - Vercel Serverless Functions
- [x] Next.js
- [x] Typescript
- [x] Tailwindcss
- [x] Yahoo API
- [x] TradingView API
- [x] https://economia.awesomeapi.com.br
- [x] Infraestrutura na Vercel

## Contribuições

Sinta-se livre para contribuir ou reportar algum erro ou sujestão.

## Exoneração de Responsabilidade

Esse projeto não é afiliado de qualquer forma à Yahoo ou TradingView or qualquer outra empresa mencionada aqui ou no site [brapi.dev](brapi.dev).

Essa é uma API para fins informativos. Não garantimos a precisão dos dados
fornecidos pela API ou contidos nesta página, uma vez que devem
ser utilizados apenas para efeitos informativos. Trabalhamos pela
estabilidade e precisão dos dados, porém, os dados podem estar
atrasados ou errados "no estado em que se encontram", confirme
todos os dados antes de efetuar qualquer ação que possa ser
afetada por estes valores, assim como demais endpoints da API.

## Desenvolvedor

- Twitter - [@alissonsleal](https://twitter.com/alissonsleal)
- Email - [brapi@proton.me](mailto:brapi@proton.me)
- StackOverflow - [Alisson Leal](https://stackoverflow.com/users/14122260/alisson-leal)

[Subir ao Topo 🚀](#brapi)

<p align="center">
  <a href="https://brapi.dev/#gh-dark-mode-only">
    <img src="./public/logotype.svg" width="256" />
  </a>
  <a href="https://brapi.dev/#gh-light-mode-only">
    <img src="./public/favicon.svg" width="96" />
  </a>
</p>

<p align="center">
  <a
  href="https://vercel.com/?utm_source=alisson-oss&utm_campaign=oss"
  rel="noreferrer noopener"
  target="\_blank">
  <img src="https://www.datocms-assets.com/31049/1618983297-powered-by-vercel.svg" />
  </a>
</p>
