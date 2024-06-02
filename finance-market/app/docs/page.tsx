import { Metadata } from 'next';
import { MainDocs } from '~/components/MainDocs';
import { createOg } from '~/utils/og';
import '~/styles/swagger-ui.css';

export const metadata: Metadata = {
  title: 'Documentação',
  openGraph: {
    ...createOg('Documentação'),
  },
};

const Docs = () => {
  return (
    <>
      <div>
        <section className="text-gray-400 bg-gray-900 body-font max-w-screen min-h-screen">
          <div className="container mx-auto flex flex-col px-5 py-0 md:py-24 items-center">
            <MainDocs />
          </div>
        </section>
      </div>
    </>
  );
};

export default Docs;
