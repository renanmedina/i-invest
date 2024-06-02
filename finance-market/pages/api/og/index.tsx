import { ImageResponse } from '@vercel/og';
import { NextRequest } from 'next/server';

export const config = {
  runtime: 'experimental-edge',
};

export default async function handler(req: NextRequest) {
  const { searchParams } = new URL(req.url);

  // ?ticker=<ticker>&logoUrl=<logoUrl>&longName=<longName>
  const hasTicker = searchParams.has('ticker');
  const hasLogoUrl = searchParams.has('logoUrl');
  const hasLongName = searchParams.has('longName');
  const ticker = hasTicker ? searchParams.get('ticker')?.slice(0, 100) : '';
  const logoUrl = hasLogoUrl ? searchParams.get('logoUrl') : '';
  const longName = hasLongName ? searchParams.get('longName') : '';

  return new ImageResponse(
    (
      <div
        tw="flex w-full h-full bg-gray-900 p-6 flex-col"
        style={{
          backgroundImage:
            'radial-gradient(circle at 25px 25px, #344261 2%, transparent 0%), radial-gradient(circle at 75px 75px, #344261 2%, transparent 0%)',
          backgroundSize: '100px 100px',
        }}
      >
        <div tw="flex flex-col">
          <div tw="flex justify-between items-center">
            <div tw="flex items-center">
              <Favicon />
              <h1 tw="text-10 font-bold text-white">brapi.dev</h1>
            </div>

            <div tw="flex">
              <h2 tw="text-8 text-white">
                Preço, variação, gráfico, notícias e mais
              </h2>
            </div>
          </div>

          <div tw="flex items-center justify-center h-full -mt-14 flex-col">
            <img tw="w-64 h-64 rounded-2xl" src={logoUrl} />
            <p tw="text-12 font-semibold text-white">
              {ticker} - {longName}
            </p>
          </div>
        </div>
      </div>
    ),
    {
      height: 630,
      width: 1200,
    },
  );
}

const Favicon = () => {
  return (
    <svg
      width="48"
      height="48"
      viewBox="0 0 512 512"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        d="M34 292.571V512H269.493C402.107 512 477.984 385.196 477.984 303.508V0H382.966L318.711 72.4592C109.535 72.4592 34 188.325 34 292.571Z"
        fill="#7347AB"
      ></path>
      <path
        d="M252.683 294.346C171.313 226.417 49.3805 213.617 49.3805 213.617C39.9185 237.441 36.668 252.825 34 283.343V512H275.303C305.025 510.098 320.914 506.621 347.762 495.935C347.762 495.935 334.053 362.278 252.683 294.346Z"
        fill="url(#paint0_linear_324_15)"
        fill-opacity="0.3"
      ></path>
      <path
        d="M179.602 290.862H112.612V435.781H179.602V290.862Z"
        fill="white"
      ></path>
      <path
        d="M282.823 167.818H215.832V435.781H282.823V167.818Z"
        fill="white"
      ></path>
      <path
        d="M386.042 0.341675H319.052V435.781H386.042V0.341675Z"
        fill="white"
      ></path>
      <defs>
        <linearGradient
          id="SvgjsLinearGradient1000"
          x1="190.881"
          y1="213.617"
          x2="190.881"
          y2="512"
          gradientUnits="userSpaceOnUse"
        >
          <stop stop-color="#7347AB"></stop>
          <stop offset="1" stop-color="white" stop-opacity="0.28"></stop>
        </linearGradient>
      </defs>
    </svg>
  );
};
