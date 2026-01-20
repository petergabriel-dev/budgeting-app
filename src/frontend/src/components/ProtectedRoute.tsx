import { Navigate, Outlet } from 'react-router-dom';
import { useAuthContext } from '../contexts/AuthContext';
import { Loading } from './Loading';

export const ProtectedRoute = () => {
    const { isAuthenticated, isLoading } = useAuthContext();

    if (isLoading) {
        return <Loading />;
    }

    if (!isAuthenticated) {
        return <Navigate to="/login" replace />;
    }

    return <Outlet />;
};
