import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/flash-cards/add/$materialId')({
  component: RouteComponent,
})

function RouteComponent() {
  const { materialId } = Route.useParams()
  return <div>Hello flash-cards/add | material id: {materialId}</div>
}
