import { useState } from 'react'
import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupLabel,
  SidebarGroupContent,
  SidebarMenu,
  SidebarProvider,
  SidebarHeader,
} from '@/components/ui/sidebar'
import { Separator } from '../ui/separator'
import { useCollections } from '@/lib/queries/library'
import type { Collection } from '@/lib/api/types'
import CollectionItem from './collection-item'
import { AddCollectionDialog } from './dialogs/add-collection'

export default function AppSidebar() {
  const { data: collections } = useCollections()
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
            <AddCollectionDialog />
            <SidebarGroupContent>
              <SidebarMenu>
                {collections?.data?.map?.((collection: Collection) => (
                  <CollectionItem
                    key={collection.id}
                    collection={collection}
                    isOpen={openCollections[collection.id]}
                    toggleCollection={toggleCollection}
                  />
                ))}
              </SidebarMenu>
            </SidebarGroupContent>
          </SidebarGroup>
        </SidebarContent>
      </Sidebar>
    </SidebarProvider>
  )
}
