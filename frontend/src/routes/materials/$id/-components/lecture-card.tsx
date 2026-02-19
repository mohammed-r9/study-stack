import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import type { Lecture } from '@/lib/api/types'
import { Calendar, HardDrive, FileText } from 'lucide-react'

// Helper to convert bytes into a human-readable string
function formatBytes(bytes: string | number, decimals = 1) {
  const b = typeof bytes === 'string' ? parseInt(bytes, 10) : bytes
  if (!b || b === 0) return '0 Bytes'

  const k = 1024
  const dm = decimals < 0 ? 0 : decimals
  const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB']

  const i = Math.floor(Math.log(b) / Math.log(k))

  return `${parseFloat((b / Math.pow(k, i)).toFixed(dm))} ${sizes[i]}`
}

type LectureCardProps = {
  lecture: Lecture
}

export default function LectureCard({ lecture }: LectureCardProps) {
  return (
    <Card className="overflow-hidden transition-all hover:shadow-md group h-full flex flex-col">
      <div className="flex items-center justify-center h-40 bg-muted/40 group-hover:bg-muted/60 transition-colors border-b hover:cursor-pointer">
        <FileText className="w-16 h-16 text-muted-foreground/60 group-hover:text-primary/80 transition-all duration-300 group-hover:scale-110" />
      </div>

      <CardHeader className="p-4 grow">
        <CardTitle
          className="text-base font-semibold line-clamp-2"
          title={lecture.title}
        >
          {lecture.title}
        </CardTitle>
      </CardHeader>

      <CardContent className="p-4 pt-0 space-y-2 text-xs text-muted-foreground">
        <div className="flex items-center gap-2">
          <HardDrive className="w-3.5 h-3.5" />
          <span className="font-medium text-foreground/80">
            {formatBytes(lecture.file_size)}
          </span>
        </div>

        <div className="flex items-center gap-2">
          <Calendar className="w-3.5 h-3.5" />
          <span>
            {new Date(lecture.created_at).toLocaleDateString(undefined, {
              month: 'short',
              day: 'numeric',
              year: 'numeric',
            })}
          </span>
        </div>
      </CardContent>
    </Card>
  )
}
