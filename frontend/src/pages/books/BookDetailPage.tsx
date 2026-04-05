import { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { booksApi } from '../../api/books';
import { reservationsApi } from '../../api/reservations';
import type { Book } from '../../types';

export default function BookDetailPage() {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const [book, setBook] = useState<Book | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [reserving, setReserving] = useState(false);
  const [dueDate, setDueDate] = useState('');
  const [showReserveForm, setShowReserveForm] = useState(false);
  const [reserveMsg, setReserveMsg] = useState('');

  useEffect(() => {
    if (!id) return;
    setLoading(true);
    booksApi.getById(id)
      .then(({ data }) => setBook(data))
      .catch((err) => {
        if (err.response?.status === 404) {
          navigate('/404', { replace: true });
        } else {
          setError('Kitap yuklenirken hata olustu');
        }
      })
      .finally(() => setLoading(false));
  }, [id, navigate]);

  const handleReserve = async () => {
    if (!id || !dueDate) return;
    setReserving(true);
    setReserveMsg('');
    try {
      await reservationsApi.create(id, dueDate);
      setReserveMsg('Rezervasyon basariyla olusturuldu!');
      setShowReserveForm(false);
      // Refresh book to update available copies
      const { data } = await booksApi.getById(id);
      setBook(data);
    } catch (err: any) {
      setReserveMsg(err.response?.data?.message || 'Rezervasyon olusturulamadi');
    } finally {
      setReserving(false);
    }
  };

  if (loading) {
    return (
      <div className="bg-white rounded-lg shadow p-8 animate-pulse">
        <div className="h-8 bg-gray-200 rounded w-1/3 mb-4" />
        <div className="h-5 bg-gray-200 rounded w-1/4 mb-3" />
        <div className="h-5 bg-gray-200 rounded w-1/5" />
      </div>
    );
  }

  if (error) {
    return (
      <div className="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded">
        {error}
      </div>
    );
  }

  if (!book) return null;

  const minDate = new Date(Date.now() + 86400000).toISOString().split('T')[0];

  return (
    <div>
      <button onClick={() => navigate(-1)} className="text-indigo-600 hover:text-indigo-800 mb-4 inline-block">
        &larr; Geri
      </button>

      <div className="bg-white rounded-lg shadow p-8">
        <div className="flex justify-between items-start">
          <div>
            <h1 className="text-3xl font-bold text-gray-900">{book.title}</h1>
            <p className="text-xl text-gray-600 mt-1">{book.author}</p>
          </div>
          <span
            className={`px-4 py-2 rounded-full text-sm font-medium ${
              book.available ? 'bg-green-100 text-green-700' : 'bg-red-100 text-red-700'
            }`}
          >
            {book.available ? 'Musait' : 'Musait Degil'}
          </span>
        </div>

        <div className="grid grid-cols-2 gap-6 mt-8">
          <div>
            <p className="text-sm text-gray-500">ISBN</p>
            <p className="text-gray-900">{book.isbn}</p>
          </div>
          <div>
            <p className="text-sm text-gray-500">Tur</p>
            <p className="text-gray-900">{book.genre || '-'}</p>
          </div>
          <div>
            <p className="text-sm text-gray-500">Yayin Yili</p>
            <p className="text-gray-900">{book.publishedYear || '-'}</p>
          </div>
          <div>
            <p className="text-sm text-gray-500">Kopya Durumu</p>
            <p className="text-gray-900">{book.availableCopies} / {book.totalCopies} musait</p>
          </div>
        </div>

        {reserveMsg && (
          <div className={`mt-4 px-4 py-3 rounded ${reserveMsg.includes('basari') ? 'bg-green-50 text-green-700' : 'bg-red-50 text-red-700'}`}>
            {reserveMsg}
          </div>
        )}

        {book.available && !showReserveForm && (
          <button
            onClick={() => setShowReserveForm(true)}
            className="mt-6 px-6 py-3 bg-indigo-600 text-white rounded-md hover:bg-indigo-700"
          >
            Rezervasyon Yap
          </button>
        )}

        {showReserveForm && (
          <div className="mt-6 p-4 border border-gray-200 rounded-md">
            <label className="block text-sm font-medium text-gray-700 mb-2">
              Iade Tarihi
            </label>
            <input
              type="date"
              value={dueDate}
              min={minDate}
              onChange={(e) => setDueDate(e.target.value)}
              className="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500"
              aria-label="Iade tarihi"
            />
            <div className="mt-3 flex gap-2">
              <button
                onClick={handleReserve}
                disabled={!dueDate || reserving}
                className="px-4 py-2 bg-indigo-600 text-white rounded-md hover:bg-indigo-700 disabled:opacity-50"
              >
                {reserving ? 'Olusturuluyor...' : 'Onayla'}
              </button>
              <button
                onClick={() => setShowReserveForm(false)}
                className="px-4 py-2 border border-gray-300 rounded-md hover:bg-gray-50"
              >
                Iptal
              </button>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}
