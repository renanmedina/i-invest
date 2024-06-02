import clsx from 'clsx';
import NextLink from 'next/link';
import { IQuoteList } from '~/services/getQuoteList';
import { numberToMoney } from '~/utils/formatNumbers';

export const QuoteSimpleCard = ({
  stock,
  name,
  change,
  close,
  logo,
}: IQuoteList) => {
  return (
    <NextLink
      className="flex flex-col justify-center px-4 py-4 bg-base-300 rounded-md flex-grow min-w-[192px]"
      role="button"
      href={`/quote/${stock}`}
    >
      <div className="flex justify-between">
        <div className="text-xl font-bold">{stock}</div>
        <img
          src={logo}
          alt={stock}
          className="w-8 h-8 rounded-lg"
          width={32}
          height={32}
          loading="lazy"
          // dont send cookies to other domains
          crossOrigin="anonymous"
        />
      </div>
      <div className="text-sm capitalize">{name?.toLowerCase()}</div>
      <div className="flex justify-between">
        <div className="text-sm">{numberToMoney(close)}</div>
        <div
          className={clsx({
            'text-sm': true,
            'text-success': change > 0,
            'text-error': change < 0,
          })}
        >
          {change > 0 ? '▲' : '▼'}
          {(Math.floor(change * 100) / 100).toFixed(2)}%
        </div>
      </div>
    </NextLink>
  );
};
