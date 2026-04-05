import { Link, useNavigate } from 'react-router-dom';
import { useAuthStore } from '../../store/authStore';

export default function Navbar() {
  const { user, logout, isAdmin } = useAuthStore();
  const navigate = useNavigate();

  const handleLogout = () => {
    logout();
    navigate('/login');
  };

  return (
    <nav className="bg-white border-b border-gray-200 px-6 py-3">
      <div className="max-w-7xl mx-auto flex items-center justify-between">
        <Link to="/books" className="text-xl font-bold text-indigo-600">
          LibraNet
        </Link>

        {user && (
          <div className="flex items-center gap-6">
            <Link to="/books" className="text-gray-600 hover:text-gray-900">
              Kitaplar
            </Link>
            {!isAdmin() && (
              <Link to="/reservations" className="text-gray-600 hover:text-gray-900">
                Rezervasyonlar
              </Link>
            )}
            {isAdmin() && (
              <>
                <Link to="/admin/books" className="text-gray-600 hover:text-gray-900">
                  Kitap Yonetimi
                </Link>
                <Link to="/admin/reports" className="text-gray-600 hover:text-gray-900">
                  Raporlar
                </Link>
              </>
            )}
            <div className="flex items-center gap-3 ml-4 pl-4 border-l border-gray-200">
              <span className="text-sm text-gray-500">
                {user.name}
                {isAdmin() && (
                  <span className="ml-1 px-1.5 py-0.5 bg-indigo-100 text-indigo-700 text-xs rounded">
                    Admin
                  </span>
                )}
              </span>
              <button
                onClick={handleLogout}
                className="text-sm text-red-500 hover:text-red-700"
              >
                Cikis
              </button>
            </div>
          </div>
        )}
      </div>
    </nav>
  );
}
