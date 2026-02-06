import { Button } from '@/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import { Pen, Trash2 } from 'lucide-react'

export function UpdateMaterialDialog() {
  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button variant="outline">
          <Pen />
        </Button>
      </DialogTrigger>
      <DialogContent className="rounded-2xl">
        <DialogHeader>
          <DialogTitle className="text-xl">Update Material</DialogTitle>
        </DialogHeader>

        <div className="py-4">form</div>

        <DialogFooter className="flex justify-between">
          <Button variant="destructive" className="gap-2">
            <Trash2 className="size-4" />
            Delete
          </Button>
          <Button type="submit">Save Changes</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}
