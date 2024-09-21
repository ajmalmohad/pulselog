import React from 'react';
import { Navigate } from 'react-router-dom';
import { useAuth } from '@app/hooks/useAuth';

interface ProtectedInverseRouteProps {
    children: React.ReactNode;
}

export const ProtectedInverseRoute: React.FC<ProtectedInverseRouteProps> = ({ children }) => {
    const { isAuthenticated } = useAuth();

    if (isAuthenticated) {
        return <Navigate to="/home" replace />;
    }

    return <>{children}</>;
};