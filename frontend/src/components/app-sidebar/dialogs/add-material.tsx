import { Button } from '@/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import { BookPlus } from 'lucide-react'

export function AddMaterialDialog() {
  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button>
          <BookPlus />
        </Button>
      </DialogTrigger>
      <DialogContent className="rounded-2xl">
        <DialogHeader>
          <DialogTitle className="text-xl">Add New Material</DialogTitle>
        </DialogHeader>

        <div className="py-4">form</div>

        <DialogFooter>
          <Button type="submit">Create</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}
