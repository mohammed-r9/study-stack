import {
  Field,
  FieldDescription,
  FieldError,
  FieldLabel,
} from '@/components/ui/field'
import { Input } from '@/components/ui/input'
import { useFieldContext } from '@/hooks/form-context'
import type { HTMLInputTypeAttribute } from 'react'

type FieldProps = {
  id: string
  label: string
  description?: string
  placeholder?: string
  type?: HTMLInputTypeAttribute
}

/**
 * Field > Label + Description + Input
 */
function AppField({
  id,
  label,
  description,
  placeholder,
  type = 'text',
}: FieldProps) {
  const field = useFieldContext<string>()
  const hasError =
    field.state.meta.isTouched && field.state.meta.errors.length > 0

  return (
    <Field>
      <FieldLabel htmlFor={id}>{label}</FieldLabel>

      {description && <FieldDescription>{description}</FieldDescription>}

      <Input
        id={id}
        type={type}
        placeholder={placeholder}
        value={(field.state.value as string) ?? ''}
        onBlur={field.handleBlur}
        onChange={(e) => field.handleChange(e.target.value)}
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
