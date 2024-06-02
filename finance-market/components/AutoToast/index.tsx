'use client';

import hotToast, { Toaster as HotToaster } from 'react-hot-toast';
import { useSearchParams } from 'next/navigation';
import { useEffect } from 'react';

export const AutoToast = () => {
  const searchParams = useSearchParams();

  useEffect(() => {
    const toastMessage = searchParams.get('show-toast') || '';

    if (toastMessage) {
      hotToast(toastMessage);
    }
  }, []);

  return <HotToaster />;
};
