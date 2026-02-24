import { useMaterials } from '@/lib/queries/library'
import {
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarMenuSub,
  SidebarMenuSubButton,
  SidebarMenuSubItem,
} from '../ui/sidebar'
import { BookOpen, Folder, FolderOpen } from 'lucide-react'
import UpdateCollectionDialog from './dialogs/update-collection'
import type { Collection, Material } from '@/lib/api/types'
import { Link } from '@tanstack/react-router'
import { AddMaterialDialog } from './dialogs/add-material'
import { useState } from 'react'
import { cn } from '@/lib/utils'

export default function CollectionItem({
  collection,
  isOpen,
  toggleCollection,
}: {
  collection: Collection
  isOpen: boolean
  toggleCollection: (id: string) => void
}) {
  const materialsQuery = useMaterials(collection.id, { enabled: isOpen })

  const [isHoveredOn, setIsHoveredOn] = useState(false)
  const materialsData = materialsQuery.data

  return (
    <>
      <SidebarMenuItem
        onMouseEnter={() => setIsHoveredOn(true)}
        onMouseLeave={() => setIsHoveredOn(false)}
      >
        <div key={collection.id} className="flex gap-2">
          <SidebarMenuButton
            isActive={isOpen}
            className="hover:cursor-pointer flex items-center justify-between h-9 rounded-sm"
            onClick={() => toggleCollection(collection.id)}
          >
            <div className="flex items-center gap-2">
              {isOpen ? (
                <FolderOpen className="size-4" />
              ) : (
                <Folder className="size-4" />
              )}
              <span className="truncate">{collection.title}</span>
            </div>
          </SidebarMenuButton>

          <div className={cn(isHoveredOn ? 'opacity-100' : 'opacity-0')}>
            <UpdateCollectionDialog
              collectionID={collection.id}
              setIsHoveredOn={setIsHoveredOn}
            />
          </div>
        </div>
        {isOpen && (
          <SidebarMenuSub>
            <AddMaterialDialog collectionID={collection.id} />
            {materialsData?.data?.map?.((material: Material) => (
              <SidebarMenuSubItem key={material.id}>
                <SidebarMenuSubButton asChild>
                  <Link
                    to="/materials/$id"
                    params={{ id: material.id }}
                    search={{ title: material.title }}
                    activeOptions={{ exact: true }}
                    className="w-72 rounded-none"
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
    </>
  )
}
