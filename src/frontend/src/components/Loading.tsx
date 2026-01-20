import './DashboardLayout.css';

export const Loading = () => {
    return (
        <div className="loading-container" style={{ minHeight: '100vh', backgroundColor: '#0f172a' }}>
            <div className="spinner"></div>
        </div>
    );
};
