import { Button } from '@/components/ui/button';
import { useAuth } from '@app/hooks/useAuth';
import { Input } from '@components/ui/input';
import { Eye, EyeOff } from 'lucide-react';
import React, { useState } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Link } from 'react-router-dom';

interface ILoginFormInput {
    email: string;
    password: string;
}

export const LoginPage: React.FC = () => {
    const { register, handleSubmit } = useForm<ILoginFormInput>()
    const { login } = useAuth();
    const [passwordVisible, setPasswordVisible] = useState(false);
    const onSubmit: SubmitHandler<ILoginFormInput> = async (data) => {
        await login(data.email, data.password);
    }

    const togglePasswordVisibility = () => {
        setPasswordVisible(!passwordVisible);
    };

    return (
        <div className='w-full max-w-[600px]'>
            <h2 className="text-4xl font-bold mb-4">Welcome back</h2>
            <p className="mb-8">
                Don't have an account?{' '}
                <Link to="/auth/signup" replace className="text-purple-500 hover:underline">
                    Sign Up
                </Link>
            </p>
            <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
                <Input
                    type="email"
                    placeholder="Email"
                    {...register('email')}
                />
                <div className="relative">
                    <Input
                        type={passwordVisible ? 'text' : 'password'}
                        placeholder="Enter your password"
                        {...register('password')}
                    />
                    {passwordVisible ? (
                        <EyeOff
                            className="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 cursor-pointer"
                            onClick={togglePasswordVisibility}
                        />
                    ) : (
                        <Eye
                            className="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 cursor-pointer"
                            onClick={togglePasswordVisibility}
                        />
                    )}
                </div>
                <Button type="submit" className="w-full bg-purple-600 hover:bg-purple-700">
                    Login
                </Button>
            </form>
        </div>
    );
};