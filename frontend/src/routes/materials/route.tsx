import AppSidebar from '@/components/app-sidebar'
import { createFileRoute, Outlet } from '@tanstack/react-router'

export const Route = createFileRoute('/materials')({
  component: RouteComponent,
})

function RouteComponent() {
  return (
    <div className="h-svh w-full p-2 bg-background">
      <div className="flex h-full w-full rounded-2xl overflow-hidden border border-accent">
        <div className="w-64 z-10 overflow-y-auto border-r border-r-accent shrink-0">
          <AppSidebar />
        </div>

        {/* Main Content */}
        <div className="flex-1 h-full overflow-auto bg-background z-30">
          <Outlet />
        </div>
      </div>
    </div>
  )
}
