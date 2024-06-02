import { QuoteSimpleCard } from '~/components/MainQuotes/QuoteSimpleCard';
import { getQuoteList } from '~/services/getQuoteList';

export const SideScroller = async () => {
  let data = await getQuoteList();

  return (
    <div className="flex space-x-2 overflow-hidden pb-5 hover:overflow-auto hover:pb-[5px] container mx-auto px-5 md:pt-6 bg-gray-900">
      {data.map((quote) => (
        <QuoteSimpleCard
          key={quote.stock}
          stock={quote.stock}
          name={quote.name}
          close={quote.close}
          change={quote.change}
          logo={quote.logo}
        />
      ))}
    </div>
  );
};
