import { QueryClient } from '@tanstack/react-query'
import {
  Outlet,
  createRootRouteWithContext,
  useRouterState,
} from '@tanstack/react-router'
import { Navbar } from '../components/Navbar'
import { Spinner } from '../components/Spinner'

interface MyRouterContext {
  auth: boolean
  queryClient: QueryClient
}

// TODO: can this be a skeleton loader?
function RouterSpinner() {
  const isLoading = useRouterState({ select: (s) => s.status === 'pending' })
  return <Spinner show={isLoading} />
}

export const Route = createRootRouteWithContext<MyRouterContext>()({
  component: () => (
    <>
      <RouterSpinner />
      <Navbar />
      <hr />
      <Outlet />
      {/* <ReactQueryDevtools buttonPosition="top-right" />
      <TanStackRouterDevtools position="bottom-right" /> */}
    </>
  ),
})
