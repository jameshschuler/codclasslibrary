import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/loadouts/')({
  component: () => (
    <div>
      <h1>My Library</h1>
      Hello /loadouts/!
    </div>
  ),
})
