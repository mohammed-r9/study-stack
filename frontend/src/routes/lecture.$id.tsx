import { useSignedURL } from '@/lib/queries/lectures'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/lecture/$id')({
  component: RouteComponent,
})

function RouteComponent() {
  const { id } = Route.useParams()
  const { data } = useSignedURL(id)
  return <iframe src={data?.data.url} width="100%" height="600px"></iframe>
}
