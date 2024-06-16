import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import { ScrollArea } from '@/components/ui/scroll-area'
import { Separator } from '@/components/ui/separator'
import { createFileRoute, redirect } from '@tanstack/react-router'

export const Route = createFileRoute('/loadouts/')({
  beforeLoad: ({ context, location }) => {
    console.log('context', context.auth)
    if (!context.auth) {
      throw redirect({
        to: '/login',
        search: {
          redirect: location.href,
        },
      })
    }
  },
  component: () => (
    <main className="" style={{ height: 'calc(100% - 56px)' }}>
      <PageHeader title="Your Library" />
      <Separator />
      <div className="flex mt-4">
        <div className="p-4" style={{ flex: '0 0 300px' }}>
          <p>Filters</p>
          <div className="space-y-4 py-4">
            <div className="px-3 py-2">
              <h2 className="mb-2 px-4 text-lg font-semibold tracking-tight">
                Discover
              </h2>
              <div className="space-y-1">
                <button className="inline-flex items-center whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 bg-secondary text-secondary-foreground shadow-sm hover:bg-secondary/80 h-9 px-4 py-2 w-full justify-start">
                  Listen Now
                </button>
                <button className="inline-flex items-center whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 hover:bg-accent hover:text-accent-foreground h-9 px-4 py-2 w-full justify-start">
                  Browse
                </button>
              </div>
            </div>

            <div className="px-3 py-2">
              <h2 className="mb-2 px-4 text-lg font-semibold tracking-tight">
                Discover
              </h2>
              <div className="space-y-1">
                <button className="inline-flex items-center whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 bg-secondary text-secondary-foreground shadow-sm hover:bg-secondary/80 h-9 px-4 py-2 w-full justify-start">
                  Listen Now
                </button>
                <button className="inline-flex items-center whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 bg-secondary text-secondary-foreground shadow-sm hover:bg-secondary/80 h-9 px-4 py-2 w-full justify-start">
                  Listen Now
                </button>
              </div>
            </div>
          </div>
        </div>
        <div className="p-4" style={{ flex: '0 0 400px' }}>
          List
          <ScrollArea className="h-[100dvh] rounded-md border p-4">
            {[1, 2, 3, 4, 5, 6, 7].map((v) => {
              return (
                <Card key={v} className="my-2">
                  <CardHeader>
                    <CardTitle>Card Title {v}</CardTitle>
                    <CardDescription>Card Description</CardDescription>
                  </CardHeader>
                  <CardContent>
                    <p>Card Content</p>
                  </CardContent>
                  <CardFooter>
                    <p>Card Footer</p>
                  </CardFooter>
                </Card>
              )
            })}
          </ScrollArea>
        </div>
        <div className="p-4 flex-1">Details</div>
      </div>
    </main>
  ),
})

interface PageHeaderProps {
  title: string
}

function PageHeader({ title }: PageHeaderProps) {
  return (
    <header className="flex items-center p-4">
      <h1 className="text-2xl font-medium">{title}</h1>
    </header>
  )
}
