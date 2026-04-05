import { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { reservationsApi } from '../../api/reservations';
import type { Reservation } from '../../types';

export default function ReservationDetailPage() {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const [reservation, setReservation] = useState<Reservation | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [returning, setReturning] = useState(false);
  const [returnMsg, setReturnMsg] = useState('');

  useEffect(() => {
    if (!id) return;
    reservationsApi.getById(id)
      .then(({ data }) => setReservation(data))
      .catch((err) => {
        if (err.response?.status === 404) navigate('/404', { replace: true });
        else if (err.response?.status === 403) navigate('/403', { replace: true });
        else setError('Rezervasyon yuklenirken hata olustu');
      })
      .finally(() => setLoading(false));
  }, [id, navigate]);

  const handleReturn = async () => {
    if (!id) return;
    setReturning(true);
    setReturnMsg('');
    try {
      await reservationsApi.returnBook(id);
      setReturnMsg('Kitap basariyla iade edildi!');
      const { data } = await reservationsApi.getById(id);
      setReservation(data);
    } catch (err: any) {
      setReturnMsg(err.response?.data?.message || 'Iade islemi basarisiz');
    } finally {
      setReturning(false);
    }
  };

  const formatDate = (dateStr: string) => new Date(dateStr).toLocaleDateString('tr-TR');

  if (loading) {
    return (
      <div className="bg-white rounded-lg shadow p-8 animate-pulse">
        <div className="h-6 bg-gray-200 rounded w-1/3 mb-4" />
        <div className="h-4 bg-gray-200 rounded w-1/4" />
      </div>
    );
  }

  if (error) {
    return <div className="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded">{error}</div>;
  }

  if (!reservation) return null;

  return (
    <div>
      <button onClick={() => navigate(-1)} className="text-indigo-600 hover:text-indigo-800 mb-4 inline-block">
        &larr; Geri
      </button>

      <div className="bg-white rounded-lg shadow p-8">
        <div className="flex justify-between items-start">
          <h1 className="text-2xl font-bold text-gray-900">Rezervasyon Detayi</h1>
          <span
            className={`px-3 py-1 rounded-full text-sm font-medium ${
              reservation.status === 'active'
                ? 'bg-blue-100 text-blue-700'
                : 'bg-gray-100 text-gray-600'
            }`}
          >
            {reservation.status === 'active' ? 'Aktif' : 'Iade Edildi'}
          </span>
        </div>

        <div className="grid grid-cols-2 gap-6 mt-8">
          <div>
            <p className="text-sm text-gray-500">Rezervasyon ID</p>
            <p className="text-gray-900 font-mono text-sm">{reservation._id}</p>
          </div>
          <div>
            <p className="text-sm text-gray-500">Kitap ID</p>
            <p className="text-gray-900 font-mono text-sm">{reservation.bookId}</p>
          </div>
          <div>
            <p className="text-sm text-gray-500">Rezervasyon Tarihi</p>
            <p className="text-gray-900">{formatDate(reservation.reservedAt)}</p>
          </div>
          <div>
            <p className="text-sm text-gray-500">Iade Tarihi</p>
            <p className="text-gray-900">{formatDate(reservation.dueDate)}</p>
          </div>
          {reservation.returnedAt && (
            <div>
              <p className="text-sm text-gray-500">Iade Edildi</p>
              <p className="text-gray-900">{formatDate(reservation.returnedAt)}</p>
            </div>
          )}
        </div>

        {returnMsg && (
          <div className={`mt-4 px-4 py-3 rounded ${returnMsg.includes('basari') ? 'bg-green-50 text-green-700' : 'bg-red-50 text-red-700'}`}>
            {returnMsg}
          </div>
        )}

        {reservation.status === 'active' && (
          <button
            onClick={handleReturn}
            disabled={returning}
            className="mt-6 px-6 py-3 bg-red-600 text-white rounded-md hover:bg-red-700 disabled:opacity-50"
          >
            {returning ? 'Iade ediliyor...' : 'Kitabi Iade Et'}
          </button>
        )}
      </div>
    </div>
  );
}
