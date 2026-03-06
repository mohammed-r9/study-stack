import { createFileRoute, Link } from '@tanstack/react-router'
import Flashcard from './-components/flashcard'
import { useState } from 'react'
import { Button } from '@/components/ui/button'
import { useFlashcard } from '@/lib/queries/flashcards'
import { AlertTriangle, ArrowRight } from 'lucide-react'
import { AnimatePresence, motion } from 'motion/react'
import {
  Empty,
  EmptyContent,
  EmptyDescription,
  EmptyHeader,
  EmptyMedia,
  EmptyTitle,
} from '@/components/ui/empty'

export const Route = createFileRoute('/flash-cards')({
  component: RouteComponent,
})

function RouteComponent() {
  const [cursor, setCursor] = useState(0)
  const { data: flashcard, isLoading } = useFlashcard(cursor)

  if (!flashcard?.data && !isLoading)
    return (
      <Empty className="border col-span-4">
        <EmptyHeader>
          <EmptyMedia variant={'icon'}>
            {' '}
            <AlertTriangle />{' '}
          </EmptyMedia>
          <EmptyTitle>No flashcards found</EmptyTitle>
          <EmptyDescription>Please add some flashcards first.</EmptyDescription>
          <EmptyContent>
            <Button asChild variant={'secondary'}>
              <Link to="/materials">Go to materials</Link>
            </Button>
          </EmptyContent>
        </EmptyHeader>
      </Empty>
    )

  return (
    <div className="p-4 w-full min-h-screen flex flex-col items-center justify-center gap-6">
      <div className="w-full max-w-md flex justify-center">
        <AnimatePresence mode="wait">
          <motion.div
            key={cursor}
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            exit={{ opacity: 0, y: -20 }}
            transition={{ duration: 0.2 }}
          >
            <Flashcard
              front={flashcard?.data.front}
              back={flashcard?.data.back}
            />
          </motion.div>
        </AnimatePresence>
      </div>

      <div className="flex justify-center w-full">
        <Button
          variant="outline"
          onClick={() => setCursor((prev) => prev + 1)}
          disabled={isLoading}
        >
          <ArrowRight />
        </Button>
      </div>
    </div>
  )
}
