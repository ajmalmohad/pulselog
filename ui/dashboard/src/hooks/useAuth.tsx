import { useDispatch, useSelector } from 'react-redux';
import { RootState, AppDispatch } from '@app/store';
import { clearTokens } from '@app/store/auth/authSlice';

export const useAuth = () => {
    const dispatch = useDispatch<AppDispatch>();
    const { accessToken, refreshToken } = useSelector((state: RootState) => state.auth);

    const login = async (email: string, password: string) => {
        try {
            // Replace this with your actual API call
            console.log('Logging in with:', email, password);
            
            return true;
        } catch (error) {
            console.error('Login error:', error);
            return false;
        }
    };

    const logout = () => {
        // Call your logout API here
        dispatch(clearTokens());
    };

    const isAuthenticated = !!accessToken;

    return { login, logout, isAuthenticated, accessToken, refreshToken };
};