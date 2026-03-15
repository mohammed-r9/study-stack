import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { createFileRoute, redirect } from '@tanstack/react-router'
import { useState } from 'react'

export const Route = createFileRoute('/image-flashcard/')({
  component: RouteComponent,
  beforeLoad: () => {
    throw redirect({ to: '/image-flashcard/add', replace: true })
  },
})

function RouteComponent() {
  const [image, setImage] = useState<string | null>(null)
  return (
    <>
      Hello "/image-flashcard/"!
      <Label htmlFor="image">image</Label>
      <Input
        type="file"
        name="image"
        id="image"
        onChange={(e) => {
          const files = e.target.files
          if (!files) return
          const file = files[0]
          const url = URL.createObjectURL(file)
          setImage(url)
        }}
      />
      {image && <img src={image} />}
    </>
  )
}
