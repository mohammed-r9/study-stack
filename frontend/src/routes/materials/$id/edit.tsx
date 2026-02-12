import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/materials/$id/edit')({
  component: RouteComponent,
})

function RouteComponent() {
  const params = Route.useParams()
  return <div>Hello "/materials/{params.id}/edit"!</div>
}
