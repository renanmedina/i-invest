import { Metadata } from 'next';
import { createOg } from '~/utils/og';
import { MainContact } from '../../components/MainContact';

export const metadata: Metadata = {
  title: 'Contato',
  openGraph: {
    ...createOg('Contato'),
  },
};

const ContactPage = () => {
  return (
    <div>
      <MainContact />
    </div>
  );
};

export default ContactPage;
