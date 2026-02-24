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
import { useState, type Dispatch, type SetStateAction } from 'react'
import { toast } from 'sonner'

export default function UpdateCollectionDialog({
  collectionID,
  setIsHoveredOn,
}: {
  collectionID: string
  setIsHoveredOn: Dispatch<SetStateAction<boolean>>
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
  setIsHoveredOn(false)
  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button variant="ghost" className="p-3!">
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
          <DialogFooter className="flex mt-4 w-full">
            <Button variant="destructive" className="gap-2 flex-1">
              <Trash2 className="size-4" />
              Delete
            </Button>
            <Button type="submit" className="flex-1">
              Save Changes
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  )
}
