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
      <Button asChild variant={'ghost'} size={'lg'}>
        <Link {...props} ref={ref} />
      </Button>
    )
  },
)
export const AppHeader = () => {
  return (
    <Card className="sticky top-0 z-50 rounded-none py-2 mb-0 w-full shadow-none flex justify-between items-center flex-row px-10 border-t-0">
      <div>
        <ModeToggle />
      </div>
      <SignedIn>
        <div className="flex gap-4 items-center">
          <NavLink to="/materials">Materials</NavLink>
          <NavLink to="/flash-cards/study">Flash cards</NavLink>
        </div>
      </SignedIn>
      <div>
        <SignedIn>
          <UserButton />
        </SignedIn>
      </div>
    </Card>
  )
}
