import * as React from "react"
import { cn } from "@/lib/utils"
import {
    ResizableHandle,
    ResizablePanel,
    ResizablePanelGroup,
} from "@/components/ui/resizable"
import { TooltipProvider } from "@/components/ui/tooltip"
import { Outlet } from "react-router-dom"

export const DashboardLayout: React.FC = () => {
    const [isCollapsed, setIsCollapsed] = React.useState(false)
    const defaultLayout = [30, 70];
    const navCollapsedSize = 4;

    return (
        <div className="h-screen w-screen">
            <TooltipProvider delayDuration={0}>
                <ResizablePanelGroup
                    direction="horizontal"
                    onLayout={(sizes: number[]) => {
                        document.cookie = `react-resizable-panels:layout:mail=${JSON.stringify(
                            sizes
                        )}`
                    }}
                    className="h-full items-stretch"
                >
                    <ResizablePanel
                        defaultSize={defaultLayout[0]}
                        collapsedSize={navCollapsedSize}
                        collapsible={true}
                        minSize={15}
                        maxSize={20}
                        onCollapse={() => {
                            setIsCollapsed(true)
                            console.log("collapsed")
                            document.cookie = `react-resizable-panels:collapsed=${JSON.stringify(
                                true
                            )}`
                        }}
                        onResize={() => {
                            setIsCollapsed(false)
                            document.cookie = `react-resizable-panels:collapsed=${JSON.stringify(
                                false
                            )}`
                        }}
                        className={cn(
                            isCollapsed &&
                            "min-w-[50px] transition-all duration-300 ease-in-out"
                        )}
                    >
                        <h1>{isCollapsed ? 'Closed' : 'Open'}</h1>
                    </ResizablePanel>
                    <ResizableHandle withHandle />
                    <ResizablePanel defaultSize={defaultLayout[1]} minSize={30}>
                        <Outlet />
                    </ResizablePanel>
                </ResizablePanelGroup>
            </TooltipProvider>
        </div>
    )
};