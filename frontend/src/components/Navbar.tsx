import {
  NavigationMenu,
  NavigationMenuItem,
  NavigationMenuList,
  navigationMenuTriggerStyle,
} from '@/components/ui/navigation-menu'
import { Link } from '@tanstack/react-router'
import { useAuth } from '../AuthProvider'

export function Navbar() {
  const { auth } = useAuth()
  return (
    <NavigationMenu className="flex justify-between max-w-100 px-4">
      <NavigationMenuList className="my-2">
        <NavigationMenuItem>
          <Link
            to="/"
            className={`${navigationMenuTriggerStyle()} font-semibold text-lg`}
          >
            CoD Class Library
          </Link>
        </NavigationMenuItem>
        <NavigationMenuItem>
          <Link to="/community" className={navigationMenuTriggerStyle()}>
            Browse Community Library
          </Link>
        </NavigationMenuItem>
        <NavigationMenuItem>
          <Link to="/loadouts" className={navigationMenuTriggerStyle()}>
            Your Library
          </Link>
        </NavigationMenuItem>
      </NavigationMenuList>
      <NavigationMenuList>
        {!auth && (
          <NavigationMenuItem>
            <Link to="/login" className={navigationMenuTriggerStyle()}>
              Documentation
            </Link>
          </NavigationMenuItem>
        )}
      </NavigationMenuList>
    </NavigationMenu>
  )
}
