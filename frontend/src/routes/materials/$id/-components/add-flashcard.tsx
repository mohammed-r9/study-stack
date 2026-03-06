import AppField from '@/components/form/app-field'
import { Button } from '@/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import { useAppForm } from '@/hooks/form'
import { httpClient } from '@/lib/api'
import { flashcardSchema } from '@/lib/schemas/post'
import { Plus } from 'lucide-react'
import { useState } from 'react'
import { toast } from 'sonner'

type Props = {
  materialId: string
}

export default function AddFlashcard({ materialId }: Props) {
  const [open, setOpen] = useState(false)
  const form = useAppForm({
    defaultValues: {
      back: '',
      front: '',
    },
    validators: {
      onChange: flashcardSchema,
    },
    onSubmit: async ({ value }) => {
      try {
        httpClient.createFlashcard({
          front: value.front,
          back: value.back,
          material_id: materialId,
        })
        setOpen(false)
        form.reset()
        toast.success('Flashcard created')
      } catch (err) {
        toast.error(err as string)
      }
    },
  })
  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button variant={'secondary'}>
          Add a flashcard
          <Plus className="w-4 h-4" />
        </Button>
      </DialogTrigger>
      <DialogContent className="rounded-2xl">
        <DialogHeader>
          <DialogTitle className="text-xl">Add New Flashcard</DialogTitle>
        </DialogHeader>

        <form
          className="flex flex-col gap-5"
          onSubmit={(e) => {
            e.preventDefault()
            form.handleSubmit()
          }}
        >
          <form.AppField name="front">
            {() => <AppField type="text" id="front" label="Front" />}
          </form.AppField>
          <form.AppField name="back">
            {() => <AppField type="text" id="back" label="Back" />}
          </form.AppField>
          <Button type="submit">Create</Button>
        </form>
      </DialogContent>
    </Dialog>
  )
}
