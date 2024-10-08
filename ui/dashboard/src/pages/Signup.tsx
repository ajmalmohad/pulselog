import React, { useState } from 'react';
import { useAuth } from '@app/hooks/useAuth';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { useForm, SubmitHandler } from 'react-hook-form';
import { Eye, EyeOff } from 'lucide-react';
import { Link } from 'react-router-dom';

interface ISignupFormInput {
    name: string;
    email: string;
    password: string;
}

export const SignupPage: React.FC = () => {
    const { register, handleSubmit } = useForm<ISignupFormInput>();
    const { signup } = useAuth();
    const [passwordVisible, setPasswordVisible] = useState(false);

    const onSubmit: SubmitHandler<ISignupFormInput> = async (data) => {
        await signup(data.name, data.email, data.password);
    };

    const togglePasswordVisibility = () => {
        setPasswordVisible(!passwordVisible);
    };

    return (
        <div className='w-full max-w-[600px]'>
            <h2 className="text-4xl font-semibold mb-4">Create an account</h2>
            <p className="mb-8">
                Already have an account?{' '}
                <Link to="/auth/login" replace className="text-purple-500 hover:underline">
                    Log in
                </Link>
            </p>
            <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
                <div className="flex">
                    <Input
                        type="text"
                        placeholder="Name"
                        {...register('name')}
                    />
                </div>
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
                    Create account
                </Button>
            </form>
        </div>
    );
};