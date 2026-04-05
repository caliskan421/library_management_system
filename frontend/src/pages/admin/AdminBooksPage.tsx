import { useState, useEffect, useCallback } from 'react';
import { booksApi } from '../../api/books';
import type { Book, CreateBookInput } from '../../types';

export default function AdminBooksPage() {
  const [books, setBooks] = useState<Book[]>([]);
  const [page, setPage] = useState(1);
  const [totalCount, setTotalCount] = useState(0);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');
  const [showForm, setShowForm] = useState(false);
  const [editingBook, setEditingBook] = useState<Book | null>(null);
  const [formData, setFormData] = useState<CreateBookInput>({
    title: '', author: '', isbn: '', genre: '', publishedYear: 0, totalCopies: 1,
  });
  const [formError, setFormError] = useState('');
  const [submitting, setSubmitting] = useState(false);
  const limit = 10;

  const fetchBooks = useCallback(async () => {
    setLoading(true);
    try {
      const { data } = await booksApi.search({ page, limit });
      setBooks(data.data);
      setTotalCount(data.totalCount);
    } catch {
      setError('Kitaplar yuklenirken hata olustu');
    } finally {
      setLoading(false);
    }
  }, [page]);

  useEffect(() => { fetchBooks(); }, [fetchBooks]);

  const resetForm = () => {
    setFormData({ title: '', author: '', isbn: '', genre: '', publishedYear: 0, totalCopies: 1 });
    setEditingBook(null);
    setShowForm(false);
    setFormError('');
  };

  const openEdit = (book: Book) => {
    setEditingBook(book);
    setFormData({
      title: book.title,
      author: book.author,
      isbn: book.isbn,
      genre: book.genre || '',
      publishedYear: book.publishedYear || 0,
      totalCopies: book.totalCopies,
    });
    setShowForm(true);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setFormError('');
    if (!formData.title || !formData.author || !formData.isbn || formData.totalCopies < 1) {
      setFormError('Baslik, yazar, ISBN ve kopya sayisi zorunludur');
      return;
    }
    setSubmitting(true);
    try {
      if (editingBook) {
        await booksApi.update(editingBook._id, formData);
      } else {
        await booksApi.create(formData);
      }
      resetForm();
      fetchBooks();
    } catch (err: any) {
      setFormError(err.response?.data?.message || 'Islem basarisiz');
    } finally {
      setSubmitting(false);
    }
  };

  const handleDelete = async (id: string) => {
    if (!confirm('Bu kitabi silmek istediginizden emin misiniz?')) return;
    try {
      await booksApi.delete(id);
      fetchBooks();
    } catch (err: any) {
      alert(err.response?.data?.message || 'Silme basarisiz');
    }
  };

  const totalPages = Math.ceil(totalCount / limit);

  return (
    <div>
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-2xl font-bold text-gray-900">Kitap Yonetimi</h1>
        <button
          onClick={() => { resetForm(); setShowForm(true); }}
          className="px-4 py-2 bg-indigo-600 text-white rounded-md hover:bg-indigo-700"
        >
          + Yeni Kitap
        </button>
      </div>

      {showForm && (
        <div className="bg-white rounded-lg shadow p-6 mb-6">
          <h2 className="text-lg font-semibold mb-4">
            {editingBook ? 'Kitap Duzenle' : 'Yeni Kitap Ekle'}
          </h2>
          {formError && (
            <div className="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded mb-4">{formError}</div>
          )}
          <form onSubmit={handleSubmit} className="grid grid-cols-2 gap-4">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">Baslik *</label>
              <input
                type="text" value={formData.title}
                onChange={(e) => setFormData({ ...formData, title: e.target.value })}
                className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500"
                required aria-label="Kitap basligi"
              />
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">Yazar *</label>
              <input
                type="text" value={formData.author}
                onChange={(e) => setFormData({ ...formData, author: e.target.value })}
                className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500"
                required aria-label="Yazar"
              />
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">ISBN *</label>
              <input
                type="text" value={formData.isbn}
                onChange={(e) => setFormData({ ...formData, isbn: e.target.value })}
                className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500"
                required aria-label="ISBN"
              />
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">Tur</label>
              <input
                type="text" value={formData.genre}
                onChange={(e) => setFormData({ ...formData, genre: e.target.value })}
                className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500"
                aria-label="Tur"
              />
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">Yayin Yili</label>
              <input
                type="number" value={formData.publishedYear || ''}
                onChange={(e) => setFormData({ ...formData, publishedYear: parseInt(e.target.value) || 0 })}
                className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500"
                aria-label="Yayin yili"
              />
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">Kopya Sayisi *</label>
              <input
                type="number" value={formData.totalCopies} min={1}
                onChange={(e) => setFormData({ ...formData, totalCopies: parseInt(e.target.value) || 1 })}
                className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500"
                required aria-label="Kopya sayisi"
              />
            </div>
            <div className="col-span-2 flex gap-2">
              <button type="submit" disabled={submitting}
                className="px-4 py-2 bg-indigo-600 text-white rounded-md hover:bg-indigo-700 disabled:opacity-50">
                {submitting ? 'Kaydediliyor...' : editingBook ? 'Guncelle' : 'Ekle'}
              </button>
              <button type="button" onClick={resetForm}
                className="px-4 py-2 border border-gray-300 rounded-md hover:bg-gray-50">
                Iptal
              </button>
            </div>
          </form>
        </div>
      )}

      {error && (
        <div className="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded mb-4">{error}</div>
      )}

      {loading ? (
        <div className="animate-pulse space-y-3">
          {[...Array(3)].map((_, i) => <div key={i} className="h-16 bg-gray-200 rounded" />)}
        </div>
      ) : (
        <div className="bg-white rounded-lg shadow overflow-hidden">
          <table className="w-full">
            <thead className="bg-gray-50">
              <tr>
                <th className="px-4 py-3 text-left text-sm font-medium text-gray-600">Baslik</th>
                <th className="px-4 py-3 text-left text-sm font-medium text-gray-600">Yazar</th>
                <th className="px-4 py-3 text-left text-sm font-medium text-gray-600">ISBN</th>
                <th className="px-4 py-3 text-center text-sm font-medium text-gray-600">Kopya</th>
                <th className="px-4 py-3 text-right text-sm font-medium text-gray-600">Islemler</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-gray-200">
              {books.map((book) => (
                <tr key={book._id} className="hover:bg-gray-50">
                  <td className="px-4 py-3 text-sm text-gray-900">{book.title}</td>
                  <td className="px-4 py-3 text-sm text-gray-600">{book.author}</td>
                  <td className="px-4 py-3 text-sm text-gray-500 font-mono">{book.isbn}</td>
                  <td className="px-4 py-3 text-sm text-center text-gray-600">
                    {book.availableCopies}/{book.totalCopies}
                  </td>
                  <td className="px-4 py-3 text-right space-x-2">
                    <button onClick={() => openEdit(book)}
                      className="text-sm text-indigo-600 hover:text-indigo-800">Duzenle</button>
                    <button onClick={() => handleDelete(book._id)}
                      className="text-sm text-red-600 hover:text-red-800">Sil</button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}

      {totalPages > 1 && (
        <div className="flex justify-center gap-2 mt-6">
          <button onClick={() => setPage((p) => Math.max(1, p - 1))} disabled={page === 1}
            className="px-4 py-2 border border-gray-300 rounded-md disabled:opacity-50 hover:bg-gray-50">Onceki</button>
          <span className="px-4 py-2 text-gray-600">{page} / {totalPages}</span>
          <button onClick={() => setPage((p) => Math.min(totalPages, p + 1))} disabled={page === totalPages}
            className="px-4 py-2 border border-gray-300 rounded-md disabled:opacity-50 hover:bg-gray-50">Sonraki</button>
        </div>
      )}
    </div>
  );
}
