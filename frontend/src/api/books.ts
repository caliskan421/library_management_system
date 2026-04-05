import api from './client';
import type { Book, CreateBookInput, UpdateBookInput, PaginatedResponse } from '../types';

interface SearchParams {
  query?: string;
  genre?: string;
  available?: boolean;
  page?: number;
  limit?: number;
}

export const booksApi = {
  search: (params: SearchParams = {}) =>
    api.get<PaginatedResponse<Book>>('/books', { params }),

  getById: (id: string) =>
    api.get<Book>(`/books/${id}`),

  create: (data: CreateBookInput) =>
    api.post<Book>('/books', data),

  update: (id: string, data: UpdateBookInput) =>
    api.put<Book>(`/books/${id}`, data),

  delete: (id: string) =>
    api.delete(`/books/${id}`),
};
