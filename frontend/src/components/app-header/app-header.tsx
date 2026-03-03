import { Link, type LinkProps } from '@tanstack/react-router'
import { SignedIn } from '../auth/auth'
import ModeToggle from '../mode-toggle'
import { Card } from '../ui/card'
import { UserButton } from './user-button'
import { forwardRef } from 'react'
import { Button } from '../ui/button'

type NavLinkProps = LinkProps
export const NavLink = forwardRef<HTMLAnchorElement, NavLinkProps>(
  (props, ref) => {
    return (
      <Button asChild variant={'ghost'}>
        <Link {...props} ref={ref} />
      </Button>
    )
  },
)
export const AppHeader = () => {
  return (
    <Card className="sticky top-0 z-50 rounded-none p-2 mb-0 w-full shadow-none flex justify-between flex-row px-10 border-t-0">
      <div>
        <ModeToggle />
      </div>
      <div className="flex gap-4 items-center">
        <NavLink to="/materials">Materials</NavLink>
        <NavLink to="/flash-cards">Flash cards</NavLink>
      </div>
      <div>
        <SignedIn>
          <UserButton />
        </SignedIn>
      </div>
    </Card>
  )
}
