import { ModeToggle } from '@/components/themes/mode-toggle';
import React from 'react';

export const LandingPage: React.FC = () => {
    return (
        <div>
            <ModeToggle />
            <h1>Welcome to My Landing Page</h1>
            <p>This is a mini landing page created with React.</p>
            <button>Get Started</button>
        </div>
    );
};
