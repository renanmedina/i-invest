import { SearchResult } from '~/components/SearchInput';
import { getCurrentQuote } from '~/services/getCurrentQuote';

interface IRecommendedQuotesProps {
  quote: string;
}

const getRecommendedQuotes = async (quote: string) => {
  try {
    const res = await fetch(
      `https://query2.finance.yahoo.com/v6/finance/recommendationsbysymbol/${quote}.SA`,
      { cache: 'no-cache' },
    );
    const data = await res?.json();

    const recommendedSymbols = data?.finance?.result?.[0]?.recommendedSymbols?.map(
      (recommendedSymbol) => recommendedSymbol?.symbol?.replace('.SA', ''),
    ) as string[];

    const recommendedQuotes = await getCurrentQuote({
      stocks: recommendedSymbols,
      fundamental: false,
    });

    return recommendedQuotes;
  } catch (err) {}
};

export const RecommendedQuotes = async ({ quote }: IRecommendedQuotesProps) => {
  const recommendedQuotes = await getRecommendedQuotes(quote);

  if (!recommendedQuotes?.length) {
    return <div>Não foi possível carregar ações relacionadas a {quote}</div>;
  }

  return (
    <div className="flex flex-col space-y-2">
      <h2 className="text-xl font-bold">Ações relacionadas</h2>

      <div className="grid grid-cols-1">
        {recommendedQuotes.map((recommendedQuote) => (
          <div key={recommendedQuote.symbol}>
            <SearchResult
              name={recommendedQuote.longName}
              stock={recommendedQuote.symbol}
              change={recommendedQuote.regularMarketChange}
              close={recommendedQuote.regularMarketPrice}
              logo={recommendedQuote.logourl}
            />
          </div>
        ))}
      </div>
    </div>
  );
};
