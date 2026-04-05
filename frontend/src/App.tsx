import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import { lazy, Suspense } from 'react';
import Layout from './components/layout/Layout';
import ProtectedRoute from './components/ui/ProtectedRoute';

const LoginPage = lazy(() => import('./pages/auth/LoginPage'));
const RegisterPage = lazy(() => import('./pages/auth/RegisterPage'));
const BooksPage = lazy(() => import('./pages/books/BooksPage'));
const BookDetailPage = lazy(() => import('./pages/books/BookDetailPage'));
const ReservationsPage = lazy(() => import('./pages/reservations/ReservationsPage'));
const ReservationDetailPage = lazy(() => import('./pages/reservations/ReservationDetailPage'));
const AdminBooksPage = lazy(() => import('./pages/admin/AdminBooksPage'));
const AdminReportsPage = lazy(() => import('./pages/admin/AdminReportsPage'));
const NotFoundPage = lazy(() => import('./pages/NotFoundPage'));
const ForbiddenPage = lazy(() => import('./pages/ForbiddenPage'));

function LoadingFallback() {
  return (
    <div className="flex items-center justify-center min-h-screen">
      <div className="animate-spin h-8 w-8 border-4 border-indigo-600 border-t-transparent rounded-full" />
    </div>
  );
}

export default function App() {
  return (
    <BrowserRouter>
      <Suspense fallback={<LoadingFallback />}>
        <Routes>
          {/* Public routes */}
          <Route path="/login" element={<LoginPage />} />
          <Route path="/register" element={<RegisterPage />} />

          {/* Protected routes */}
          <Route element={<ProtectedRoute><Layout /></ProtectedRoute>}>
            <Route path="/books" element={<BooksPage />} />
            <Route path="/books/:id" element={<BookDetailPage />} />
            <Route path="/reservations" element={<ReservationsPage />} />
            <Route path="/reservations/:id" element={<ReservationDetailPage />} />

            {/* Admin routes */}
            <Route path="/admin/books" element={
              <ProtectedRoute adminOnly><AdminBooksPage /></ProtectedRoute>
            } />
            <Route path="/admin/reports" element={
              <ProtectedRoute adminOnly><AdminReportsPage /></ProtectedRoute>
            } />
          </Route>

          {/* Error pages */}
          <Route path="/403" element={<ForbiddenPage />} />
          <Route path="/404" element={<NotFoundPage />} />
          <Route path="/" element={<Navigate to="/books" replace />} />
          <Route path="*" element={<NotFoundPage />} />
        </Routes>
      </Suspense>
    </BrowserRouter>
  );
}
