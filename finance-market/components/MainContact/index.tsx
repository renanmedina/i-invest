export const MainContact = () => {
  return (
    <section className="text-gray-400 bg-gray-900 body-font relative">
      <div className="container px-5 py-24 mx-auto">
        <div className="flex flex-col text-center w-full mb-12">
          <h1 className="sm:text-3xl text-2xl font-medium title-font mb-4 text-white">
            Entre em Contato
          </h1>
          <p className="lg:w-2/3 mx-auto leading-relaxed text-base">
            Encontrou algum problema? Tem alguma dúvida ou feedback?
          </p>
        </div>
        <div className="lg:w-1/2 md:w-2/3 mx-auto">
          <div className="flex flex-wrap -m-2">
            <form
              className="flex flex-wrap -m-2"
              action="https://formsubmit.co/166feb7796ce3daa45ec9848d1ae062e"
              method="POST"
            >
              <input
                type="hidden"
                name="_next"
                value={encodeURI(
                  'https://brapi.dev/contact?show-toast=Recebemos o seu email',
                )}
              />
              <input
                type="hidden"
                name="_autoresponse"
                value="Recebemos seu email!"
              />
              <input type="hidden" name="_subject" value="Contato brapi" />
              <input type="hidden" name="_captcha" value="false" />
              <input type="text" name="_honey" style={{ display: 'none' }} />
              <input type="hidden" name="_template" value="table" />

              <input
                type="hidden"
                name="_webhook"
                value="https://brapi.dev/api/webhook/form/contact"
              />

              <input type="text" name="_honey" style={{ display: 'none' }} />

              <div className="p-2 w-1/2">
                <div className="relative">
                  <label
                    htmlFor="name"
                    className="leading-7 text-sm text-gray-400"
                  >
                    Nome
                  </label>
                  <input
                    type="text"
                    id="name"
                    name="name"
                    placeholder="Bill Gates - Microsoft"
                    className="w-full bg-gray-800 bg-opacity-40 rounded border border-gray-700 focus:border-primary-500 focus:bg-gray-900 focus:ring-2 focus:ring-primary-900 text-base outline-none text-gray-100 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out"
                  />
                </div>
              </div>
              <div className="p-2 w-1/2">
                <div className="relative">
                  <label
                    htmlFor="email"
                    className="leading-7 text-sm text-gray-400"
                  >
                    Email
                  </label>
                  <input
                    type="email"
                    id="email"
                    name="email"
                    placeholder="bill@microsoft.com"
                    className="w-full bg-gray-800 bg-opacity-40 rounded border border-gray-700 focus:border-primary-500 focus:bg-gray-900 focus:ring-2 focus:ring-primary-900 text-base outline-none text-gray-100 py-1 px-3 leading-8 transition-colors duration-200 ease-in-out"
                  />
                </div>
              </div>
              <div className="p-2 w-full">
                <div className="relative">
                  <label
                    htmlFor="message"
                    className="leading-7 text-sm text-gray-400"
                  >
                    Mensagem
                  </label>
                  <textarea
                    id="message"
                    name="message"
                    placeholder="Olá, gostaria de saber mais sobre a brapi."
                    className="w-full bg-gray-800 bg-opacity-40 rounded border border-gray-700 focus:border-primary-500 focus:bg-gray-900 focus:ring-2 focus:ring-primary-900 h-32 text-base outline-none text-gray-100 py-1 px-3 resize-none leading-6 transition-colors duration-200 ease-in-out"
                  ></textarea>
                </div>
              </div>
              <div className="p-2 w-full flex justify-center">
                <button className="btn btn-accent">Enviar</button>
              </div>
            </form>
            <div className="p-2 w-full pt-8 mt-8 border-t border-gray-800 flex flex-col items-center">
              <div className="w-fit flex flex-col">
                <span className="leading-normal my-5 hover:text-white">
                  <a
                    href="mailto:brapi@proton.me"
                    className="text-primary-400 flex space-x-2 items-center"
                  >
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
                      <path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"></path>
                      <polyline points="22,6 12,13 2,6"></polyline>
                    </svg>
                    <span>brapi@proton.me</span>
                  </a>
                </span>

                <span>
                  <a
                    href="https://github.com/alissonsleal/brapi"
                    target="_blank"
                    rel="noopener noreferrer"
                    className="hover:text-white transition-colors flex space-x-2 items-center"
                  >
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
                    </svg>
                    <span>github</span>
                  </a>
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};
