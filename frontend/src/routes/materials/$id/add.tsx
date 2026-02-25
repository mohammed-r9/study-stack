import { createFileRoute, Link, useRouter } from '@tanstack/react-router'
import type { MaterialSerach } from './route'
import { useAppForm } from '@/hooks/form'
import { lectureFormSchema } from '@/lib/schemas/post'
import AppField from '@/components/form/app-field'
import { FieldGroup } from '@/components/ui/field'
import { Button } from '@/components/ui/button'
import { ArrowLeft, ArrowUp } from 'lucide-react'
import { useCreateLecture } from '@/lib/queries/lectures'
import type { createLectureReq } from '@/lib/api/types'

export const Route = createFileRoute('/materials/$id/add')({
  component: RouteComponent,
  validateSearch: (search: Record<string, string>): MaterialSerach => {
    return {
      title: search?.title ?? '',
    }
  },
})

function RouteComponent() {
  const { id } = Route.useParams()
  const router = useRouter()
  const { title } = Route.useSearch()
  const { mutate } = useCreateLecture(id, title)
  const form = useAppForm({
    defaultValues: {
      lecture_title: '',
      lecture_file: null as File | null,
    },
    validators: {
      onChange: lectureFormSchema,
    },
    onSubmit: async ({ value }) => {
      const payload: createLectureReq = {
        lecture_title: value.lecture_title,
        lecture_file: value.lecture_file!,
        material_id: id,
      }
      mutate(payload)
    },
  })
  return (
    <div className="flex flex-col gap-5">
      <Button
        asChild
        variant={'secondary'}
        onClick={() => router.history.back()}
        className="self-start"
      >
        <Link to="/materials" params={{ id: id }} search={{ title: title }}>
          <ArrowLeft /> Go Back
        </Link>
      </Button>
      <form
        className="w-1/2"
        onSubmit={(e) => {
          e.preventDefault()
          form.handleSubmit()
        }}
      >
        <FieldGroup>
          <form.AppField name="lecture_title">
            {() => (
              <AppField id="lecture_title" type="text" label="Lecture Title" />
            )}
          </form.AppField>
          <form.AppField name="lecture_file">
            {() => (
              <AppField
                id="lecture_file"
                type="file"
                label="Lecture File"
                accept="application/pdf"
              />
            )}
          </form.AppField>
        </FieldGroup>
        <FieldGroup className="mt-4">
          <Button>
            {' '}
            <ArrowUp /> Upload Lecture{' '}
          </Button>
        </FieldGroup>
      </form>
    </div>
  )
}
