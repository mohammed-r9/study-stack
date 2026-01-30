import { useState } from 'react'
import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupLabel,
  SidebarGroupContent,
  SidebarMenu,
  SidebarMenuItem,
  SidebarMenuButton,
  SidebarMenuSub,
  SidebarMenuSubItem,
  SidebarMenuSubButton,
  SidebarProvider,
} from '@/components/ui/sidebar'
import {
  Award,
  Book,
  BookOpen,
  Clipboard,
  Folder,
  FolderOpen,
} from 'lucide-react'
import { Link } from '@tanstack/react-router'

const collections = [
  {
    id: 'collection-1',
    name: 'Collection 1',
    materials: [
      { id: 'material-1', name: 'Material 1' },
      { id: 'material-2', name: 'Material 2' },
      { id: 'material-3', name: 'Material 3' },
    ],
  },
  {
    id: 'collection-2',
    name: 'Collection 2',
    materials: [
      { id: 'material-4', name: 'Material 4' },
      { id: 'material-5', name: 'Material 5' },
    ],
  },
]

export default function AppSidebar() {
  const [openCollections, setOpenCollections] = useState<
    Record<string, boolean>
  >({})

  const toggleCollection = (id: string) => {
    setOpenCollections((prev) => ({ ...prev, [id]: !prev[id] }))
  }

  return (
    <SidebarProvider>
      <Sidebar collapsible="none" className="h-screen">
        <SidebarContent>
          <SidebarGroup>
            <SidebarGroupLabel>Collections</SidebarGroupLabel>

            <SidebarGroupContent>
              <SidebarMenu>
                {collections.map((collection) => (
                  <SidebarMenuItem key={collection.id}>
                    {/* Collection button */}
                    <SidebarMenuButton
                      className="hover:cursor-pointer"
                      onClick={() => toggleCollection(collection.id)}
                    >
                      {openCollections[collection.id] ? (
                        <FolderOpen className="mr-2" />
                      ) : (
                        <Folder className="mr-2" />
                      )}
                      {collection.name}
                    </SidebarMenuButton>

                    {/* Submenu, collapsible */}
                    {openCollections[collection.id] && (
                      <SidebarMenuSub>
                        {collection.materials.map((material) => (
                          <SidebarMenuSubItem key={material.id}>
                            <SidebarMenuSubButton asChild>
                              <Link
                                to={`/dashboard/material/${material.id}`}
                                activeOptions={{ exact: true }}
                                activeProps={{
                                  className:
                                    'bg-primary/20 border-2 border-primary/20 text-accent-foreground',
                                }}
                              >
                                <BookOpen className="mr-2" />
                                {material.name}
                              </Link>
                            </SidebarMenuSubButton>
                          </SidebarMenuSubItem>
                        ))}
                      </SidebarMenuSub>
                    )}
                  </SidebarMenuItem>
                ))}
              </SidebarMenu>
            </SidebarGroupContent>
          </SidebarGroup>
        </SidebarContent>
      </Sidebar>
    </SidebarProvider>
  )
}
