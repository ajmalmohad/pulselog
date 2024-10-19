import React from 'react';
import { useParams } from 'react-router-dom';

const Projects: React.FC = () => {
    const { projectId } = useParams<{ projectId: string }>();
    
    return (
        <div>
            <h1 className="text-3xl font-bold underline">
                Hello Project {projectId}!
            </h1>
        </div>
    );
};

export default Projects;