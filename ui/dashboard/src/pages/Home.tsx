import { useGetAllProjectsQuery,useCreateProjectMutation } from '@/api/projects/projectApi';
import React, { useState } from 'react';
import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "@/components/ui/dialog"
import {
    Table,
    TableBody,
    TableCell,
    TableFooter,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table"
import { Plus } from 'lucide-react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { toast } from 'sonner';
import { format } from 'date-fns';

type IProjectFormInput = {
    name: string;
}

const Home: React.FC = () => {
    const { register, handleSubmit } = useForm<IProjectFormInput>()
    const [createProject] = useCreateProjectMutation();
    const [isDialogOpen, setIsDialogOpen] = useState(false);

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

    if (!data) {
        return <div>No data</div>;
    }

    const onSubmit: SubmitHandler<IProjectFormInput> = async (data) => {
        const createProjectPromise = createProject({
            name: data.name,
        });

        toast.promise(createProjectPromise, {
            loading: 'Creating project...',
            success: 'Project created successfully',
            error: 'Failed to create project',
        });
        
        await createProjectPromise.then(() => {
            setIsDialogOpen(false);
        });
    }

    console.log(data);

    return (
        <div>
            <h1 className='text-2xl font-medium mb-4'>Project Dashboard</h1>
            {
                <Table>
                    <TableHeader>
                        <TableRow>
                            <TableHead>Project ID</TableHead>
                            <TableHead>Name</TableHead>
                            <TableHead>Created At</TableHead>
                            <TableHead className="text-right">Owner</TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {data.data.map((project) => (
                            <TableRow key={project.id}>
                                <TableCell className="font-medium">{project.id}</TableCell>
                                <TableCell className="font-medium">{project.name}</TableCell>
                                <TableCell>{format(new Date(project.created_at), 'yyyy-MM-dd HH:mm:ss')}</TableCell>
                                <TableCell className='text-right'>{project.owner.name}</TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                    <TableFooter>
                        <Dialog  open={isDialogOpen} onOpenChange={setIsDialogOpen}>
                            <DialogTrigger asChild>
                                <TableRow className="cursor-pointer">
                                    <TableCell colSpan={4} className='text-gray-500'>
                                        <Plus />
                                    </TableCell>
                                </TableRow>
                            </DialogTrigger>
                            <DialogContent>
                                <DialogHeader>
                                    <DialogTitle>Create Project</DialogTitle>
                                </DialogHeader>
                                <form onSubmit={handleSubmit(onSubmit)}>
                                    <DialogDescription>
                                        <Input type="text" placeholder="Project Name" {...register('name')} />
                                    </DialogDescription>
                                    <DialogFooter>
                                        <Button type="submit" className="mt-4 bg-purple-600 hover:bg-purple-700">
                                            Create Project
                                        </Button>
                                    </DialogFooter>
                                </form>
                            </DialogContent>
                        </Dialog>
                    </TableFooter>
                </Table>
            }
        </div >
    );
};

export default Home;