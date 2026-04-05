import { Link } from 'react-router-dom';

export default function ForbiddenPage() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50">
      <div className="text-center">
        <h1 className="text-6xl font-bold text-gray-300">403</h1>
        <p className="text-xl text-gray-600 mt-4">Bu sayfaya erisim yetkiniz bulunmuyor</p>
        <Link to="/books" className="mt-6 inline-block px-6 py-3 bg-indigo-600 text-white rounded-md hover:bg-indigo-700">
          Ana Sayfaya Don
        </Link>
      </div>
    </div>
  );
}
