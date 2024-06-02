import Link from 'next/link';

export const CTA = () => {
  return (
    <section className="text-gray-400 bg-gray-900 body-font">
      <div className="container px-5 py-24 mx-auto">
        <div className="lg:w-2/3 flex flex-col sm:flex-row sm:items-center items-start mx-auto">
          <h1 className="flex-grow sm:pr-16 text-2xl font-medium title-font text-white">
            Pronto para desenvolver uma aplicação?
          </h1>
          <Link prefetch={false} href="/docs" className="btn btn-primary">
            Começar
          </Link>
        </div>
      </div>
    </section>
  );
};
