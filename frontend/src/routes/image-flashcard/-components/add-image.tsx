import { Input } from '@/components/ui/input'
import { Image } from 'lucide-react'
import { useRef, type ChangeEvent } from 'react'

type propsT = {
  setImageFile: React.Dispatch<React.SetStateAction<File | null>>
  imageFile: File | null
}
export default function AddImage({ setImageFile, imageFile }: propsT) {
  const fileInputRef = useRef<HTMLInputElement>(null)
  const preview = imageFile ? URL.createObjectURL(imageFile) : null
  function handleUpload() {
    fileInputRef.current?.click()
  }

  function handleChange(e: ChangeEvent<HTMLInputElement>) {
    if (!e.target.files || e.target.files.length === 0) return
    const file = e.target.files[0]
    setImageFile(file)
  }

  return (
    <div
      onClick={handleUpload}
      className="w-64 h-64 border-2 border-dotted border-muted-foreground rounded-xl p-4 flex flex-col justify-center items-center gap-4 hover:border-primary hover:bg-primary/5 transition-colors cursor-pointer"
    >
      {preview ? (
        <img
          src={preview}
          alt="Preview"
          className="w-full h-full object-contain rounded-lg"
        />
      ) : (
        <>
          <Image size={48} className="text-muted-foreground" />
          <p className="text-sm text-muted-foreground text-center">
            Click to select an image
          </p>
        </>
      )}
      <Input
        type="file"
        ref={fileInputRef}
        onChange={handleChange}
        accept="image/png, image/jpeg, image/jpg"
        style={{ display: 'none' }}
      />
    </div>
  )
}
