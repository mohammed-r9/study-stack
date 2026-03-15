import { Button } from '@/components/ui/button'
import { cn } from '@/lib/utils'
import { createFileRoute } from '@tanstack/react-router'
import {
  ArrowBigLeftDash,
  ArrowBigRightDash,
  Image,
  Shapes,
  type LucideIcon,
} from 'lucide-react'
import { useState } from 'react'
import AddImage from './-components/add-image'
import Labels from './-components/labels'

export const Route = createFileRoute('/image-flashcard/add')({
  component: RouteComponent,
})

type StepProps = {
  imageFile: File | null
  setImageFile: React.Dispatch<React.SetStateAction<File | null>>
}

type Step = {
  title: string
  desc: string
  icon: LucideIcon
  component: React.ComponentType<StepProps>
}

const STEPS: Step[] = [
  {
    title: 'Select Image',
    desc: 'Upload the image',
    icon: Image,
    component: AddImage,
  },
  {
    title: 'Add Labels',
    desc: 'Add labels on top of the image',
    icon: Shapes,
    component: Labels,
  },
]

function RouteComponent() {
  const [current, setCurrent] = useState(0)
  const [imageFile, setImageFile] = useState<File | null>(null)
  const StepComponent = STEPS[current].component
  return (
    <div className="flex h-[90dvh] overflow-scroll">
      {/* Sidebar */}
      <aside className="w-56 border-r bg-background/30 flex flex-col gap-2 p-4">
        {STEPS.map((s, i) => {
          const Icon = s.icon
          const isActive = current === i
          return (
            <p
              key={i}
              className={cn(
                'flex items-center gap-3 px-3 py-2 rounded-md w-full text-left hover:cursor-default',
                isActive
                  ? 'bg-primary text-primary-foreground'
                  : 'text-muted-foreground',
              )}
            >
              <Icon
                className={cn(
                  'h-5 w-5 shrink-0',
                  isActive
                    ? 'text-primary-foreground'
                    : 'text-muted-foreground',
                )}
              />
              <div className="flex flex-col">
                <span className="text-sm font-medium">{s.title}</span>
                <span
                  className={cn(
                    'text-xs',
                    !isActive ? 'text-muted-foreground' : '',
                  )}
                >
                  {s.desc}
                </span>
              </div>
            </p>
          )
        })}
      </aside>

      <div className="flex-1 flex flex-col p-3 overflow-auto w-full outline">
        <div className="flex gap-3 justify-end w-full mb-auto">
          <Button
            variant="outline"
            size="lg"
            disabled={current === 0}
            onClick={() => setCurrent((c) => c - 1)}
          >
            <ArrowBigLeftDash />
            Back
          </Button>
          <Button
            size="lg"
            disabled={current === STEPS.length - 1 || imageFile === null}
            onClick={() => setCurrent((c) => c + 1)}
          >
            Continue
            <ArrowBigRightDash />
          </Button>
        </div>

        <div className="flex-1 flex items-center justify-center">
          {<StepComponent setImageFile={setImageFile} imageFile={imageFile} />}
        </div>
      </div>
    </div>
  )
}
