import Image from 'next/image';
import Link from 'next/link';

const companyLinks = [
  {
    href: '/about',
    label: 'Sobre',
  },
  {
    href: '/contact',
    label: 'Contato',
  },
  {
    href: '/career',
    label: 'Carreira',
  },
  {
    href: '/press-kit',
    label: 'Press kit',
  },
];

const legalLinks = [
  {
    href: '/terms-of-use',
    label: 'Termos de uso',
  },
  {
    href: '/privacy-policy',
    label: 'Política de privacidade',
  },
  {
    href: '/cookie-policy',
    label: 'Política de cookies',
  },
];

const Footer = () => {
  return (
    <footer className="text-gray-400 bg-gray-900 body-font border-t border-gray-800">
      <div className="container px-5 py-8 flex flex-wrap mx-auto items-center flex-row justify-between">
        <div className="w-full flex flex-wrap md:flex-nowrap gap-8 justify-center md:justify-start">
          <div className="footer flex space-x-8 max-w-fit">
            <div>
              <span className="mb-2 font-bold uppercase opacity-80">
                Empresa
              </span>
              {companyLinks.map((link) => (
                <Link
                  key={link.href}
                  href={link.href}
                  className="link link-hover"
                >
                  {link.label}
                </Link>
              ))}
            </div>
            <div>
              <span className="mb-2 font-bold uppercase opacity-80">Legal</span>
              {legalLinks.map((link) => (
                <Link
                  key={link.href}
                  href={link.href}
                  className="link link-hover"
                >
                  {link.label}
                </Link>
              ))}
            </div>
          </div>
          <div className="flex md:flex-nowrap flex-wrap justify-center md:justify-end w-full gap-8 md:gap-2">
            <p className="text-gray-400 text-sm md:ml-6 md:mt-0 mt-2 text-center sm:text-left max-w-md">
              Ajudamos desenvolvedores a construir o futuro das fintechs
              democratizando o acesso aos dados do mercado financeiro
              brasileiro.
            </p>
            <form
              action="https://formsubmit.co/166feb7796ce3daa45ec9848d1ae062e"
              method="POST"
              className="relative sm:w-64 w-full max-w-full min-w-[200px] sm:mr-4 mr-2"
            >
              <input
                type="hidden"
                name="_next"
                value={encodeURI(
                  'https://brapi.dev?show-toast=Agora você receberá nossas novidades!',
                )}
              />
              <input type="hidden" name="_subject" value="Newsletter brapi" />
              <input type="hidden" name="_captcha" value="false" />
              <input type="text" name="_honey" style={{ display: 'none' }} />
              <input type="hidden" name="_template" value="table" />
              <input
                type="hidden"
                name="_webhook"
                value="https://brapi.dev/api/webhook/form/newsletter"
              />
              <input type="text" name="_honey" style={{ display: 'none' }} />
              <label
                htmlFor="email"
                className="leading-7 text-sm text-gray-400"
              >
                Receba novidades por email
              </label>
              <input
                type="email"
                placeholder="email@empresa.com"
                id="email"
                name="email"
                className="w-full bg-gray-800 bg-opacity-40 rounded border border-gray-700 focus:ring-2 focus:ring-brand-900 focus:bg-transparent focus:border-brand-500 text-base outline-none text-gray-100 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out"
              />
              <button className="hidden">Cadastrar</button>
            </form>
          </div>
        </div>
      </div>

      <div className="footer py-4 border-t bg-base-200 border-base-300">
        <div className="flex px-5 items-center container flex-col md:flex-row justify-center md:justify-between text-center mx-auto">
          <div className="flex items-center grid-flow-col space-x-2">
            <Image
              src="/favicon.svg"
              className="max-h-5"
              alt="Logo brapi"
              width={20}
              height={20}
            />
            <p className="text-gray-400 text-sm text-center sm:text-left flex justify-center space-x-2">
              © {new Date().getFullYear()} brapi —
              <a
                href="mailto:brapi@proton.me"
                className="text-gray-400 ml-1 hover:text-white"
                target="_blank"
                rel="noopener noreferrer"
                tabIndex={0}
              >
                brapi@proton.me
              </a>
              <a
                href="https://github.com/alissonsleal/brapi"
                target="_blank"
                rel="noopener noreferrer"
                className="text-gray-400 hover:text-white transition-colors"
              >
                <span className="sr-only">Alissonsleal Github</span>
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="w-5 h-5"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  strokeWidth="2"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                >
                  <path d="M9 19c-5 1.5-5-2.5-7-3m14 6v-3.87a3.37 3.37 0 0 0-.94-2.61c3.14-.35 6.44-1.54 6.44-7A5.44 5.44 0 0 0 20 4.77 5.07 5.07 0 0 0 19.91 1S18.73.65 16 2.48a13.38 13.38 0 0 0-7 0C6.27.65 5.09 1 5.09 1A5.07 5.07 0 0 0 5 4.77a5.44 5.44 0 0 0-1.5 3.78c0 5.42 3.3 6.61 6.44 7A3.37 3.37 0 0 0 9 18.13V22"></path>
                </svg>{' '}
              </a>
            </p>
          </div>
          <div className="md:place-self-center md:justify-self-end">
            <div className="grid grid-flow-col gap-4">
              <a
                href="https://vercel.com/?utm_source=alisson-oss&utm_campaign=oss"
                className="sm:ml-auto sm:mt-0 mt-2 sm:w-auto w-full sm:text-left text-center text-gray-400 text-sm hover:text-white transition-colors"
                tabIndex={0}
              >
                Powered by ▲ Vercel
              </a>
            </div>
          </div>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
