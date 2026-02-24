import AppSidebar from '@/components/app-sidebar/app-sidebar'
import { SidebarProvider } from '@/components/ui/sidebar'
import { createFileRoute, Outlet } from '@tanstack/react-router'

export const Route = createFileRoute('/materials')({
  component: RouteComponent,
})

function RouteComponent() {
  return (
    <div className="h-svh w-full p-2 bg-background">
      <div className="flex h-full w-full rounded-2xl overflow-hidden border border-accent">
        <div className="w-72 z-10 overflow-y-auto border-r border-r-accent shrink-0 overflow-x-hidden">
          <SidebarProvider>
            <AppSidebar />
          </SidebarProvider>
        </div>

        {/* Main Content */}
        <div className="flex-1 h-full overflow-auto bg-background z-30 w-full">
          <Outlet />
        </div>
      </div>
    </div>
  )
}
