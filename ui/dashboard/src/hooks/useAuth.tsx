import { useDispatch, useSelector } from 'react-redux';
import { RootState, AppDispatch } from '@app/store';
import { clearTokens, setTokens } from '@app/store/auth/authSlice';
import { identityAPIHandler } from '@app/api/handlers';

export const useAuth = () => {
    const dispatch = useDispatch<AppDispatch>();
    const { accessToken, refreshToken } = useSelector((state: RootState) => state.auth);

    const signup = async (name: string, email: string, password: string) => {
        try {
            const { data } = await identityAPIHandler.post("/auth/signup", {
                name,
                email,
                password,
            });

            dispatch(
                setTokens({
                    accessToken: data.data.access_token,
                    refreshToken: data.data.refresh_token
                })
            );
        } catch (error) {
            console.error('Signup error:', error);
        }
    }

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
        } catch (error) {
            console.error('Login error:', error);
        }
    };

    const logout = async () => {
        try {
            const { data } = await identityAPIHandler.delete("/users/logout", {
                data: {
                    refresh_token: refreshToken
                }
            });

            dispatch(
                clearTokens()
            );
        } catch (error) {
            console.error('Logout error:', error);
        }
    };

    const isAuthenticated = !!accessToken;

    return { login, logout, signup, isAuthenticated, accessToken, refreshToken };
};