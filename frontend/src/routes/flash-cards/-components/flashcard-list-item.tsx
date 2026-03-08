import { Button } from '@/components/ui/button'
import { Card, CardContent, CardFooter, CardHeader } from '@/components/ui/card'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { Badge } from '@/components/ui/badge' // Assuming you have badge installed
import { MoreHorizontal, Pencil, Trash, BookOpen, Layers } from 'lucide-react'

type Props = {
  id: string
  material_title: string
  back: string
  front: string
  last_used: Date
}

export default function FlashcardListItem({
  id,
  material_title,
  front,
  back,
  last_used,
}: Props) {
  return (
    <Card className="group w-full transition-all duration-200 hover:border-primary/50 hover:shadow-sm">
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-3">
        <div className="flex flex-col gap-1">
          <div className="flex items-center gap-2">
            <Layers className="h-4 w-4 text-muted-foreground" />
            <span className="text-xs font-medium uppercase tracking-wider text-muted-foreground">
              Flashcard
            </span>
          </div>
          {material_title && (
            <Badge variant="secondary" className="w-fit font-normal">
              {material_title}
            </Badge>
          )}
        </div>

        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button
              variant="ghost"
              size="icon"
              className="h-8 w-8 opacity-0 group-hover:opacity-100 transition-opacity"
            >
              <MoreHorizontal className="h-4 w-4" />
              <span className="sr-only">Open menu</span>
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent align="end" className="w-[160px]">
            <DropdownMenuItem className="cursor-pointer">
              <Pencil className="mr-2 h-4 w-4" />
              Edit Card
            </DropdownMenuItem>
            <DropdownMenuSeparator />
            <DropdownMenuItem className="cursor-pointer text-destructive focus:text-destructive">
              <Trash className="mr-2 h-4 w-4" />
              Delete
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </CardHeader>

      <CardContent className="grid gap-4 py-4">
        <div className="grid grid-cols-2 gap-4">
          <div className="space-y-1">
            <p className="text-[10px] font-bold uppercase text-muted-foreground/70">
              Front
            </p>
            <p className="text-sm leading-relaxed">{front}</p>
          </div>
          <div className="space-y-1 border-l pl-4">
            <p className="text-[10px] font-bold uppercase text-muted-foreground/70">
              Back
            </p>
            <p className="text-sm leading-relaxed font-medium text-foreground/90">
              {back}
            </p>
          </div>
        </div>
      </CardContent>

      <CardFooter className="flex items-center border-t bg-muted/30 px-6 py-3 self-stretch justify-end">
        <div className="flex items-center gap-1.5 text-xs text-muted-foreground">
          <BookOpen className="h-3 w-3" />
          <span>Last practiced {new Date(last_used).toLocaleDateString()}</span>
        </div>
      </CardFooter>
    </Card>
  )
}
