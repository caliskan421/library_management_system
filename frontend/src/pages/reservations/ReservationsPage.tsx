import { useState, useEffect, useCallback } from 'react';
import { Link } from 'react-router-dom';
import { reservationsApi } from '../../api/reservations';
import { useAuthStore } from '../../store/authStore';
import type { Reservation } from '../../types';

export default function ReservationsPage() {
  const { user } = useAuthStore();
  const [reservations, setReservations] = useState<Reservation[]>([]);
  const [status, setStatus] = useState('all');
  const [page, setPage] = useState(1);
  const [totalCount, setTotalCount] = useState(0);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');
  const limit = 10;

  const fetchReservations = useCallback(async () => {
    if (!user) return;
    setLoading(true);
    setError('');
    try {
      const { data } = await reservationsApi.listByUser(user._id, { status, page, limit });
      setReservations(data.data);
      setTotalCount(data.totalCount);
    } catch {
      setError('Rezervasyonlar yuklenirken hata olustu');
    } finally {
      setLoading(false);
    }
  }, [user, status, page]);

  useEffect(() => {
    fetchReservations();
  }, [fetchReservations]);

  const totalPages = Math.ceil(totalCount / limit);

  const formatDate = (dateStr: string) => new Date(dateStr).toLocaleDateString('tr-TR');

  return (
    <div>
      <h1 className="text-2xl font-bold text-gray-900 mb-6">Rezervasyonlarim</h1>

      <div className="flex gap-2 mb-6">
        {['all', 'active', 'returned'].map((s) => (
          <button
            key={s}
            onClick={() => { setStatus(s); setPage(1); }}
            className={`px-4 py-2 rounded-md text-sm font-medium ${
              status === s
                ? 'bg-indigo-600 text-white'
                : 'bg-white border border-gray-300 text-gray-700 hover:bg-gray-50'
            }`}
          >
            {s === 'all' ? 'Tumu' : s === 'active' ? 'Aktif' : 'Iade Edildi'}
          </button>
        ))}
      </div>

      {error && (
        <div className="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded mb-4">
          {error}
          <button onClick={fetchReservations} className="ml-2 underline">Tekrar dene</button>
        </div>
      )}

      {loading ? (
        <div className="space-y-4">
          {[...Array(3)].map((_, i) => (
            <div key={i} className="bg-white rounded-lg shadow p-6 animate-pulse">
              <div className="h-5 bg-gray-200 rounded w-1/3 mb-3" />
              <div className="h-4 bg-gray-200 rounded w-1/4" />
            </div>
          ))}
        </div>
      ) : reservations.length === 0 ? (
        <p className="text-gray-500 text-center py-8">Rezervasyon bulunamadi</p>
      ) : (
        <div className="space-y-4">
          {reservations.map((res) => (
            <Link
              key={res._id}
              to={`/reservations/${res._id}`}
              className="block bg-white rounded-lg shadow hover:shadow-md transition p-6"
            >
              <div className="flex justify-between items-center">
                <div>
                  <p className="text-sm text-gray-500">Kitap ID: {res.bookId}</p>
                  <p className="text-sm text-gray-500 mt-1">
                    Rezervasyon: {formatDate(res.reservedAt)} &mdash; Iade: {formatDate(res.dueDate)}
                  </p>
                </div>
                <span
                  className={`px-3 py-1 rounded-full text-sm font-medium ${
                    res.status === 'active'
                      ? 'bg-blue-100 text-blue-700'
                      : 'bg-gray-100 text-gray-600'
                  }`}
                >
                  {res.status === 'active' ? 'Aktif' : 'Iade Edildi'}
                </span>
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
          <span className="px-4 py-2 text-gray-600">{page} / {totalPages}</span>
          <button
            onClick={() => setPage((p) => Math.min(totalPages, p + 1))}
            disabled={page === totalPages}
            className="px-4 py-2 border border-gray-300 rounded-md disabled:opacity-50 hover:bg-gray-50"
          >
            Sonraki
          </button>
        </div>
      )}
    </div>
  );
}
