import Main from '~/components/Main';

export default function IndexPage() {
  /* @ts-expect-error Server Component */
  return <Main />;
}
