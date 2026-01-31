import AppField from '@/components/form/app-field'
import { Button } from '@/components/ui/button'
import { Card, CardDescription, CardTitle } from '@/components/ui/card'
import { FieldGroup } from '@/components/ui/field'
import { useAppForm } from '@/hooks/form'
import { httpClient } from '@/lib/api'
import { registerSchema } from '@/lib/schemas/auth'
import { createFileRoute, Link, useNavigate } from '@tanstack/react-router'
import { toast } from 'sonner'

export const Route = createFileRoute('/register')({
  component: RouteComponent,
})

function RouteComponent() {
  const navigate = useNavigate()
  const form = useAppForm({
    defaultValues: {
      name: '',
      email: '',
      password: '',
      confirm_password: '',
    },
    validators: {
      onChange: registerSchema,
    },

    onSubmit: async ({ value }) => {
      try {
        await httpClient.register({
          email: value.email,
          password: value.password,
          confirm_password: value.confirm_password,
          name: value.name,
        })
        navigate({ to: '/login', replace: true })
      } catch (e: any) {
        console.error(e)
        toast.error(e?.response.data)
      }
    },
  })
  return (
    <div className="h-screen w-full flex justify-center items-center">
      <Card className="p-4 w-md">
        <CardTitle>Login</CardTitle>
        <CardDescription>Login to your account please</CardDescription>
        <form
          onSubmit={(e) => {
            e.preventDefault()
            form.handleSubmit()
          }}
        >
          <FieldGroup>
            <form.AppField name="name">
              {(field) => (
                <AppField
                  type="text"
                  id="name"
                  label="Name"
                  placeholder="John Doe"
                />
              )}
            </form.AppField>
            <form.AppField name="email">
              {(field) => (
                <AppField
                  type="email"
                  id="email"
                  label="Email"
                  placeholder="me@example.com"
                />
              )}
            </form.AppField>
            <div className="flex gap-3">
              <form.AppField name="password">
                {(field) => (
                  <AppField
                    type="password"
                    id="password"
                    label="Password"
                    placeholder="password"
                  />
                )}
              </form.AppField>

              <form.AppField name="confirm_password">
                {(field) => (
                  <AppField
                    type="password"
                    id="confirm_password"
                    label="Confirm password"
                    placeholder="cofirm your password"
                  />
                )}
              </form.AppField>
            </div>
          </FieldGroup>
          <div className="mt-3 flex flex-col gap-3">
            <Button type="submit" className="w-full">
              Register
            </Button>
            <div className="flex flex-col items-center">
              Already have an account?{' '}
              <Button asChild variant={'link'} className="inline">
                <Link to="/login">Login</Link>
              </Button>
            </div>
          </div>
        </form>
      </Card>
    </div>
  )
}
