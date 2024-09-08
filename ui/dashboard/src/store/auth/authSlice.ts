import { createSlice, PayloadAction } from '@reduxjs/toolkit';

type Tokens = {
    accessToken: string | null;
    refreshToken: string | null;
};

interface AuthState extends Tokens {}

const initialState: AuthState = {
    accessToken: null,
    refreshToken: null,
};

const authSlice = createSlice({
    name: 'auth',
    initialState,
    reducers: {
        setTokens: (state, action: PayloadAction<Partial<Tokens>>) => {
            if (action.payload.accessToken !== undefined) {
                state.accessToken = action.payload.accessToken;
            }
            if (action.payload.refreshToken !== undefined) {
                state.refreshToken = action.payload.refreshToken;
            }
        },
        clearTokens: (state) => {
            state.accessToken = null;
            state.refreshToken = null;
        },
    },
});

export const { setTokens, clearTokens } = authSlice.actions;

export default authSlice.reducer;