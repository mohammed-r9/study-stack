import ModeToggle from '../mode-toggle'
import { Card } from '../ui/card'
import { UserButton } from './user-button'

export const AppHeader = () => {
  return (
    <Card className="sticky top-0 z-50 rounded-none p-2 mb-0 w-full shadow-none flex justify-between flex-row px-10 border-t-0">
      <div>
        <ModeToggle />
      </div>
      <div>
        <UserButton />
      </div>
    </Card>
  )
}
