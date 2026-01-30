import AppSidebar from '@/components/app-sidebar/app-sidebar'
import { createFileRoute, Outlet } from '@tanstack/react-router'

export const Route = createFileRoute('/dashboard')({
  component: RouteComponent,
})

function RouteComponent() {
  return (
    <div className="flex h-svh w-full bg-background my-2 gap-2">
      <div className="w-64 z-10 rounded-r-2xl overflow-clip">
        <AppSidebar />
      </div>

      <div className="h-full w-full overflow-auto p-2 rounded-l-2xl bg-accent z-30 flex flex-1">
        <Outlet />
      </div>
    </div>
  )
}
