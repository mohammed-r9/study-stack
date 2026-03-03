import { createFileRoute, Outlet } from '@tanstack/react-router'

export const Route = createFileRoute('/flash-cards')({
  component: RouteComponent,
})

function RouteComponent() {
  return (
    <div>
      Hello "/falsh-cards"! <Outlet />{' '}
    </div>
  )
}
