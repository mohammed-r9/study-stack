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
import { ThemeProvider } from '@/components/theme-provider'
import { Toaster } from 'sonner'
import { useAuthStore } from '@/lib/store/auth'

export interface MyRouterContext {
  queryClient: QueryClient
  auth: AuthContext
}

export const Route = createRootRouteWithContext<MyRouterContext>()({
  component: () => (
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
      <Outlet />

      <TanStackDevtools
        config={{
          position: 'bottom-right',
        }}
        plugins={[
          {
            name: 'Tanstack Router',
            render: <TanStackRouterDevtoolsPanel />,
          },
          TanStackQueryDevtools,
        ]}
      />
      <Toaster />
    </ThemeProvider>
  ),
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
