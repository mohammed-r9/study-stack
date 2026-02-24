import AppSidebar from '@/components/app-sidebar/app-sidebar'
import { ScrollArea } from '@/components/ui/scroll-area'
import { SidebarProvider } from '@/components/ui/sidebar'
import { createFileRoute, Outlet } from '@tanstack/react-router'

export const Route = createFileRoute('/materials')({
  component: RouteComponent,
})

function RouteComponent() {
  return (
    <div className="h-svh w-full p-2 bg-background pt-0">
      <div className="flex h-full w-full overflow-hidden border border-accent border-t-0">
        <div className="w-72 z-10 overflow-y-auto border-r border-r-accent shrink-0 overflow-x-hidden">
          <SidebarProvider>
            <AppSidebar />
          </SidebarProvider>
        </div>

        {/* Main Content */}
        <ScrollArea className="flex-1 h-full bg-background z-30 w-full">
          <Outlet />
        </ScrollArea>
      </div>
    </div>
  )
}
