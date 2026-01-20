import { createContext, useContext, type ReactNode } from 'react';
import { useCurrentUser } from '../features/auth/hooks/useAuth';
import type { User } from '../features/auth/types';

interface AuthContextType {
    user: User | null | undefined;
    isLoading: boolean;
    isAuthenticated: boolean;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider = ({ children }: { children: ReactNode }) => {
    const { data: user, isLoading } = useCurrentUser();

    return (
        <AuthContext.Provider
            value={{
                user,
                isLoading,
                isAuthenticated: !!user,
            }}
        >
            {children}
        </AuthContext.Provider>
    );
};

export const useAuthContext = () => {
    const context = useContext(AuthContext);
    if (context === undefined) {
        throw new Error('useAuthContext must be used within an AuthProvider');
    }
    return context;
};
