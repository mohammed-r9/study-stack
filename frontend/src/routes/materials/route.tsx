import AppSidebar from '@/components/app-sidebar/app-sidebar'
import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert'
import { Button } from '@/components/ui/button'
import { ScrollArea } from '@/components/ui/scroll-area'
import { SidebarProvider } from '@/components/ui/sidebar'
import { createFileRoute, Outlet, useNavigate } from '@tanstack/react-router'
import { Info, X } from 'lucide-react'

type Search = {
  alert?: boolean
}

export const Route = createFileRoute('/materials')({
  component: RouteComponent,
  validateSearch: (search: Search): Search => {
    return {
      alert: search?.alert ?? false,
    }
  },
})

function RouteComponent() {
  const { alert } = Route.useSearch()
  const navigate = useNavigate({ from: Route.fullPath })

  const handleClose = () => {
    navigate({
      search: (prev) => {
        const { alert: showAlert, ...rest } = prev
        return rest
      },
    })
  }

  return (
    <div className="relative h-screen w-full bg-background overflow-hidden">
      {alert && (
        <div className="absolute top-6 right-6 z-50 w-full max-w-sm animate-in fade-in slide-in-from-bottom-4 duration-300">
          <Alert className="border-border bg-card shadow-lg ring-1 ring-black/5 dark:ring-white/10">
            <div className="flex w-full items-start justify-between gap-2">
              <div className="grid gap-2">
                <AlertTitle className="font-semibold leading-none tracking-tight text-foreground flex items-center gap-3">
                  <Info className="h-4 w-4" />
                  Action Required
                </AlertTitle>
                <AlertDescription className="text-sm text-muted-foreground leading-relaxed">
                  Please select a material first, then you'll be able to add
                  your flashcard.
                </AlertDescription>
              </div>
              <Button
                variant="ghost"
                size="icon"
                className="-mt-1 -mr-2 h-7 w-7 text-muted-foreground hover:text-foreground"
                onClick={handleClose}
              >
                <X className="h-4 w-4" />
                <span className="sr-only">Dismiss</span>
              </Button>
            </div>
          </Alert>
        </div>
      )}

      <div className="flex h-full w-full border border-accent border-t-0">
        <div className="w-72 shrink-0 border-r border-accent">
          <SidebarProvider>
            <AppSidebar />
          </SidebarProvider>
        </div>

        <ScrollArea className="flex-1 h-full bg-background">
          <Outlet />
        </ScrollArea>
      </div>
    </div>
  )
}
