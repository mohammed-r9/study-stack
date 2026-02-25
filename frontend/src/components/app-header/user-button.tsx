'use client'

import { useQuery } from '@tanstack/react-query'
import { createUserQueryOptions } from '@/lib/queries/user'

import {
  DropdownMenu,
  DropdownMenuTrigger,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
} from '@/components/ui/dropdown-menu'

import { Button } from '@/components/ui/button'
import { Avatar, AvatarFallback } from '@/components/ui/avatar'

import { LogOut, Settings, User as UserIcon } from 'lucide-react'
import { Spinner } from '../ui/spinner'
import { Link } from '@tanstack/react-router'

export function UserButton() {
  const { data: user, isLoading, isError } = useQuery(createUserQueryOptions())

  if (isLoading) {
    return (
      <Button variant="ghost" size="icon" disabled>
        <Avatar className="h-8 w-8">
          <AvatarFallback>
            <Spinner />
          </AvatarFallback>
        </Avatar>
      </Button>
    )
  }

  if (isError || !user?.data) {
    return null
  }

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild className="rounded-full h-9 w-9">
        <Button variant="ghost" className="flex items-center gap-2 p-2">
          <Avatar className="h-8 w-8">
            <AvatarFallback>
              {user?.data?.name?.charAt(0).toUpperCase()}
            </AvatarFallback>
          </Avatar>
        </Button>
      </DropdownMenuTrigger>

      <DropdownMenuContent align="end" className="w-56">
        <div className="px-2 py-1.5 text-sm">
          <p className="font-medium">{user.data.name}</p>
          <p className="text-muted-foreground text-xs">{user.data.email}</p>
        </div>

        <DropdownMenuSeparator />

        <DropdownMenuItem>
          <UserIcon className="mr-2 h-4 w-4" />
          <Link to="/profile" className="w-full">
            Profile
          </Link>
        </DropdownMenuItem>

        <DropdownMenuItem>
          <Settings className="mr-2 h-4 w-4" />
          <Link to="/settings" className="w-full">
            Settings
          </Link>
        </DropdownMenuItem>

        <DropdownMenuSeparator />

        <DropdownMenuItem className="hover:cursor-pointer">
          <LogOut className="mr-2 h-4 w-4" />
          Logout
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}
