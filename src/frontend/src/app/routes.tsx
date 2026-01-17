import { BrowserRouter, Routes, Route } from 'react-router-dom';

// Placeholder pages - replace with actual feature components
const HomePage = () => <div>Home</div>;
const LoginPage = () => <div>Login</div>;
const DashboardPage = () => <div>Dashboard</div>;

export const AppRoutes = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<HomePage />} />
                <Route path="/login" element={<LoginPage />} />
                <Route path="/dashboard" element={<DashboardPage />} />
                {/* Add more routes as features are built */}
            </Routes>
        </BrowserRouter>
    );
};
