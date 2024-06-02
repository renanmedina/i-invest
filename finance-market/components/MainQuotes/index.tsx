import { MainQuote } from '~/components/MainQuotes/MainQuote';
import { QuoteChart } from '~/components/MainQuotes/QuoteChart';
import { RecommendedQuotes } from '~/components/MainQuotes/RecommendedQuotes';
import { getCurrentQuote } from '~/services/getCurrentQuote';

interface IMainQuotesProps {
  currentStock: string;
}

const MainQuotes = async ({ currentStock }: IMainQuotesProps) => {
  const [currentQuote] = await getCurrentQuote({
    stocks: currentStock,
    interval: '1mo',
    range: 'max',
  });

  if (!currentQuote) {
    return <div>Não foi possível carregar os dados de {currentStock}</div>;
  }

  return (
    <div className="py-4 flex w-full flex-col space-y-4">
      {/* @ts-expect-error Server Component */}
      <MainQuote quote={currentQuote} />

      <div className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-8">
        <div className="h-[400px] md:col-span-2 lg:col-span-3">
          <QuoteChart
            historicalDataPrices={currentQuote?.historicalDataPrice || []}
            source={`https://brapi.dev/api/quote/${currentQuote.symbol}?range=max&interval=1d&fundamental=true`}
          />
        </div>

        {/* @ts-expect-error Server Component */}
        <RecommendedQuotes quote={currentQuote.symbol} />
      </div>
    </div>
  );
};

export default MainQuotes;
