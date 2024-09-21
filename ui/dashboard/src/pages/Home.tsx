import { useAuth } from '@app/hooks/useAuth';
import React from 'react';

const Home: React.FC = () => {
    const { logout } = useAuth();
    return (
        <div>
            <h1>Welcome to the Home page!</h1>
            <button
                onClick={() => {
                    logout();
                }}
            >Logout</button>
        </div>
    );
};

export default Home;