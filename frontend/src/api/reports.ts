import api from './client';
import type { Report } from '../types';

interface ReportParams {
  type?: string;
  from?: string;
  to?: string;
}

export const reportsApi = {
  get: (params: ReportParams = {}) =>
    api.get<Report>('/reports', { params }),
};
