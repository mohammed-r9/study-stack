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

export default function CollectionItem({
  collection,
  isOpen,
  toggleCollection,
}: {
  collection: Collection
  isOpen: boolean
  toggleCollection: (id: string) => void
}) {
  const materialsQuery = isOpen
    ? useMaterials(collection.id)
    : { data: { data: [] } }

  const materialsData = materialsQuery.data

  return (
    <SidebarMenuItem key={collection.id}>
      <SidebarMenuButton
        className="hover:cursor-pointer flex items-center justify-between"
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

        <UpdateCollectionDialog collectionID={collection.id} />
      </SidebarMenuButton>

      {isOpen && (
        <SidebarMenuSub>
          <AddMaterialDialog collectionID={collection.id} />
          {materialsData?.data?.map?.((material: Material) => (
            <SidebarMenuSubItem key={material.id}>
              <SidebarMenuSubButton asChild>
                <Link
                  to="/materials/$id/"
                  params={{ id: material.id }}
                  search={{ title: material.title }}
                  activeOptions={{ exact: true }}
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
  )
}
