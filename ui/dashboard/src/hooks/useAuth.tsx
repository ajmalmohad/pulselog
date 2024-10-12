import { useDispatch, useSelector } from 'react-redux';
import { RootState, AppDispatch } from '@app/store';
import { clearTokens, setTokens } from '@app/store/auth/authSlice';
import { identityAPIHandler } from '@app/api/handlers';
import { toast } from 'sonner';

export const useAuth = () => {
    const dispatch = useDispatch<AppDispatch>();
    const { accessToken, refreshToken } = useSelector((state: RootState) => state.auth);

    const signup = async (name: string, email: string, password: string) => {
        const signupPromise = identityAPIHandler.post("/auth/signup", { name, email, password });

        toast.promise(
            signupPromise,
            {
                loading: 'Signing up...',
                success: 'Signed up successfully!',
                error: 'Signup failed. Please try again.',
            }
        );

        await signupPromise.then(({ data }) => {
            dispatch(
                setTokens({
                    accessToken: data.data.access_token,
                    refreshToken: data.data.refresh_token
                })
            );
        });
    };

    const login = async (email: string, password: string) => {
        const loginPromise = identityAPIHandler.post("/auth/login", { email, password });

        toast.promise(
            loginPromise,
            {
                loading: 'Logging in...',
                success: 'Logged in successfully!',
                error: 'Login failed. Please try again.',
            }
        );

        await loginPromise.then(({ data }) => {
            dispatch(
                setTokens({
                    accessToken: data.data.access_token,
                    refreshToken: data.data.refresh_token
                })
            );
        });
    };

    const logout = async () => {
        const logoutPromise = identityAPIHandler.delete("/users/logout", {
            data: { refresh_token: refreshToken }
        });

        toast.promise(
            logoutPromise,
            {
                loading: 'Logging out...',
                success: 'Logged out successfully!',
                error: 'Logout failed. Please try again.',
            }
        );

        await logoutPromise.then(() => {
            dispatch(clearTokens());
        });
    };

    const isAuthenticated = !!accessToken;

    return { login, logout, signup, isAuthenticated, accessToken, refreshToken };
};