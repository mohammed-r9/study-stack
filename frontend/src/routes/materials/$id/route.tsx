import { Button } from '@/components/ui/button'
import { Separator } from '@/components/ui/separator'
import {
  createFileRoute,
  Link,
  Outlet,
  useLocation,
  useMatchRoute,
} from '@tanstack/react-router'
import { Plus } from 'lucide-react'

export type MaterialSerach = {
  title: string
}

export type MaterialParams = {
  id: string
}

export const Route = createFileRoute('/materials/$id')({
  component: RouteComponent,
  validateSearch: (search: Record<string, string>): MaterialSerach => {
    return {
      title: search?.title ?? '',
    }
  },
})

function RouteComponent() {
  const { title } = Route.useSearch()
  const { id } = Route.useParams()
  const location = useLocation()
  const isAddRoute = location.pathname.endsWith('/add')
  return (
    <>
      <div className="sticky top-0 z-20 bg-background">
        <div className="flex items-center justify-between px-4 py-4.5">
          <p className="font-bold text-2xl">
            {title || 'Invalid material name'}
          </p>

          {!isAddRoute && (
            <Button asChild>
              <Link
                to="/materials/$id/add"
                params={{ id }}
                search={{ title }}
                className="flex items-center gap-2"
              >
                Add a lecture
                <Plus className="w-4 h-4" />
              </Link>
            </Button>
          )}
        </div>

        <Separator orientation="horizontal" />
      </div>

      <div className="p-4">
        <Outlet />
      </div>
    </>
  )
}
