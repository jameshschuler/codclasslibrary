import {
  NavigationMenu,
  NavigationMenuItem,
  NavigationMenuList,
  navigationMenuTriggerStyle,
} from '@/components/ui/navigation-menu'
import { Link } from '@tanstack/react-router'

export function Navbar() {
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
            Community Library
          </Link>
        </NavigationMenuItem>
        <NavigationMenuItem>
          <Link to="/loadouts" className={navigationMenuTriggerStyle()}>
            My Library
          </Link>
        </NavigationMenuItem>
      </NavigationMenuList>
      <NavigationMenuList>
        <NavigationMenuItem>
          <Link to="/login" className={navigationMenuTriggerStyle()}>
            Documentation
          </Link>
        </NavigationMenuItem>
      </NavigationMenuList>
    </NavigationMenu>
  )
}
