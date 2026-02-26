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
import { Tooltip, TooltipContent, TooltipTrigger } from '../ui/tooltip'
import { ScrollArea } from '../ui/scroll-area'

export default function AppSidebar() {
  const { data: collections } = useCollections()
  const [openCollections, setOpenCollections] = useState<
    Record<string, boolean>
  >({})

  const toggleCollection = (id: string) => {
    setOpenCollections((prev) => ({ ...prev, [id]: !prev[id] }))
  }

  return (
    <Sidebar collapsible="none" className="w-full h-screen">
      <SidebarHeader className="font-bold text-2xl flex flex-row items-center gap-4 justify-between py-4">
        <div className="flex items-center justify-center gap-2 py-1">
          <BookOpen className="size-7 text-primary" />
          <p>Study Stack</p>
        </div>
        <Tooltip>
          <TooltipTrigger asChild>
            <span className="inline-flex">
              {' '}
              <AddCollectionDialog />
            </span>
          </TooltipTrigger>
          <TooltipContent side="bottom">
            <p>New collection</p>
          </TooltipContent>
        </Tooltip>
      </SidebarHeader>
      <Separator orientation="horizontal" />

      <SidebarContent className="flex-1 relative pb-10">
        <SidebarGroup>
          <SidebarGroupLabel>Collections</SidebarGroupLabel>
          <SidebarGroupContent>
            <ScrollArea className="overflow-y-auto h-[60svh] mask-[linear-gradient(to_bottom,transparent,black_10%,black_90%,transparent)]">
              <SidebarMenu>
                {collections?.data?.map((collection: Collection) => (
                  <CollectionItem
                    key={collection.id}
                    collection={collection}
                    isOpen={openCollections[collection.id]}
                    toggleCollection={toggleCollection}
                    className="first:mt-7 last:mb-7"
                  />
                ))}
              </SidebarMenu>
            </ScrollArea>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
    </Sidebar>
  )
}
