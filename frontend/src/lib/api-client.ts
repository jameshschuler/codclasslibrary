import Axios, { InternalAxiosRequestConfig } from 'axios'
import { supabase } from '../supabase'

async function authRequestInterceptor(config: InternalAxiosRequestConfig) {
  if (config.headers) {
    config.headers.Accept = 'application/json'
  }

  if (config.url?.includes('me')) {
    try {
      const user = await supabase.auth.getUser()
      if (user) {
        const {
          data: { session },
        } = await supabase.auth.getSession()

        if (session) {
          config.headers.Authorization = `Bearer ${session?.access_token}`
        }
      }
    } catch (e) {
      const searchParams = new URLSearchParams()
      const redirectTo = searchParams.get('redirectTo')
      window.location.href = `/login?redirectTo=${redirectTo}`
    }
  }

  config.withCredentials = true
  return config
}

export const api = Axios.create({
  baseURL: import.meta.env.VITE_APP_API_URL,
})

api.interceptors.request.use(authRequestInterceptor)
api.interceptors.response.use(
  (response) => {
    return response.data
  },
  (error) => {
    const message = error.response?.data?.message || error.message
    // TODO: handle errors?
    // useNotifications.getState().addNotification({
    //   type: 'error',
    //   title: 'Error',
    //   message,
    // });
    if (error.response?.status === 401) {
      const searchParams = new URLSearchParams()
      const redirectTo = searchParams.get('redirectTo')
      window.location.href = `/login?redirectTo=${redirectTo}`
    }

    return Promise.reject(error)
  },
)
