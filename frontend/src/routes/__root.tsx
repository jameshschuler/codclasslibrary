import { Outlet, createRootRoute } from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/router-devtools'
import { Navbar } from '../components/Navbar'

export const Route = createRootRoute({
  component: () => (
    <>
      <Navbar />

      {/* <div className="p-2 flex gap-2">
        <Link to="/" className="[&.active]:font-bold">
          CoD Class Library
        </Link>
      </div> */}
      <hr />
      <Outlet />
      <TanStackRouterDevtools />
    </>
  ),
})
