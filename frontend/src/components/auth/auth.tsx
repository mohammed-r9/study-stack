import { useAuthStore } from '@/lib/store/auth'
import type { ReactNode } from 'react'

export function SignedIn({ children }: { children: ReactNode }) {
  const accessToken = useAuthStore((state) => state.accessToken)
  const isAuthenticated = !!accessToken

  if (!isAuthenticated) return null

  return <>{children}</>
}

export function SignedOut({ children }: { children: ReactNode }) {
  const accessToken = useAuthStore((state) => state.accessToken)
  const isAuthenticated = !!accessToken

  if (isAuthenticated) return null

  return <>{children}</>
}
