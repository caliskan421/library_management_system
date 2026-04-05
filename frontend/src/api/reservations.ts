import api from './client';
import type { Reservation, PaginatedResponse } from '../types';

interface ListParams {
  status?: string;
  page?: number;
  limit?: number;
}

export const reservationsApi = {
  create: (bookId: string, dueDate: string) =>
    api.post<Reservation>('/reservations', { bookId, dueDate }),

  getById: (id: string) =>
    api.get<Reservation>(`/reservations/${id}`),

  returnBook: (id: string) =>
    api.delete(`/reservations/${id}`),

  listByUser: (userId: string, params: ListParams = {}) =>
    api.get<PaginatedResponse<Reservation>>(`/users/${userId}/reservations`, { params }),
};
