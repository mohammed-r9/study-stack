import AppSidebar from '@/components/app-sidebar/app-sidebar'
import { ScrollArea } from '@/components/ui/scroll-area'
import { SidebarProvider } from '@/components/ui/sidebar'
import { createFileRoute, Outlet } from '@tanstack/react-router'

export const Route = createFileRoute('/materials')({
  component: RouteComponent,
})

function RouteComponent() {
  return (
    <div className="h-screen w-full bg-background overflow-hidden">
      <div className="flex h-full w-full border border-accent border-t-0">
        <div className="w-72 shrink-0 border-r border-accent">
          <SidebarProvider>
            <AppSidebar />
          </SidebarProvider>
        </div>

        <ScrollArea className="flex-1 h-full overflow-y-auto bg-background">
          <Outlet />
        </ScrollArea>
      </div>
    </div>
  )
}
