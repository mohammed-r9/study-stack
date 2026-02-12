import { Separator } from '@/components/ui/separator'
import { createFileRoute, Outlet } from '@tanstack/react-router'

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
  return (
    <div>
      <p className="font-bold text-2xl p-2">
        {title || 'Invalid material name'}
      </p>
      <Separator orientation="horizontal" />
      <div className="p-4">
        {' '}
        <Outlet />
      </div>
    </div>
  )
}
