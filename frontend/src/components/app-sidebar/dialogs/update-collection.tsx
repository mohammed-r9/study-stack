import AppField from '@/components/form/app-field'
import { Button } from '@/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import { FieldGroup } from '@/components/ui/field'
import { useAppForm } from '@/hooks/form'
import { useMutateCollection } from '@/lib/queries/library'
import { updateCollectionSchema } from '@/lib/schemas/update'
import { Pen, Trash2 } from 'lucide-react'
import { useState } from 'react'
import { toast } from 'sonner'

export default function UpdateCollectionDialog({
  collectionID,
}: {
  collectionID: string
}) {
  const { mutate, isError } = useMutateCollection(collectionID)
  const [open, setOpen] = useState(false)
  const form = useAppForm({
    defaultValues: {
      new_title: '',
      new_description: '',
    },
    onSubmit: ({ value }) => {
      mutate({
        body: {
          new_title: value.new_title,
          new_description: value.new_description,
        },
      })
      if (isError) {
        toast.error('Failed to update collection')
        return
      }
      setOpen(false)
      toast.success('Collection updated successfully')
    },
    validators: {
      onChange: updateCollectionSchema,
    },
  })
  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button variant="ghost" className="bg-accent-foreground/10 p-1.5!">
          <Pen />
        </Button>
      </DialogTrigger>
      <DialogContent className="rounded-2xl">
        <DialogHeader>
          <DialogTitle className="text-xl">Update Collection</DialogTitle>
        </DialogHeader>

        <form
          onSubmit={(e) => {
            e.preventDefault()
            form.handleSubmit()
          }}
        >
          <FieldGroup>
            <form.AppField name="new_title">
              {() => (
                <AppField
                  type="text"
                  id="new_title"
                  label="New Title"
                  placeholder=""
                />
              )}
            </form.AppField>
            <form.AppField name="new_description">
              {() => (
                <AppField
                  type="text"
                  id="new_description"
                  label="New Description"
                  placeholder=""
                />
              )}
            </form.AppField>
          </FieldGroup>
          <DialogFooter className="flex justify-between mt-4">
            <Button variant="destructive" className="gap-2">
              <Trash2 className="size-4" />
              Delete
            </Button>
            <Button type="submit">Save Changes</Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  )
}
