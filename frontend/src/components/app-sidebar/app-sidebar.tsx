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
import { useLibrary } from '@/lib/queries/user'
import UpdateCollectionDialog from './dialogs/update-collection'

export default function AppSidebar() {
  const { data: library, isError, isLoading } = useLibrary()
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
                {library?.data?.map?.((collection) => (
                  <SidebarMenuItem key={collection.id}>
                    {/* Collection button */}
                    <SidebarMenuButton
                      className="hover:cursor-pointer flex items-center justify-between"
                      onClick={() => toggleCollection(collection.id)}
                    >
                      <div className="flex items-center gap-2">
                        {openCollections[collection.id] ? (
                          <FolderOpen className="size-4" />
                        ) : (
                          <Folder className="size-4" />
                        )}
                        <span className="truncate">{collection.title}</span>
                      </div>

                      <UpdateCollectionDialog collectionID={collection.id} />
                    </SidebarMenuButton>

                    {/* Submenu, collapsible */}
                    {openCollections[collection.id] && (
                      <SidebarMenuSub>
                        {collection?.materials.map((material) => (
                          <SidebarMenuSubItem key={material.id}>
                            <SidebarMenuSubButton asChild>
                              <Link
                                to="/materials/$id"
                                params={{ id: material.id }}
                                search={{ title: material.title }}
                                activeOptions={{
                                  exact: true,
                                }}
                                activeProps={{
                                  className:
                                    'bg-primary/20 border-2 border-primary/20 text-accent-foreground hover:bg-primary/20! font-bold',
                                }}
                              >
                                <BookOpen className="mr-2" />
                                {material.title}
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
