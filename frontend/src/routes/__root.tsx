import {
  Outlet,
  createRootRouteWithContext,
  redirect,
} from '@tanstack/react-router'
import { TanStackRouterDevtoolsPanel } from '@tanstack/react-router-devtools'
import { TanStackDevtools } from '@tanstack/react-devtools'

import TanStackQueryDevtools from '../integrations/tanstack-query/devtools'

import type { QueryClient } from '@tanstack/react-query'
import type { AuthContext } from '@/lib/context/auth'
import { authLoader } from '@/lib/auth-loader'
import { Toaster } from 'sonner'
import { useAuthStore } from '@/lib/store/auth'
import { AppHeader } from '@/components/app-header/app-header'
import { useEffect } from 'react'
import { useThemeStore } from '@/lib/store/theme'
import { TooltipProvider } from '@/components/ui/tooltip'

export interface MyRouterContext {
  queryClient: QueryClient
  auth: AuthContext
}

export const Route = createRootRouteWithContext<MyRouterContext>()({
  component: () => {
    const theme = useThemeStore((state) => state.theme)

    useEffect(() => {
      const root = window.document.documentElement
      root.classList.remove('light', 'dark')
      root.classList.add(theme)
    }, [theme])

    return (
      <TooltipProvider>
        <div className="h-screen flex flex-col overflow-hidden">
          <AppHeader />

          <div className="flex-1 overflow-hidden">
            <Outlet />
          </div>
        </div>

        <TanStackDevtools
          config={{ position: 'bottom-right' }}
          plugins={[
            {
              name: 'Tanstack Router',
              render: <TanStackRouterDevtoolsPanel />,
            },
            TanStackQueryDevtools,
          ]}
        />

        <Toaster richColors theme={theme} position="top-center" />
      </TooltipProvider>
    )
  },
  loader: async ({ context: ctx }) => {
    await authLoader(ctx)

    const path = window.location.pathname
    const isAuth = ctx.auth.isAuthenticated
    console.log(useAuthStore.getState().accessToken)

    if (
      !isAuth &&
      !path.startsWith('/login') &&
      !path.startsWith('/register')
    ) {
      throw redirect({ to: '/login', replace: true })
    }

    if (isAuth && (path.startsWith('/login') || path.startsWith('/register'))) {
      throw redirect({
        to: '/materials',
        replace: true,
      })
    }
  },
})
