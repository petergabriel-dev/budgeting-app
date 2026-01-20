import { useState } from 'react';
import { Outlet, NavLink, useLocation } from 'react-router-dom';
import {
    LayoutDashboard,
    Settings,
    LogOut,
    Menu,
    Users,
    Package,
    PieChart,
    FileText
} from 'lucide-react';
import { useAuthContext } from '../contexts/AuthContext';
import { useLogout } from '../features/auth/hooks/useAuth';
import './DashboardLayout.css';

export const DashboardLayout = () => {
    const { user } = useAuthContext();
    const logoutMutation = useLogout();
    const [isSidebarOpen, setIsSidebarOpen] = useState(false);
    const location = useLocation();

    const toggleSidebar = () => setIsSidebarOpen(!isSidebarOpen);
    const closeSidebar = () => setIsSidebarOpen(false);

    const handleLogout = () => {
        logoutMutation.mutate();
    };

    // Define navigation items
    // If role is missing, default to standard user view or limited view
    const role = user?.role || 'user';

    const getNavItems = () => {
        const items = [
            {
                icon: LayoutDashboard,
                label: role === 'admin' ? 'Overview' : 'Dashboard',
                path: '/dashboard'
            }
        ];

        if (role === 'admin') {
            items.push(
                { icon: Users, label: 'Clients', path: '/admin/clients' },
                { icon: Package, label: 'Packages', path: '/admin/packages' },
                { icon: PieChart, label: 'Reports', path: '/admin/reports' }
            );
        } else {
            // Standard User Links
            items.push(
                { icon: FileText, label: 'My Plan', path: '/plan' }
            );
        }

        // Common items at bottom usually, but here we just append
        items.push({ icon: Settings, label: 'Settings', path: '/settings' });

        return items;
    };

    const navItems = getNavItems();

    const getPageTitle = () => {
        const current = navItems.find(item => item.path === location.pathname);
        if (current) return current.label;
        if (location.pathname.startsWith('/admin')) return 'Admin Area';
        return 'Dashboard';
    };

    return (
        <div className="dashboard-container">
            {/* Mobile Overlay */}
            <div
                className={`sidebar-overlay ${isSidebarOpen ? 'open' : ''}`}
                onClick={closeSidebar}
                aria-hidden="true"
            />

            {/* Sidebar */}
            <aside className={`dashboard-sidebar ${isSidebarOpen ? 'open' : ''}`}>
                <div className="sidebar-header">
                    <span>BudgetApp</span>
                </div>

                <nav className="sidebar-nav">
                    {navItems.map((item) => (
                        <NavLink
                            key={item.path}
                            to={item.path}
                            className={({ isActive }) => `nav-item ${isActive ? 'active' : ''}`}
                            onClick={closeSidebar}
                        >
                            <item.icon size={20} />
                            <span>{item.label}</span>
                        </NavLink>
                    ))}
                </nav>

                <div className="sidebar-footer">
                    <div className="user-profile">
                        <div className="user-avatar">
                            {user?.email?.charAt(0).toUpperCase() || 'U'}
                        </div>
                        <div className="user-info">
                            <span className="user-name" title={user?.email}>{user?.email?.split('@')[0]}</span>
                            <span className="user-email" title={role}>{role}</span>
                        </div>
                    </div>
                    <button
                        onClick={handleLogout}
                        className="nav-item"
                        style={{ width: '100%', marginTop: '8px', color: '#ef4444', justifyContent: 'flex-start', background: 'transparent', border: 'none' }}
                    >
                        <LogOut size={20} />
                        <span>Logout</span>
                    </button>
                </div>
            </aside>

            {/* Main Content */}
            <main className="dashboard-main">
                <header className="dashboard-header">
                    <div style={{ display: 'flex', alignItems: 'center', gap: '16px' }}>
                        <button
                            className="toggle-sidebar-btn"
                            onClick={toggleSidebar}
                            aria-label="Toggle Sidebar"
                        >
                            <Menu size={24} />
                        </button>
                        <h1 className="header-title">{getPageTitle()}</h1>
                    </div>

                    {/* Placeholder for header actions */}
                    <div>
                        {/* e.g. Notifications */}
                    </div>
                </header>

                <div className="dashboard-content">
                    <Outlet />
                </div>
            </main>
        </div>
    );
};
