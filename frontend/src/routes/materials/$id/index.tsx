import { Button } from '@/components/ui/button'
import { useInfiniteLectures } from '@/lib/queries/lectures'
import { createFileRoute, Link } from '@tanstack/react-router'
import { AlertTriangle, Plus } from 'lucide-react'
import LectureCard from './-components/lecture-card'
import { Spinner } from '@/components/ui/spinner'
import { useInView } from 'react-intersection-observer'
import { useEffect } from 'react'
import {
  Empty,
  EmptyContent,
  EmptyDescription,
  EmptyHeader,
  EmptyMedia,
  EmptyTitle,
} from '@/components/ui/empty'

export const Route = createFileRoute('/materials/$id/')({
  component: RouteComponent,
})

function RouteComponent() {
  const { title } = Route.useSearch()
  const { id } = Route.useParams()
  const { data, fetchNextPage, hasNextPage, isFetchingNextPage, status } =
    useInfiniteLectures(id)
  const { ref, inView } = useInView()

  useEffect(() => {
    if (inView && hasNextPage && !isFetchingNextPage) fetchNextPage()
  }, [inView, fetchNextPage()])

  const allLectures = data?.pages.flatMap((page) => page.lectures) ?? []

  if (status === 'pending') {
    return <LecturesLoader />
  }
  return (
    <div>
      <div className="grid-cols-4 grid gap-4 my-3">
        {allLectures.map((lecture) =>
          lecture ? (
            <LectureCard key={lecture.id} lecture={lecture} />
          ) : (
            <Empty className="border col-span-4">
              <EmptyHeader>
                <EmptyMedia variant={'icon'}>
                  {' '}
                  <AlertTriangle />{' '}
                </EmptyMedia>
                <EmptyTitle>No lectures found</EmptyTitle>
                <EmptyDescription>
                  Please add some lectures first.
                </EmptyDescription>
                <EmptyContent>
                  <Button asChild variant={'secondary'}>
                    <Link
                      to="/materials/$id/add"
                      params={{ id }}
                      search={{ title }}
                      className="flex items-center gap-2"
                    >
                      Add a lecture
                      <Plus className="w-4 h-4" />
                    </Link>
                  </Button>
                </EmptyContent>
              </EmptyHeader>
            </Empty>
          ),
        )}
      </div>

      {isFetchingNextPage && <LecturesLoader />}
      <div ref={ref}></div>
    </div>
  )
}

function LecturesLoader() {
  return (
    <div className="flex flex-col items-center justify-center min-h-[50vh] w-full">
      <Spinner className="size-10 text-primary" />
      <p className="mt-4 text-sm text-muted-foreground animate-pulse">
        Loading lectures...
      </p>
    </div>
  )
}
