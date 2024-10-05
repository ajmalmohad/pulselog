import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { ArrowRight } from "lucide-react";
import { Link } from 'react-router-dom';
import React from 'react';

export const LandingPage: React.FC = () => {
    return (
        <div>
            <section className="container w-full">
                <div className="grid place-items-center lg:max-w-screen-xl gap-8 mx-auto py-20 md:py-32">
                    <div className="text-center space-y-8">
                        <Badge variant="outline" className="text-sm py-2">
                            <span className="mr-2 text-[orange]">
                                <Badge>New</Badge>
                            </span>
                            <span> Check it out now! </span>
                        </Badge>

                        <div className="max-w-screen-md mx-auto text-center text-4xl md:text-6xl font-bold">
                            <h1>
                                Track all your
                                <span className="text-transparent px-2 bg-gradient-to-r from-[#D247BF] to-[orange] bg-clip-text">
                                    events
                                </span>
                                in a single place
                            </h1>
                        </div>

                        <p className="max-w-screen-sm mx-auto text-xl text-muted-foreground">
                            Pulselog is a simple and intuitive platform to track all your application's events in a single place.
                        </p>

                        <div className="space-y-4 md:space-y-0 md:space-x-4">
                            <Button asChild className="w-5/6 md:w-1/4 font-bold group/arrow">
                                <Link to="/auth/signup">
                                    Get Started
                                    <ArrowRight className="size-5 ml-2 group-hover/arrow:translate-x-1 transition-transform" />
                                </Link>
                            </Button>

                            <Button
                                asChild
                                variant="secondary"
                                className="w-5/6 md:w-1/4 font-bold"
                            >
                                <Link to="/auth/login">Login</Link>
                            </Button>
                        </div>
                    </div>

                    <div className="relative group mt-14">
                        <div className="absolute top-2 lg:-top-8 left-1/2 transform -translate-x-1/2 w-[90%] mx-auto h-24 lg:h-80 bg-[orange]/50 rounded-full blur-3xl"></div>
                        <img
                            width={1200}
                            height={400}
                            className="w-full md:w-[1200px] mx-auto rounded-lg relative rouded-lg leading-none flex items-center border border-t-2 border-secondary  border-t-[orange]/30"
                            src="/hero-image.jpg"
                            alt="dashboard"
                        />
                        <div className="absolute bottom-0 left-0 w-full h-20 md:h-28 bg-gradient-to-b from-background/0 via-background/50 to-background rounded-lg"></div>
                    </div>
                </div>
            </section>
        </div>
    );
};