import { Separator } from '@/components/ui/separator'
import { createFileRoute } from '@tanstack/react-router'

type MaterialSerach = {
  title: string
}
export const Route = createFileRoute('/dashboard/material/$id')({
  component: RouteComponent,
  validateSearch: (search: Record<string, string>): MaterialSerach => {
    return {
      title: search?.title ?? '',
    }
  },
})

function RouteComponent() {
  const { id } = Route.useParams()
  // title is temporary, it should be gotten from the backend
  // it works for now tho
  const { title } = Route.useSearch()
  return (
    <div>
      <p className="font-bold text-2xl p-2">
        {title || 'Invalid material name'}
      </p>
      <Separator orientation="horizontal" />
      <div className="p-2">
        <p>Material id: {id}</p>
      </div>
    </div>
  )
}
