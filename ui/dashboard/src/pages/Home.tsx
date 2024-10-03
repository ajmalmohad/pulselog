import { useAuth } from '@app/hooks/useAuth';
import React from 'react';

const Home: React.FC = () => {
    const { logout } = useAuth();
    return (
        <div>
            <h1 className="text-3xl font-bold underline">
                Hello world!
            </h1>
            <button
                onClick={() => {
                    logout();
                }}
            >Logout</button>
        </div>
    );
};

export default Home;