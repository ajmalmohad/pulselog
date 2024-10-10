import { LogOut, LucideIcon } from "lucide-react"
import {
    Home,
    Settings
} from "lucide-react"

export type NavbarLink = {
    title: string
    label?: string
    icon: LucideIcon
    variant: "default" | "ghost"
    path?: string
    onClick?: () => void
}

export const navbarLinks: NavbarLink[] = [
    {
        title: "Home",
        icon: Home,
        variant: "ghost",
        path: "/home",
    },
    {
        title: "Settings",
        icon: Settings,
        variant: "ghost",
        path: "/home/settings",
    },
    {
        title: "Logout",
        icon: LogOut,
        variant: "ghost",
    },
]