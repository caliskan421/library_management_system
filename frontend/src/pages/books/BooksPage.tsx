import { useState, useEffect, useCallback } from 'react';
import { Link } from 'react-router-dom';
import { booksApi } from '../../api/books';
import type { Book } from '../../types';

export default function BooksPage() {
  const [books, setBooks] = useState<Book[]>([]);
  const [query, setQuery] = useState('');
  const [genre, setGenre] = useState('');
  const [page, setPage] = useState(1);
  const [totalCount, setTotalCount] = useState(0);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');
  const limit = 10;

  const fetchBooks = useCallback(async () => {
    setLoading(true);
    setError('');
    try {
      const params: Record<string, any> = { page, limit };
      if (query) params.query = query;
      if (genre) params.genre = genre;
      const { data } = await booksApi.search(params);
      setBooks(data.data);
      setTotalCount(data.totalCount);
    } catch {
      setError('Kitaplar yuklenirken hata olustu');
    } finally {
      setLoading(false);
    }
  }, [query, genre, page]);

  useEffect(() => {
    const timer = setTimeout(fetchBooks, 400);
    return () => clearTimeout(timer);
  }, [fetchBooks]);

  const totalPages = Math.ceil(totalCount / limit);

  return (
    <div>
      <h1 className="text-2xl font-bold text-gray-900 mb-6">Kitaplar</h1>

      <div className="flex gap-4 mb-6">
        <input
          type="text"
          value={query}
          onChange={(e) => { setQuery(e.target.value); setPage(1); }}
          placeholder="Kitap adi, yazar veya ISBN ara..."
          className="flex-1 px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500"
          aria-label="Kitap ara"
        />
        <input
          type="text"
          value={genre}
          onChange={(e) => { setGenre(e.target.value); setPage(1); }}
          placeholder="Tur filtresi"
          className="w-48 px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500"
          aria-label="Tur filtresi"
        />
      </div>

      {error && (
        <div className="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded mb-4">
          {error}
          <button onClick={fetchBooks} className="ml-2 underline">Tekrar dene</button>
        </div>
      )}

      {loading ? (
        <div className="space-y-4">
          {[...Array(3)].map((_, i) => (
            <div key={i} className="bg-white rounded-lg shadow p-6 animate-pulse">
              <div className="h-5 bg-gray-200 rounded w-1/3 mb-3" />
              <div className="h-4 bg-gray-200 rounded w-1/4 mb-2" />
              <div className="h-4 bg-gray-200 rounded w-1/5" />
            </div>
          ))}
        </div>
      ) : (
        <>
          {books.length === 0 ? (
            <p className="text-gray-500 text-center py-8">Kitap bulunamadi</p>
          ) : (
            <div className="space-y-4">
              {books.map((book) => (
                <Link
                  key={book._id}
                  to={`/books/${book._id}`}
                  className="block bg-white rounded-lg shadow hover:shadow-md transition p-6"
                >
                  <div className="flex justify-between items-start">
                    <div>
                      <h2 className="text-lg font-semibold text-gray-900">{book.title}</h2>
                      <p className="text-gray-600">{book.author}</p>
                      <p className="text-sm text-gray-400 mt-1">ISBN: {book.isbn}</p>
                      {book.genre && (
                        <span className="inline-block mt-2 px-2 py-1 bg-gray-100 text-gray-600 text-xs rounded">
                          {book.genre}
                        </span>
                      )}
                    </div>
                    <div className="text-right">
                      <span
                        className={`inline-block px-3 py-1 rounded-full text-sm font-medium ${
                          book.available
                            ? 'bg-green-100 text-green-700'
                            : 'bg-red-100 text-red-700'
                        }`}
                      >
                        {book.available ? 'Musait' : 'Musait Degil'}
                      </span>
                      <p className="text-sm text-gray-500 mt-1">
                        {book.availableCopies}/{book.totalCopies} kopya
                      </p>
                    </div>
                  </div>
                </Link>
              ))}
            </div>
          )}

          {totalPages > 1 && (
            <div className="flex justify-center gap-2 mt-6">
              <button
                onClick={() => setPage((p) => Math.max(1, p - 1))}
                disabled={page === 1}
                className="px-4 py-2 border border-gray-300 rounded-md disabled:opacity-50 hover:bg-gray-50"
              >
                Onceki
              </button>
              <span className="px-4 py-2 text-gray-600">
                {page} / {totalPages}
              </span>
              <button
                onClick={() => setPage((p) => Math.min(totalPages, p + 1))}
                disabled={page === totalPages}
                className="px-4 py-2 border border-gray-300 rounded-md disabled:opacity-50 hover:bg-gray-50"
              >
                Sonraki
              </button>
            </div>
          )}
        </>
      )}
    </div>
  );
}
