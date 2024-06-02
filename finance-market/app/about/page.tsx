import { Metadata } from 'next';
import React from 'react';
import { createOg } from '~/utils/og';
import MainAbout from '../../components/MainAbout';

export const metadata: Metadata = {
  title: 'Sobre',
  openGraph: {
    ...createOg('Sobre'),
  },
};

export default function AboutPage() {
  return (
    <div>
      <MainAbout />
    </div>
  );
}
