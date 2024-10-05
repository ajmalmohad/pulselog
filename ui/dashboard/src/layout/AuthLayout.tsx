import React from 'react';
import { Outlet } from 'react-router-dom';


export const AuthLayout: React.FC = () => {
    return (
        <div className="flex h-screen">
            <div className="w-1/2 relative rounded p-5">
                <img
                    src="/auth-screen.jpg"
                    alt="Image"
                    className="object-cover h-full w-full rounded-md"
                />
                <div className="absolute top-10 left-10">
                    <h1 className="text-xl font-bold uppercase drop-shadow-xl">Pulselog</h1>
                </div>
            </div>
            <div className="w-1/2 p-12 flex flex-col items-center justify-center">
                <div className='w-full max-w-[600px]'>
                    <Outlet />
                </div>
            </div>
        </div>
    );
};