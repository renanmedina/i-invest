import Link from 'next/link';

export const WhatWeDo = () => {
  return (
    <section className="text-gray-400 body-font bg-gray-900">
      <div className="container flex flex-wrap px-5 py-24 mx-auto items-center">
        <div className="md:w-1/2 md:pr-12 md:py-8 md:border-r md:border-b-0 md:mb-0 mb-10 pb-10 border-b border-gray-800">
          <h1 className="sm:text-3xl text-2xl font-medium title-font mb-2 text-white">
            O que realmente fazemos
          </h1>
          <p className="leading-relaxed text-base">
            A brapi é uma API de ações, moedas e criptomoedas onde você tem
            acesso aos dados de qualquer ação, como preço atual, preço mínimo e
            máximo do dia ou ano, variação, volume, volume transacionado,
            histórico de preços, etc. Para moedas e criptomoedas você ainda pode
            converter de uma moeda para qualquer outra.
          </p>
          <Link
            prefetch={false}
            href="/docs"
            className="text-brand-400 inline-flex items-center mt-4"
            tabIndex={0}
          >
            Começar
            <svg
              fill="none"
              stroke="currentColor"
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth="2"
              className="w-4 h-4 ml-2"
              viewBox="0 0 24 24"
            >
              <path d="M5 12h14M12 5l7 7-7 7"></path>
            </svg>
          </Link>
        </div>
        <div className="flex flex-col md:w-1/2 md:pl-12">
          <h2 className="title-font font-semibold text-white tracking-wider text-sm mb-3">
            CATEGORIAS
          </h2>
          <ul className="flex flex-wrap list-none -mb-1">
            <li className="lg:w-1/3 mb-1 w-1/2">
              <Link
                prefetch={false}
                href="/docs#stocks"
                className="hover:text-white"
                tabIndex={0}
              >
                Bovespa
              </Link>
            </li>
            <li className="lg:w-1/3 mb-1 w-1/2">
              <Link
                prefetch={false}
                href="/docs#stocks-history"
                className="hover:text-white"
                tabIndex={0}
              >
                Dados históricos
              </Link>
            </li>
            <li className="lg:w-1/3 mb-1 w-1/2">
              <Link
                prefetch={false}
                href="/docs#stocks-fundamentalist"
                className="hover:text-white"
                tabIndex={0}
              >
                Dados fundamentalistas
              </Link>
            </li>
            <li className="lg:w-1/3 mb-1 w-1/2">
              <Link
                prefetch={false}
                href="/docs#crypto"
                className="hover:text-white"
                tabIndex={0}
              >
                Criptomoedas
              </Link>
            </li>
            <li className="lg:w-1/3 mb-1 w-1/2">
              <Link
                prefetch={false}
                href="/docs#currency"
                className="hover:text-white"
                tabIndex={0}
              >
                Moedas
              </Link>
            </li>
            <li className="lg:w-1/3 mb-1 w-1/2">
              <Link
                prefetch={false}
                href="/docs#search"
                className="hover:text-white"
                tabIndex={0}
              >
                Pesquisa rápida
              </Link>
            </li>
          </ul>
        </div>
      </div>
    </section>
  );
};
