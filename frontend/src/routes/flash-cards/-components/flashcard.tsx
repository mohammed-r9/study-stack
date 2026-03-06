import { motion, AnimatePresence } from 'motion/react'
import { useState } from 'react'

import { Button } from '@/components/ui/button'
import {
  Card,
  CardHeader,
  CardTitle,
  CardContent,
  CardFooter,
} from '@/components/ui/card'

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
          <Card className="p-4 text-center">
            <CardHeader>
              <CardTitle>{flip ? 'Back' : 'Front'}</CardTitle>
            </CardHeader>

            <CardContent>
              {flip
                ? back
                  ? back
                  : 'Loading...'
                : front
                  ? front
                  : 'Loading...'}
            </CardContent>

            <CardFooter className="flex justify-end">
              <Button onClick={() => setFlip((prev) => !prev)}>Flip</Button>
            </CardFooter>
          </Card>
        </motion.div>
      </AnimatePresence>
    </div>
  )
}
