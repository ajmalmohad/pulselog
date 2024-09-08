import { useDispatch, useSelector } from 'react-redux';
import { RootState, AppDispatch } from '@app/store';
import { setTokens, clearTokens } from '@app/store/auth/authSlice';

export const useAuth = () => {
    const dispatch = useDispatch<AppDispatch>();
    const { accessToken, refreshToken } = useSelector((state: RootState) => state.auth);

    const login = async (username: string, password: string) => {
        try {
            // Replace this with your actual API call
            const response = await fetch('/api/login', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ username, password }),
            });

            if (!response.ok) {
                throw new Error('Login failed');
            }

            const data = await response.json();
            dispatch(setTokens({
                accessToken: data.accessToken,
                refreshToken: data.refreshToken,
            }));

            return true;
        } catch (error) {
            console.error('Login error:', error);
            return false;
        }
    };

    const logout = () => {
        dispatch(clearTokens());
    };

    const isAuthenticated = !!accessToken;

    return { login, logout, isAuthenticated, accessToken, refreshToken };
};