import { motion, AnimatePresence } from 'motion/react'
import { useState } from 'react'

import { Button } from '@/components/ui/button'
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card'
import { cn } from '@/lib/utils'
import { Spinner } from '@/components/ui/spinner'

type Props = {
  front: string | undefined
  back: string | undefined
}

export default function Flashcard({ front, back }: Props) {
  const [flip, setFlip] = useState(false)

  return (
    <div className="w-xl">
      <AnimatePresence mode="wait">
        <motion.div
          key={flip ? 'front' : 'back'}
          initial={{ rotateY: -90 }}
          animate={{ rotateY: 0 }}
          exit={{ rotateY: 90 }}
          transition={{ duration: 0.2 }}
        >
          <Card
            className={cn(
              'p-4 text-center',
              flip ? 'dark:bg-card/30 bg-border/25' : '',
            )}
          >
            <CardHeader>
              <CardTitle className="text-lg">
                {flip ? 'Back' : 'Front'}
              </CardTitle>
            </CardHeader>

            <CardContent className="text-lg">
              {flip ? (
                back ? (
                  back
                ) : (
                  <Spinner className="text-primary size-6!" />
                )
              ) : front ? (
                front
              ) : (
                <Spinner className="text-primary size-6!" />
              )}
            </CardContent>

            <div className="flex justify-end">
              <Button onClick={() => setFlip((prev) => !prev)}>Flip</Button>
            </div>
          </Card>
        </motion.div>
      </AnimatePresence>
    </div>
  )
}
