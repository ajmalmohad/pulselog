import { useDispatch, useSelector } from 'react-redux';
import { RootState, AppDispatch } from '@app/store';
import { clearTokens, setTokens } from '@app/store/auth/authSlice';
import { identityAPIHandler } from '@app/api/handlers';

export const useAuth = () => {
    const dispatch = useDispatch<AppDispatch>();
    const { accessToken, refreshToken } = useSelector((state: RootState) => state.auth);

    const login = async (email: string, password: string) => {
        try {
            const { data } = await identityAPIHandler.post("/auth/login", {
                email,
                password,
            });

            dispatch(
                setTokens({
                    accessToken: data.data.access_token,
                    refreshToken: data.data.refresh_token
                })
            );

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