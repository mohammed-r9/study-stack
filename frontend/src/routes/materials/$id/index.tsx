import { Button } from '@/components/ui/button'
import { createFileRoute, Link } from '@tanstack/react-router'
import { Plus } from 'lucide-react'

export const Route = createFileRoute('/materials/$id/')({
  component: RouteComponent,
})

function RouteComponent() {
  const { title } = Route.useSearch()
  const { id } = Route.useParams()
  return (
    <div>
      <Button asChild>
        <Link
          to="/materials/$id/add"
          params={{ id: id }}
          search={{ title: title }}
        >
          Add a lecture <Plus />{' '}
        </Link>
      </Button>
    </div>
  )
}
