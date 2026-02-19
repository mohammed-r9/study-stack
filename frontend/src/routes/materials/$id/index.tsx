import { Button } from '@/components/ui/button'
import { useInfiniteLectures } from '@/lib/queries/lectures'
import { createFileRoute, Link } from '@tanstack/react-router'
import { Plus } from 'lucide-react'
import LectureCard from './-components/lecture-card'
import { Spinner } from '@/components/ui/spinner'
import { useInView } from 'react-intersection-observer'
import { useEffect } from 'react'

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
      <Button asChild>
        <Link
          to="/materials/$id/add"
          params={{ id: id }}
          search={{ title: title }}
        >
          Add a lecture <Plus />{' '}
        </Link>
      </Button>

      <div className="grid-cols-4 grid gap-4 my-3">
        {allLectures.map((lecture) => (
          <LectureCard key={lecture.id} lecture={lecture} />
        ))}
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
