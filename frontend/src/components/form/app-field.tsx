import {
  Field,
  FieldDescription,
  FieldError,
  FieldLabel,
} from '@/components/ui/field'
import { Input } from '@/components/ui/input'
import { useFieldContext } from '@/hooks/form-context'
import type { HTMLInputTypeAttribute } from 'react'

type FieldProps<T> = {
  id: string
  label: string
  description?: string
  placeholder?: string
  type?: HTMLInputTypeAttribute
  accept?: string
}

/**
 * Field > Label + Description + Input
 */
function AppField<T>({
  id,
  label,
  description,
  placeholder,
  type = 'text',
  accept,
}: FieldProps<T>) {
  const field = useFieldContext<T>()
  const hasError =
    field.state.meta.isTouched && field.state.meta.errors.length > 0
  const isFile = type === 'file'
  return (
    <Field>
      <FieldLabel htmlFor={id}>{label}</FieldLabel>

      {description && <FieldDescription>{description}</FieldDescription>}

      <Input
        id={id}
        type={type}
        accept={accept}
        placeholder={placeholder}
        value={isFile ? undefined : ((field.state.value as any) ?? '')}
        onBlur={field.handleBlur}
        onChange={(e) => {
          if (isFile) {
            field.handleChange((e.target.files?.[0] ?? null) as T)
          } else {
            field.handleChange(e.target.value as T)
          }
        }}
        aria-invalid={hasError}
      />

      {hasError && (
        <FieldError>
          {field.state.meta.errors[0]?.message ?? field.state.meta.errors[0]}
        </FieldError>
      )}
    </Field>
  )
}

export default AppField
