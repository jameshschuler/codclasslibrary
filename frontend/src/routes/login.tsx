import { Auth } from '@supabase/auth-ui-react'
import { ThemeSupa } from '@supabase/auth-ui-shared'
import { createFileRoute, useRouter } from '@tanstack/react-router'
import { useLayoutEffect } from 'react'
import { z } from 'zod'
import { supabase } from '../supabase'

export const Route = createFileRoute('/login')({
  validateSearch: z.object({
    redirect: z.string().optional(),
  }),
  component: () => {
    const router = useRouter()
    const search = Route.useSearch()

    const { auth } = Route.useRouteContext({
      select: ({ auth }) => ({ auth }),
    })

    useLayoutEffect(() => {
      if (auth && search.redirect) {
        router.history.push(search.redirect)
      }
    }, [search.redirect])

    return (
      <div style={{ width: '50%', margin: '0px auto' }}>
        <Auth
          supabaseClient={supabase}
          appearance={{
            theme: ThemeSupa,
          }}
          providers={[]}
          //theme="dark"
          redirectTo="/"
          showLinks
        ></Auth>
      </div>
    )
  },
})
