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
import { useCreateMaterial } from '@/lib/queries/library'
import { insertMaterialSchema } from '@/lib/schemas/post'
import { BookPlus } from 'lucide-react'
import { useState } from 'react'
import { toast } from 'sonner'

export function AddMaterialDialog({ collectionID }: { collectionID: string }) {
  const { mutate, isError } = useCreateMaterial(collectionID)
  const [open, setOpen] = useState(false)
  const form = useAppForm({
    defaultValues: {
      title: '',
    },
    onSubmit: ({ value }) => {
      mutate({
        body: {
          collection_id: collectionID,
          title: value.title,
        },
      })
      if (isError) {
        toast.error('Failed to update collection')
        return
      }
      setOpen(false)
      toast.success('Material Created Successfully')
    },
    validators: {
      onChange: insertMaterialSchema,
    },
  })
  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button variant={'ghost'}>
          <BookPlus />
        </Button>
      </DialogTrigger>
      <DialogContent className="rounded-2xl">
        <DialogHeader>
          <DialogTitle className="text-xl">Add New Material</DialogTitle>
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
          </FieldGroup>
          <DialogFooter className="flex justify-between mt-4">
            <Button type="submit">Create Material</Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  )
}
