import { useState, useEffect } from 'react';
import { reportsApi } from '../../api/reports';
import type { Report } from '../../types';

export default function AdminReportsPage() {
  const [report, setReport] = useState<Report | null>(null);
  const [reportType, setReportType] = useState('reservations');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  useEffect(() => {
    setLoading(true);
    setError('');
    reportsApi.get({ type: reportType })
      .then(({ data }) => setReport(data))
      .catch(() => setError('Rapor yuklenirken hata olustu'))
      .finally(() => setLoading(false));
  }, [reportType]);

  return (
    <div>
      <h1 className="text-2xl font-bold text-gray-900 mb-6">Raporlar</h1>

      <div className="flex gap-2 mb-6">
        {[
          { key: 'reservations', label: 'Rezervasyonlar' },
          { key: 'books', label: 'Kitaplar' },
          { key: 'users', label: 'Kullanicilar' },
        ].map(({ key, label }) => (
          <button
            key={key}
            onClick={() => setReportType(key)}
            className={`px-4 py-2 rounded-md text-sm font-medium ${
              reportType === key
                ? 'bg-indigo-600 text-white'
                : 'bg-white border border-gray-300 text-gray-700 hover:bg-gray-50'
            }`}
          >
            {label}
          </button>
        ))}
      </div>

      {error && (
        <div className="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded mb-4">{error}</div>
      )}

      {loading ? (
        <div className="animate-pulse space-y-4">
          <div className="h-32 bg-gray-200 rounded" />
          <div className="h-48 bg-gray-200 rounded" />
        </div>
      ) : report && (
        <>
          <div className="grid grid-cols-3 gap-4 mb-6">
            <div className="bg-white rounded-lg shadow p-6 text-center">
              <p className="text-3xl font-bold text-gray-900">{report.summary.total}</p>
              <p className="text-sm text-gray-500 mt-1">Toplam</p>
            </div>
            <div className="bg-white rounded-lg shadow p-6 text-center">
              <p className="text-3xl font-bold text-blue-600">{report.summary.active}</p>
              <p className="text-sm text-gray-500 mt-1">Aktif</p>
            </div>
            <div className="bg-white rounded-lg shadow p-6 text-center">
              <p className="text-3xl font-bold text-green-600">{report.summary.returned}</p>
              <p className="text-sm text-gray-500 mt-1">Iade Edildi</p>
            </div>
          </div>

          {report.topBooks && report.topBooks.length > 0 && (
            <div className="bg-white rounded-lg shadow p-6">
              <h2 className="text-lg font-semibold text-gray-900 mb-4">En Cok Rezerve Edilen Kitaplar</h2>
              <div className="space-y-3">
                {report.topBooks.map((book, i) => (
                  <div key={book.bookId} className="flex items-center justify-between py-2 border-b border-gray-100 last:border-0">
                    <div className="flex items-center gap-3">
                      <span className="w-8 h-8 bg-indigo-100 text-indigo-700 rounded-full flex items-center justify-center text-sm font-medium">
                        {i + 1}
                      </span>
                      <span className="text-gray-900">{book.title}</span>
                    </div>
                    <span className="text-sm font-medium text-gray-600">
                      {book.reservationCount} rezervasyon
                    </span>
                  </div>
                ))}
              </div>
            </div>
          )}

          <p className="text-xs text-gray-400 mt-4">
            Rapor tarihi: {new Date(report.generatedAt).toLocaleString('tr-TR')}
          </p>
        </>
      )}
    </div>
  );
}
