import { Button } from '@/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import { FolderPlus } from 'lucide-react'

export function AddCollectionDialog() {
  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button>
          <FolderPlus />
        </Button>
      </DialogTrigger>
      <DialogContent className="rounded-2xl">
        <DialogHeader>
          <DialogTitle className="text-xl">Add New Collection</DialogTitle>
        </DialogHeader>

        <div className="py-4">form</div>

        <DialogFooter>
          <Button type="submit">Create</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}
