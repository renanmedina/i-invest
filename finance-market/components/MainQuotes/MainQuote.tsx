import Image from 'next/image';
import { IHistoricalDataPrice } from '~/@types/IHistoricalDataPrice';
import { QuoteProps } from '~/@types/QuoteProps';
import { numberToMoney, numberToSIMoney } from '~/utils/formatNumbers';

interface IMainQuoteProps {
  quote: QuoteProps & {
    historicalDataPrice: IHistoricalDataPrice[];
  };
}

export const MainQuote = async ({ quote }: IMainQuoteProps) => {
  return (
    <div className="flex space-y-6 max-w-full flex-col">
      <div className="flex justify-between md:flex-row flex-col md:space-y-0 space-y-2">
        <div className="flex space-x-2 items-center">
          <img
            src={quote.logourl}
            className="w-12 h-12 rounded-md"
            width={48}
            height={48}
          />
          <div className="flex flex-col">
            <span className="text-2xl font-bold">{quote.symbol}</span>
            <span className="text-sm">{quote.longName}</span>
          </div>
        </div>

        <div className="flex flex-col md:items-end">
          <span className="text-sm w-fit">Última atualização</span>
          <span className="text-sm font-bold w-fit">
            {new Date(quote.regularMarketTime).toLocaleString('pt-BR')}
          </span>
        </div>
      </div>

      <div className="flex justify-between lg:flex-row flex-col md:space-y-0 space-y-2">
        <div className="flex flex-col">
          <span className="text-sm">Preço</span>
          <span className="text-2xl font-bold">
            {numberToMoney(quote.regularMarketPrice)}
          </span>
        </div>

        <div className="flex flex-col">
          <span className="text-sm">Variação (dia)</span>
          <span className="text-2xl font-bold">
            {numberToMoney(quote.regularMarketChange)} (
            {quote.regularMarketChangePercent.toFixed(2)}%)
            {quote.regularMarketChangePercent > 0 ? '▲' : '▼'}
          </span>
        </div>

        <div className="flex flex-col">
          <span className="text-sm">Min. 52 Semanas</span>
          <span className="text-2xl font-bold">
            {numberToMoney(quote.fiftyTwoWeekLow)}
          </span>
        </div>

        <div className="flex flex-col">
          <span className="text-sm">Máx. 52 Semanas</span>
          <span className="text-2xl font-bold">
            {numberToMoney(quote.fiftyTwoWeekHigh)}
          </span>
        </div>

        <div className="flex flex-col">
          <span className="text-sm">Capitalização de mercado</span>
          <span
            className="text-2xl font-bold"
            title={numberToMoney(quote.marketCap)}
          >
            {numberToSIMoney(quote.marketCap)}
          </span>
        </div>
      </div>
    </div>
  );
};
