import { QueryClientProvider } from '@tanstack/react-query'
import {
  ErrorComponent,
  Link,
  RouterProvider,
  createRouter,
} from '@tanstack/react-router'
import AuthProvider, { useAuth } from './AuthProvider'
import { Spinner } from './components/Spinner'
import { queryClient } from './lib/react-query'
import { routeTree } from './routeTree.gen'

declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router
  }
}

const router = createRouter({
  routeTree,
  // TODO: use full screen splash screen?
  defaultPendingComponent: () => (
    <div className={`p-2 text-2xl`}>
      <Spinner />
    </div>
  ),
  defaultErrorComponent: ({ error }) => <ErrorComponent error={error} />,
  context: {
    auth: undefined!,
    queryClient,
  },
  defaultPreload: 'intent',
  defaultPreloadStaleTime: 0,
  defaultNotFoundComponent: () => {
    return (
      <div>
        <p>Not found!</p>
        <Link to="/">Go home</Link>
      </div>
    )
  },
})

function InnerApp() {
  const { auth } = useAuth()
  return (
    <RouterProvider
      router={router}
      context={{ auth }}
      defaultPreload="intent"
    />
  )
}

export function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <AuthProvider>
        <InnerApp />
      </AuthProvider>
    </QueryClientProvider>
  )
}
