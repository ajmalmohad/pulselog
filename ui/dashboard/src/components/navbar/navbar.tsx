import { cn } from "@/lib/utils"
import { buttonVariants } from "@/components/ui/button"
import {
    Tooltip,
    TooltipContent,
    TooltipTrigger,
} from "@/components/ui/tooltip"
import { Link, useLocation } from "react-router-dom"
import { NavbarLink } from "@/data/navbar"

interface NavProps {
    isCollapsed: boolean
    links: NavbarLink[]
}

export function Nav({ links, isCollapsed }: NavProps) {
    const location = useLocation()
    const currentPath = location.pathname

    return (
        <div
            data-collapsed={isCollapsed}
            className="group flex flex-col gap-4 py-2 data-[collapsed=true]:py-2"
        >
            <nav className="grid gap-1 px-2 group-[[data-collapsed=true]]:justify-center group-[[data-collapsed=true]]:px-2">
                {links.map((link, index) => {
                    const isActive = currentPath === link.path
                    const variant = isActive ? "default" : "ghost"

                    return (
                        isCollapsed ? (
                            <Tooltip key={index} delayDuration={0}>
                                <TooltipTrigger asChild>
                                    <Link
                                        onClick={link.onClick || undefined}
                                        to={link.path || "#"}
                                        className={cn(
                                            buttonVariants({ variant: link.variant, size: "icon" }),
                                            "h-9 w-9",
                                            variant === "default" &&
                                            "dark:bg-muted dark:text-muted-foreground dark:hover:bg-muted dark:hover:text-white"
                                        )}
                                    >
                                        <link.icon className="h-4 w-4" />
                                        <span className="sr-only">{link.title}</span>
                                    </Link>
                                </TooltipTrigger>
                                <TooltipContent side="right" className="flex items-center gap-4">
                                    {link.title}
                                    {link.label && (
                                        <span className="ml-auto text-muted-foreground">
                                            {link.label}
                                        </span>
                                    )}
                                </TooltipContent>
                            </Tooltip>
                        ) : (
                            <Link
                                key={index}
                                onClick={link.onClick || undefined}
                                to={link.path || "#"}
                                className={cn(
                                    buttonVariants({ variant: link.variant, size: "sm" }),
                                    variant === "default" &&
                                    "dark:bg-muted dark:text-white dark:hover:bg-muted dark:hover:text-white",
                                    "justify-start"
                                )}
                            >
                                <link.icon className="mr-2 h-4 w-4" />
                                {link.title}
                                {link.label && (
                                    <span
                                        className={cn(
                                            "ml-auto",
                                            link.variant === "default" &&
                                            "text-background dark:text-white"
                                        )}
                                    >
                                        {link.label}
                                    </span>
                                )}
                            </Link>
                        )
                    )
                }
                )}
            </nav>
        </div>
    )
}