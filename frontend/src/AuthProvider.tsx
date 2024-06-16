import { User } from '@supabase/auth-js/dist/module/lib/types'
import { AuthError, AuthTokenResponsePassword } from '@supabase/supabase-js'
import {
  PropsWithChildren,
  createContext,
  useContext,
  useEffect,
  useState,
} from 'react'
import { supabase } from './supabase'

interface AuthContext {
  auth?: boolean
  user?: User | null
  login?: (
    email: string,
    password: string,
  ) => Promise<AuthTokenResponsePassword>
  signOut?: () => Promise<{ error: AuthError | null }>
}

const AuthContext = createContext<AuthContext>({})

export const useAuth = () => useContext(AuthContext)

const login = (email: string, password: string) =>
  supabase.auth.signInWithPassword({ email, password })

const signOut = () => supabase.auth.signOut()

const AuthProvider = ({ children }: PropsWithChildren) => {
  const [user, setUser] = useState<User | null>()
  const [auth, setAuth] = useState(false)
  const [loading, setLoading] = useState<boolean>(true)

  useEffect(() => {
    setLoading(true)
    const getUser = async () => {
      const { data } = await supabase.auth.getUser()
      const { user: currentUser } = data
      setUser(currentUser ?? null)
      setAuth(currentUser ? true : false)
      setLoading(false)
    }
    getUser()

    const { data } = supabase.auth.onAuthStateChange((event, session) => {
      if (event === 'SIGNED_IN') {
        setUser(session?.user)
        setAuth(true)
      } else if (event === 'SIGNED_OUT') {
        setUser(null)
        setAuth(false)
      }
    })
    return () => {
      data.subscription.unsubscribe()
    }
  }, [])

  return (
    <AuthContext.Provider value={{ auth, user, login, signOut }}>
      {!loading && children}
    </AuthContext.Provider>
  )
}

export default AuthProvider
