import AppField from '@/components/form/app-field'
import { Button } from '@/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import { DropdownMenuItem } from '@/components/ui/dropdown-menu'
import { FieldGroup } from '@/components/ui/field'
import { useAppForm } from '@/hooks/form'
import { useMutateFlashcard } from '@/lib/queries/flashcards'
import { updateFlashcardSchema } from '@/lib/schemas/update'
import { Pencil, Trash2 } from 'lucide-react'
import { useState } from 'react'
import { toast } from 'sonner'

export default function UpdateFlashcardDialog({
  flashcardId,
}: {
  flashcardId: string
}) {
  const { mutate, isError } = useMutateFlashcard()
  const [open, setOpen] = useState(false)
  const form = useAppForm({
    defaultValues: {
      front: '',
      back: '',
    },
    onSubmit: ({ value }) => {
      mutate({
        body: {
          front: value.front,
          back: value.back,
          id: flashcardId,
        },
      })
      if (isError) {
        return
      }
      setOpen(false)
      form.reset()
      toast.success('Flashcard updated successfully')
    },
    validators: {
      onChange: updateFlashcardSchema,
    },
  })
  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <DropdownMenuItem
          className="cursor-pointer"
          onSelect={(e) => e.preventDefault()}
        >
          <Pencil className="mr-2 h-4 w-4" />
          Edit Card
        </DropdownMenuItem>
      </DialogTrigger>
      <DialogContent className="rounded-2xl">
        <DialogHeader>
          <DialogTitle className="text-xl">Update Flashcard</DialogTitle>
        </DialogHeader>

        <form
          onSubmit={(e) => {
            e.preventDefault()
            form.handleSubmit()
          }}
        >
          <FieldGroup>
            <form.AppField name="front">
              {() => (
                <AppField type="text" id="front" label="Front" placeholder="" />
              )}
            </form.AppField>
            <form.AppField name="back">
              {() => (
                <AppField type="text" id="back" label="back" placeholder="" />
              )}
            </form.AppField>
          </FieldGroup>
          <div className="flex mt-4 w-full bg-background gap-4">
            <Button variant="destructive" className="gap-2 flex-1">
              <Trash2 className="size-4" />
              Delete
            </Button>
            <Button type="submit" className="flex-1">
              Save Changes
            </Button>
          </div>
        </form>
      </DialogContent>
    </Dialog>
  )
}
