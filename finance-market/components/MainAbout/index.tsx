import Link from 'next/link';

const MainAbout = () => {
  return (
    <main>
      <section className="text-gray-400 bg-gray-900 body-font">
        <div className="container px-5 py-0 md:py-24 mx-auto flex flex-col">
          <div className="lg:w-4/6 mx-auto">
            <div className="flex flex-col sm:flex-row mt-10">
              <div className="sm:w-1/3 text-center sm:pr-8 sm:py-8">
                <div className="avatar">
                  <div className="w-24 mask mask-hexagon">
                    <img
                      src="https://github.com/alissonsleal.png"
                      alt="avatar"
                    />
                  </div>
                </div>
                <div className="flex flex-col items-center text-center justify-center">
                  <h2 className="font-medium title-font mt-4 text-white text-lg">
                    Alisson Leal
                  </h2>
                  <div className="w-12 h-1 bg-brand-500 rounded mt-2 mb-4"></div>
                  <p className="text-base text-gray-400">
                    Founder e Software Engineer na brapi
                  </p>
                </div>
              </div>
              <div className="sm:w-2/3 sm:pl-8 sm:py-8 sm:border-l border-gray-800 sm:border-t-0 border-t mt-4 pt-4 sm:mt-0 text-center sm:text-left">
                <p className="leading-relaxed text-lg mb-4">
                  Eu desenvolvi a brapi porque um dia eu precisei usar alguns
                  dados do mercado de ações da Bovespa e foi nesse momento que
                  <b>
                    {' '}
                    eu descobri que simplesmente não existe uma API fácil e
                    gratuita para buscar esses dados,{' '}
                  </b>
                  na verdade eu até encontrei uma outra API pra isso mas era bem
                  burocrática, precisava criar uma conta com os meus dados
                  pessoais, o meu cartão de crédito e ainda precisaria
                  configurar uma chave de autenticação pra começar a fazer
                  qualquer requisição, depois de tudo isso eu ainda teria um
                  limite de requisições super baixo.
                </p>
                <p className="leading-relaxed text-lg mb-4">
                  Na brapi, nós acreditamos que os desenvolvedores devem ter as
                  ferramentas corretas e fácil acesso para que tecnologias
                  disruptivas possam ser criadas. Estamos felizes em poder
                  simplificar gratuitamente o processo de criação de novas
                  tecnologias.
                </p>
                <p className="leading-relaxed text-lg mb-4">
                  Conseguimos deixar a API gratuita e acessível para todas as
                  pessoas através da nossa principal patrocinadora, a
                  <a
                    href="https://vercel.com/?utm_source=alisson-oss&utm_campaign=oss"
                    target="_blank"
                    rel="noopener noreferrer"
                  >
                    {' '}
                    Vercel,{' '}
                  </a>
                  onde é disponibiilizada toda a sua infraestrutura global.
                </p>
                <Link
                  prefetch={false}
                  href="/docs"
                  className="text-brand-400 inline-flex items-center"
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
            </div>
          </div>
        </div>
      </section>
    </main>
  );
};

export default MainAbout;
