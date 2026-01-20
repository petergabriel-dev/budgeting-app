import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import { AuthProvider } from '../contexts/AuthContext';
import { ProtectedRoute } from '../components/ProtectedRoute';
import { LoginForm } from '../features/auth/components/LoginForm';
import { RegisterForm } from '../features/auth/components/RegisterForm';
import { DashboardLayout } from '../components/DashboardLayout';

// Placeholder pages - replace with actual feature components as they are built
const DashboardPage = () => (
    <div className="p-8">
        <h1 className="text-2xl font-bold mb-4">Dashboard</h1>
        <p>Welcome to your budgeting app!</p>
    </div>
);

export const AppRoutes = () => {
    return (
        <AuthProvider>
            <BrowserRouter>
                <Routes>
                    {/* Public Routes */}
                    <Route path="/login" element={<LoginForm />} />
                    <Route path="/register" element={<RegisterForm />} />

                    {/* Protected Routes */}
                    <Route element={<ProtectedRoute />}>
                        <Route element={<DashboardLayout />}>
                            <Route path="/dashboard" element={<DashboardPage />} />
                            {/* Default redirect to dashboard */}
                            <Route path="/" element={<Navigate to="/dashboard" replace />} />
                        </Route>
                    </Route>
                </Routes>
            </BrowserRouter>
        </AuthProvider>
    );
};
