import axios from 'axios';

const api = axios.create({
    baseURL: '/api/v1',
    headers: {
        'Content-Type': 'application/json',
    },
    withCredentials: true,
});

// Request interceptor for adding auth token if needed
api.interceptors.request.use(
    (config) => {
        // Add any request modifications here
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

// Response interceptor for handling common errors
api.interceptors.response.use(
    (response) => response,
    (error) => {
        // Only redirect on 401 for non-auth endpoints
        // The /auth/me endpoint is expected to return 401 for unauthenticated users
        if (error.response?.status === 401 && !error.config?.url?.includes('/auth/me')) {
            // Handle unauthorized (redirect to login, etc.)
            window.location.href = '/login';
        }
        return Promise.reject(error);
    }
);

export { api };
