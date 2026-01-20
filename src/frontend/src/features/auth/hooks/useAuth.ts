import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { api } from '@/lib/api';
import type { LoginInput, RegisterInput, User, AuthResponse } from '../types';
import { useNavigate } from 'react-router-dom';

// Keys
export const authKeys = {
    all: ['auth'] as const,
    user: () => [...authKeys.all, 'user'] as const,
};

// Hook for current user
export const useCurrentUser = () => {
    return useQuery({
        queryKey: authKeys.user(),
        queryFn: async (): Promise<User | null> => {
            try {
                const { data } = await api.get<{ user: User }>('/auth/me');
                return data.user;
            } catch (error) {
                // If 401, return null
                return null;
            }
        },
        retry: false,
        staleTime: 5 * 60 * 1000, // 5 minutes
    });
};

// Hook for login
export const useLogin = () => {
    const queryClient = useQueryClient();
    const navigate = useNavigate();

    return useMutation({
        mutationFn: async (input: LoginInput) => {
            const { data } = await api.post<AuthResponse>('/auth/login', input);
            return data;
        },
        onSuccess: (data) => {
            queryClient.setQueryData(authKeys.user(), data.user);
            navigate('/dashboard');
        },
    });
};

// Hook for register
export const useRegister = () => {
    const queryClient = useQueryClient();
    const navigate = useNavigate();

    return useMutation({
        mutationFn: async (input: RegisterInput) => {
            const { data } = await api.post<AuthResponse>('/auth/register', input);
            return data;
        },
        onSuccess: (data) => {
            queryClient.setQueryData(authKeys.user(), data.user);
            navigate('/dashboard');
        },
    });
};

// Hook for logout
export const useLogout = () => {
    const queryClient = useQueryClient();
    const navigate = useNavigate();

    return useMutation({
        mutationFn: async () => {
            const { data } = await api.post<AuthResponse>('/auth/logout');
            return data;
        },
        onSuccess: () => {
            queryClient.setQueryData(authKeys.user(), null);
            queryClient.clear(); // Clear all cache on logout
            navigate('/login');
        },
    });
};
