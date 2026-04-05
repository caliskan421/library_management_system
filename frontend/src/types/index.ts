// User
export interface User {
  _id: string;
  name: string;
  email: string;
  role: 'user' | 'admin';
  createdAt: string;
}

export interface AuthResponse {
  token: string;
  user: User;
}

// Book
export interface Book {
  _id: string;
  title: string;
  author: string;
  isbn: string;
  genre: string;
  publishedYear: number;
  available: boolean;
  totalCopies: number;
  availableCopies: number;
}

export interface CreateBookInput {
  title: string;
  author: string;
  isbn: string;
  genre?: string;
  publishedYear?: number;
  totalCopies: number;
}

export interface UpdateBookInput {
  title?: string;
  author?: string;
  isbn?: string;
  genre?: string;
  publishedYear?: number;
  totalCopies?: number;
}

// Reservation
export interface Reservation {
  _id: string;
  userId: string;
  bookId: string;
  status: 'active' | 'returned';
  reservedAt: string;
  dueDate: string;
  returnedAt: string | null;
}

// Report
export interface ReportSummary {
  total: number;
  active: number;
  returned: number;
}

export interface TopBookReport {
  bookId: string;
  title: string;
  reservationCount: number;
}

export interface ReservationDetailReport {
  bookTitle: string;
  userName: string;
  userEmail: string;
  status: string;
  reservedAt: string;
  dueDate: string;
}

export interface Report {
  generatedAt: string;
  type: string;
  summary: ReportSummary;
  topBooks?: TopBookReport[];
  reservations?: ReservationDetailReport[];
}

// Pagination
export interface PaginatedResponse<T> {
  data: T[];
  page: number;
  limit: number;
  totalCount: number;
}

// Error
export interface ApiError {
  message: string;
}
