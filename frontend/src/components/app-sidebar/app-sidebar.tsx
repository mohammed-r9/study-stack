import { useState } from 'react'
import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupLabel,
  SidebarGroupContent,
  SidebarMenu,
  SidebarHeader,
} from '@/components/ui/sidebar'
import { Separator } from '../ui/separator'
import { useCollections } from '@/lib/queries/library'
import type { Collection } from '@/lib/api/types'
import CollectionItem from './collection-item'
import { AddCollectionDialog } from './dialogs/add-collection'
import { BookOpen } from 'lucide-react'

export default function AppSidebar() {
  const { data: collections } = useCollections()
  const [openCollections, setOpenCollections] = useState<
    Record<string, boolean>
  >({})

  const toggleCollection = (id: string) => {
    setOpenCollections((prev) => ({ ...prev, [id]: !prev[id] }))
  }

  return (
    <Sidebar collapsible="none" className="w-full h-svh">
      <SidebarHeader className="font-bold text-2xl flex flex-row items-center gap-4 justify-between py-4">
        <div className="flex items-center justify-center gap-2">
          <BookOpen className="size-7 text-primary" />
          <p>Study Stack</p>
        </div>
        <AddCollectionDialog />
      </SidebarHeader>
      <Separator orientation="horizontal" />
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>Collections</SidebarGroupLabel>
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
  )
}
