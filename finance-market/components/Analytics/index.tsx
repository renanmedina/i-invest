'use client';

import { useEffect } from 'react';
import { Analytics as VercelAnalytics } from '@vercel/analytics/react';
import * as gtag from '~/utils/gtag';
import { usePathname } from 'next/navigation';

export const Analytics = () => {
  const pathname = usePathname();

  useEffect(() => {
    const url = new URL(pathname, window.location.origin);
    gtag.pageview(url);
  }, [pathname]);

  return <VercelAnalytics />;
};
