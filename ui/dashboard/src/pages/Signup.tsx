import { useAuth } from '@app/hooks/useAuth';
import { Input } from '@components/ui/input';
import React from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';

interface ISignupFormInput {
    name: string;
    email: string;
    password: string;
}

export const SignupPage: React.FC = () => {
    const { register, handleSubmit } = useForm<ISignupFormInput>();
    const { signup } = useAuth();
    const onSubmit: SubmitHandler<ISignupFormInput> = async (data) => {
        await signup(data.name, data.email, data.password);
    }

    return (
        <div>
            <h2>Signup</h2>
            <form onSubmit={handleSubmit(onSubmit)}>
                <div>
                    <label>Name:</label>
                    <Input type="text" {...register("name")} />
                </div>
                <div>
                    <label>Email:</label>
                    <Input type="email" {...register("email")} />
                </div>
                <div>
                    <label>Password:</label>
                    <Input type="password" {...register("password")} />
                </div>
                <button type="submit">Signup</button>
            </form>
        </div>
    );
};