import AppField from '@/components/form/app-field'
import { Button } from '@/components/ui/button'
import { Card, CardDescription, CardTitle } from '@/components/ui/card'
import { FieldGroup } from '@/components/ui/field'
import { useAppForm } from '@/hooks/form'
import { httpClient } from '@/lib/api'
import { loginSchema } from '@/lib/schemas/auth'
import { createFileRoute, Link, useNavigate } from '@tanstack/react-router'
import { toast } from 'sonner'

export const Route = createFileRoute('/login')({
  component: RouteComponent,
})

function RouteComponent() {
  const navigate = useNavigate()
  const form = useAppForm({
    defaultValues: {
      email: '',
      password: '',
    },
    onSubmit: async ({ value }) => {
      try {
        await httpClient.login({ email: value.email, password: value.password })
        toast.success('Login successful')
        navigate({ to: '/materials', replace: true })
      } catch (e: any) {
        console.error(e)
        toast.error(e?.response.data)
      }
    },
    validators: {
      onChange: loginSchema,
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
          </FieldGroup>
          <div className="mt-3 flex flex-col gap-3">
            <Button type="submit" className="w-full">
              Login
            </Button>
            <div className="flex flex-col items-center">
              Don't have an account yet?{' '}
              <Button asChild variant={'link'} className="inline">
                <Link to="/register">Create an account</Link>
              </Button>
            </div>
          </div>
        </form>
      </Card>
    </div>
  )
}
