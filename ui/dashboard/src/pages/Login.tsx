import { useAuth } from '@app/hooks/useAuth';
import { Input } from '@components/ui/input';
import React from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';

interface ILoginFormInput {
    email: string;
    password: string;
}

export const LoginPage: React.FC = () => {
    const { register, handleSubmit } = useForm<ILoginFormInput>()
    const { login } = useAuth();
    const onSubmit: SubmitHandler<ILoginFormInput> = async (data) => {
        await login(data.email, data.password);
    }

    return (
        <div>
            <h2>Login</h2>
            <form onSubmit={handleSubmit(onSubmit)}>
                <div>
                    <label>Email:</label>
                    <Input type="email" {...register("email")} />
                </div>
                <div>
                    <label>Password:</label>
                    <Input type="password" {...register("password")} /> 
                </div>
                <button type="submit">Login</button>
            </form>
        </div>
    );
};