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
import { useCreateCollection } from '@/lib/queries/library'
import { insertCollectionSchema } from '@/lib/schemas/post'
import { FolderPlus } from 'lucide-react'
import { useState } from 'react'
import { toast } from 'sonner'

export function AddCollectionDialog() {
  const [open, setOpen] = useState(false)
  const { mutate, isError } = useCreateCollection()

  const form = useAppForm({
    defaultValues: {
      title: '',
      description: '',
    },
    onSubmit: ({ value }) => {
      mutate({
        body: {
          title: value.title,
          description: value.description,
        },
      })
      if (isError) {
        toast.error('Failed to update collection')
        return
      }
      setOpen(false)
      toast.success('Collection Created Successfully')
    },
    validators: {
      onChange: insertCollectionSchema,
    },
  })
  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button variant={'outline'} className="mb-3">
          <FolderPlus />
        </Button>
      </DialogTrigger>
      <DialogContent className="rounded-2xl">
        <DialogHeader>
          <DialogTitle className="text-xl">Add New Collection</DialogTitle>
        </DialogHeader>

        <form
          onSubmit={(e) => {
            e.preventDefault()
            form.handleSubmit()
          }}
        >
          <FieldGroup>
            <form.AppField name="title">
              {() => (
                <AppField type="text" id="title" label="Title" placeholder="" />
              )}
            </form.AppField>

            <form.AppField name="description">
              {() => (
                <AppField
                  type="text"
                  id="description"
                  label="Description"
                  placeholder=""
                />
              )}
            </form.AppField>
          </FieldGroup>
          <DialogFooter className="flex justify-between mt-4">
            <Button type="submit">Create Collection</Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  )
}
