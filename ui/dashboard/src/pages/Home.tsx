import { useGetAllProjectsQuery } from '@/api/projects/projectApi';
import React from 'react';

const Home: React.FC = () => {
    const {
        data,
        isLoading,
        error,
    } = useGetAllProjectsQuery();

    if (isLoading) {
        return <div>Loading...</div>;
    }

    if (error) {
        console.error(error);
        return <div>Error</div>;
    }

    return (
        <div>
            <h1 className='text-2xl font-medium'>Project Dashboard</h1>
            <div>
                {data.data.length}
            </div>
        </div>
    );
};

export default Home;