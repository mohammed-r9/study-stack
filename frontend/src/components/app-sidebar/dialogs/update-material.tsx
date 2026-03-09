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
import { useMutateMaterial } from '@/lib/queries/library'
import { updateMaterialSchema } from '@/lib/schemas/update'
import { Pen, Trash2 } from 'lucide-react'
import { useState } from 'react'

export function UpdateMaterialDialog({ id }: { id: string }) {
  const [open, setOpen] = useState(false)
  const { mutate } = useMutateMaterial()
  const form = useAppForm({
    defaultValues: {
      title: '',
    },
    onSubmit: ({ value }) => {
      mutate({
        params: {
          title: value.title,
          id: id,
        },
      })
      setOpen(false)
      form.reset()
    },
    validators: {
      onChange: updateMaterialSchema,
    },
  })
  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button variant="secondary" size={'lg'}>
          Update material
          <Pen />
        </Button>
      </DialogTrigger>
      <DialogContent className="rounded-2xl">
        <DialogHeader>
          <DialogTitle className="text-xl">Update Material</DialogTitle>
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
          <div className="flex justify-between mt-4">
            <Button type="submit" size={'lg'} className="flex-1">
              Update material
            </Button>

            <Button variant="destructive" className="gap-2" size={'lg'}>
              <Trash2 className="size-4" />
              Delete
            </Button>
          </div>
        </form>
      </DialogContent>
    </Dialog>
  )
}
