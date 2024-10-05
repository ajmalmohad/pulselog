import React from 'react';
import { Navigate } from 'react-router-dom';
import { useAuth } from '@app/hooks/useAuth';
import useSetupInterceptors from '@app/api/interceptors';
import { identityAPIHandler } from '@app/api/handlers';

interface ProtectedRouteProps {
    children: React.ReactNode;
}

export const ProtectedRoute: React.FC<ProtectedRouteProps> = ({ children }) => {
    useSetupInterceptors(identityAPIHandler);
    
    const { isAuthenticated } = useAuth();

    if (!isAuthenticated) {
        return <Navigate to="/auth/login" replace />;
    }

    return <>{children}</>;
};