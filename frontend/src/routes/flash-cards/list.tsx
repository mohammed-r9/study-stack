import { createFileRoute } from '@tanstack/react-router'
import FlashcardListItem from './-components/flashcard-list-item'
import { Button } from '@/components/ui/button'
import { LayoutGrid, List, Plus } from 'lucide-react'
import { cn } from '@/lib/utils'
import { ScrollArea } from '@/components/ui/scroll-area'
import { useInfiniteFlashcards } from '@/lib/queries/flashcards'
import { useEffect } from 'react'
import { useInView } from 'react-intersection-observer'
import { ButtonGroup } from '@/components/ui/button-group'
import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from '@/components/ui/tooltip'

type Search = {
  view: 'grid' | 'list'
}
export const Route = createFileRoute('/flash-cards/list')({
  component: RouteComponent,
  validateSearch: (search: Search): Search => {
    return {
      view: search?.view ?? 'grid',
    }
  },
})

export default function RouteComponent() {
  const { view } = Route.useSearch()

  const navigate = Route.useNavigate()

  const setViewMode = (newMode: 'grid' | 'list') => {
    navigate({
      search: (prev) => ({ ...prev, view: newMode }),
    })
  }
  const { data, fetchNextPage, hasNextPage, isFetchingNextPage } =
    useInfiniteFlashcards()

  const { ref, inView } = useInView()

  useEffect(() => {
    if (inView && hasNextPage && !isFetchingNextPage) {
      fetchNextPage()
    }
  }, [inView, hasNextPage, isFetchingNextPage, fetchNextPage])

  const allFlashcards = data?.pages.flatMap((page) => page.flashcards) ?? []

  return (
    <ScrollArea className="h-screen bg-background p-6 lg:p-10 overflow-scroll">
      <div className="mx-auto max-w-6xl space-y-8">
        <div className="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
          <div>
            <h1 className="text-3xl font-bold tracking-tight">Flashcards</h1>
            <p className="text-muted-foreground text-sm">
              Review and manage your study set.
            </p>
          </div>

          <div className="flex items-center gap-2">
            <ButtonGroup className="flex gap-1">
              <Tooltip>
                <TooltipTrigger>
                  <Button
                    variant={view === 'list' ? 'secondary' : 'ghost'}
                    size="sm"
                    className="h-8 w-8 p-0"
                    onClick={() => setViewMode('list')}
                  >
                    <List className="h-4 w-4" />
                  </Button>
                </TooltipTrigger>
                <TooltipContent>List view</TooltipContent>
              </Tooltip>

              <Tooltip>
                <TooltipTrigger>
                  <Button
                    variant={view === 'grid' ? 'secondary' : 'ghost'}
                    size="sm"
                    className="h-8 w-8 p-0"
                    onClick={() => setViewMode('grid')}
                  >
                    <LayoutGrid className="h-4 w-4" />
                  </Button>
                </TooltipTrigger>
                <TooltipContent>Grid view</TooltipContent>
              </Tooltip>
            </ButtonGroup>

            <Button
              className="gap-2"
              size={'lg'}
              onClick={() => {
                navigate({ to: '/materials', search: { alert: true } })
              }}
            >
              <Plus className="h-4 w-4" />
              New Card
            </Button>
          </div>
        </div>

        <div
          className={cn(
            'grid gap-6 mb-8',
            view === 'grid'
              ? 'grid-cols-1 md:grid-cols-2 xl:grid-cols-3'
              : 'grid-cols-1 max-w-4xl mx-auto',
          )}
        >
          {allFlashcards.map((card) => (
            <FlashcardListItem key={card.id} {...card} />
          ))}
        </div>

        <div ref={ref}></div>
      </div>
    </ScrollArea>
  )
}
