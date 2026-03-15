type Props = {
  imageFile: File | null
}

export default function Labels({ imageFile }: Props) {
  if (!imageFile) return
  return (
    <div className="w-[90%] mt-2">
      <img src={URL.createObjectURL(imageFile)} className="max-w-full" />
    </div>
  )
}
