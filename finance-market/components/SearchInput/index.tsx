'use client';

import clsx from 'clsx';
import Link from 'next/link';
import { useEffect, useReducer } from 'react';
import { useDebounce } from '~/hooks/useDebounce';
import { getQuoteList, IQuoteList } from '~/services/getQuoteList';
import { numberToMoney } from '~/utils/formatNumbers';

interface IInputStatus {
  isFocused: boolean;
  isLoading: boolean;
}

interface IResults {
  search: string;
  results: IQuoteList[] | string;
  cachedResults: { search: string; results: IQuoteList[] | string }[];
}

interface ISearchResults {
  results: IResults['results'];
  isLoading: IInputStatus['isLoading'];
}

type ISearchResult =
  | (IQuoteList & {
      hasBorder?: boolean;
    })
  | { isLoading: boolean };

export const SearchInput = () => {
  const [results, updateResults] = useReducer(
    (state: IResults, newState: Partial<IResults>) => ({
      ...state,
      ...newState,
    }),
    {
      search: '',
      results: [],
      cachedResults: [],
    },
  );

  const [inputStatus, updateInputStatus] = useReducer(
    (state: IInputStatus, newState: Partial<IInputStatus>) => ({
      ...state,
      ...newState,
    }),
    {
      isFocused: false,
      isLoading: false,
    },
  );
  const debouncedSearchTerm = useDebounce(results.search, 500);

  useEffect(() => {
    if (debouncedSearchTerm) {
      const cached = results.cachedResults.find(
        (result) => result.search === debouncedSearchTerm,
      );

      if (cached) {
        updateResults({ results: cached.results });
        updateInputStatus({ isLoading: false });
        return;
      }

      const getData = async () => {
        updateInputStatus({ isLoading: true });

        const listResults = await getQuoteList({
          search: debouncedSearchTerm,
          limit: 5,
        });

        updateResults({
          results: listResults?.length ? listResults : 'no_results',
          cachedResults: [
            ...results.cachedResults,
            {
              search: debouncedSearchTerm,
              results: listResults?.length ? listResults : 'no_results',
            },
          ],
        });
        updateInputStatus({ isLoading: false });
      };

      getData();
    } else {
      updateResults({ results: [] });
    }
  }, [debouncedSearchTerm]);

  return (
    <div className="items-center relative hidden md:flex">
      <input
        type="text"
        placeholder="Pesquisar Ação"
        className="input input-bordered w-64"
        value={results.search}
        onChange={(e) => {
          updateResults({ search: e.target.value });

          const hasCachedResults = results.cachedResults.find(
            (result) => result.search === e.target.value,
          );

          if (hasCachedResults) {
            updateResults({ results: hasCachedResults.results });
          }
        }}
        onFocus={() => updateInputStatus({ isFocused: true })}
        onBlur={() =>
          setTimeout(() => updateInputStatus({ isFocused: false }), 150)
        }
      />
      <svg
        xmlns="http://www.w3.org/2000/svg"
        className="h-6 w-6 absolute right-4"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          strokeLinecap="round"
          strokeLinejoin="round"
          strokeWidth="2"
          d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
        />
      </svg>

      {inputStatus.isFocused && (
        <SearchResults
          results={results.results}
          isLoading={inputStatus.isLoading}
        />
      )}
    </div>
  );
};

export const SearchResult = (props: ISearchResult) => {
  if ('isLoading' in props) {
    return (
      <div className="flex items-center p-2 border-b bg-gray-800 animate-pulse">
        <div className="w-10 h-10 rounded-full bg-gray-700" />
        <div className="pl-4 flex-grow space-y-2">
          <div className="h-7 bg-gray-700 rounded" />
          <div className="h-5 bg-gray-700 rounded" />
        </div>
      </div>
    );
  }

  const { stock, name, close, change, logo } = props;

  return (
    <Link
      className={clsx({
        'flex items-center p-2 border-gray-200 hover:bg-gray-700': true,
        'border-b': props.hasBorder,
        'rounded-md': !props.hasBorder,
      })}
      href={`/quote/${stock}`}
      title={name}
    >
      <img
        src={logo}
        className="w-10 h-10 rounded-full"
        alt={name}
        width={40}
        height={40}
      />
      <div className="ml-4">
        <p className="text-lg font-medium md:w-52 lg:w-64 md:truncate">
          {name}
        </p>
        <p className="text-sm">
          <span>{stock} </span>
          <span
            className={clsx({
              'text-green-500': change > 0,
              'text-red-500': change < 0,
            })}
          >
            {change > 0 ? '▲' : '▼'}
            {Math.abs(change).toFixed(2)}%
          </span>{' '}
          | <span>{numberToMoney(close)}</span>
        </p>
      </div>
    </Link>
  );
};

const SearchResults = ({ results, isLoading }: ISearchResults) => {
  return (
    <div className="absolute top-full left-0 w-full bg-gray-800 rounded-md shadow-lg overflow-hidden rounded-t-none min-w-maxz">
      {isLoading ? (
        <SearchResult isLoading={true} />
      ) : typeof results == 'string' ? (
        <p className="text-center text-gray-400 p-4">{'Nenhum resultado :('}</p>
      ) : (
        results.map((result) => (
          <SearchResult key={result.stock} {...result} hasBorder />
        ))
      )}
    </div>
  );
};
