export const defaultOg = {
  title: 'brapi - API de ações da bolsa de valores brasileira',
  description:
    'API ilimitada da BOVESPA e Cryptomoedas. Ajudamos desenvolvedores a construir o futuro das fintechs democratizando o acesso aos dados do mercado financeiro brasileiro.',
  type: 'website',
  siteName: 'brapi',
  url: 'https://brapi.dev',
};

export const createOg = (name: string) => {
  return {
    ...defaultOg,
    images: [
      {
        url: `https://${
          process.env.NEXT_PUBLIC_VERCEL_URL || 'brapi.dev'
        }/api/og?ticker=brapi&logoUrl=https://brapi.dev/favicon.svg&longName=${name}`,
      },
    ],
  };
};
