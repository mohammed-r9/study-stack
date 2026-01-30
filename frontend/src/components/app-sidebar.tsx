import {
  Sidebar,
  SidebarContent,
  SidebarMenuItem,
  SidebarProvider,
} from '@/components/ui/sidebar'

export default function AppSidebar() {
  return (
    <SidebarProvider>
      <Sidebar>
        <SidebarContent>
          <SidebarMenuItem></SidebarMenuItem>
        </SidebarContent>
      </Sidebar>
    </SidebarProvider>
  )
}
