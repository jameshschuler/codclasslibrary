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
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/loadouts/')({
  component: () => (
    <main className="" style={{ height: 'calc(100% - 56px)' }}>
      <PageHeader title="My Library" />
      <Separator />
      <div className="flex mt-4">
        <div className="p-4" style={{ flex: '0 0 300px' }}>
          <p>Filters</p>
        </div>
        <div className="p-4" style={{ flex: '0 0 400px' }}>
          List
          {/* TODO: set height to viewport height */}
          <ScrollArea className="h-[500px] rounded-md border p-4">
            {[1, 2, 3, 4, 5, 6, 7].map((v) => {
              return (
                <Card key={v}>
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
