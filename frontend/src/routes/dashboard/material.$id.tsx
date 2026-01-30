import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/dashboard/material/$id')({
  component: RouteComponent,
})

function RouteComponent() {
  const params = Route.useParams()
  return <div>{params.id}</div>
}
