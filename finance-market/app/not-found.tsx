import Link from 'next/link';

const NotFoundPage = () => {
  return (
    <main className="flex flex-col md:justify-center items-center min-h-screen space-y-12">
      <div className="text-center">
        <h1 className="text-9xl font-bold">404</h1>
        <p>Essa página não existe :(</p>
      </div>
      <Link href="/" className="btn btn-primary">
        Voltar para a página inicial
      </Link>
    </main>
  );
};

export default NotFoundPage;
