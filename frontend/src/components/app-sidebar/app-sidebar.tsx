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
  SidebarHeader,
} from '@/components/ui/sidebar'
import { BookOpen, Folder, FolderOpen } from 'lucide-react'
import { Link } from '@tanstack/react-router'
import { Separator } from '../ui/separator'

const collections = [
  {
    id: 'collection-1',
    name: 'Collection 1',
    materials: [
      { id: '7d36b776-21fb-4a90-8896-fb4555228567', name: 'Material 1' },
      { id: 'dd666607-400d-4d75-bcac-fedeab413954', name: 'Material 2' },
      { id: 'f0ffef38-6453-411a-b3cf-46ea029646d4', name: 'Material 3' },
    ],
  },
  {
    id: 'collection-2',
    name: 'Collection 2',
    materials: [
      { id: 'f3ee6265-0c4c-473f-b655-609923682bc4', name: 'Material 4' },
      { id: 'f8c9c660-d981-4150-ad04-68562a9b81d4', name: 'Material 5' },
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
        <SidebarHeader className="font-bold text-2xl">
          Study Stack
        </SidebarHeader>
        <Separator orientation="horizontal" />
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
                                to="/dashboard/material/$id"
                                params={{ id: material.id }}
                                search={{ title: material.name }}
                                activeOptions={{
                                  exact: true,
                                }}
                                activeProps={{
                                  className:
                                    'bg-primary/20 border-2 border-primary/20 text-accent-foreground hover:bg-primary/20! font-bold',
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
