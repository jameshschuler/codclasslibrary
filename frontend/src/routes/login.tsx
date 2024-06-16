import { Auth } from '@supabase/auth-ui-react'
import { ThemeSupa } from '@supabase/auth-ui-shared'
import { createFileRoute, redirect } from '@tanstack/react-router'
import { supabase } from '../supabase'

export const Route = createFileRoute('/login')({
  beforeLoad: ({ context, location }) => {
    console.log(location)
    if (context.auth) {
      throw redirect({
        to: '/loadouts',
      })
    }
  },
  component: () => (
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
  ),
})
