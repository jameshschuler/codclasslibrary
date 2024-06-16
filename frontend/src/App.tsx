import { Link, RouterProvider, createRouter } from '@tanstack/react-router'
import AuthProvider, { useAuth } from './AuthProvider'
import { routeTree } from './routeTree.gen'

declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router
  }
}

// Create a new router instance
const router = createRouter({
  routeTree,
  context: {
    // auth will initially be undefined
    // We'll be passing down the auth state from within a React component
    auth: undefined!,
  },
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
  return <RouterProvider router={router} context={{ auth }} />
}

export function App() {
  return (
    <AuthProvider>
      <InnerApp />
    </AuthProvider>
  )
}

// TODO: delete
// import { Button } from '@/components/ui/button'
// import {
//   Form,
//   FormControl,
//   FormField,
//   FormItem,
//   FormLabel,
//   FormMessage,
// } from '@/components/ui/form'
// import { Input } from '@/components/ui/input'
// import { zodResolver } from '@hookform/resolvers/zod'
// import {
//   DefaultOptions,
//   QueryClient,
//   QueryClientProvider,
//   useQuery,
// } from '@tanstack/react-query'
// import axios from 'axios'
// import { useForm } from 'react-hook-form'

// import { z } from 'zod'

// const queryConfig: DefaultOptions = {
//   queries: {
//     refetchOnWindowFocus: false,
//     retry: false,
//   },
// }

// export const queryClient = new QueryClient({ defaultOptions: queryConfig })

// const client = axios.create({
//   baseURL: 'http://localhost:3333/', //BASE_URL
// })

// interface Loadout {
//   id: string
//   title: string
// }

// function getLoadouts(): Promise<Loadout[]> {
//   return client.get(`/loadouts`).then((res) => res.data)
// }

// export const useGetLoadouts = () => {
//   return useQuery({
//     queryKey: ['loadouts'],
//     queryFn: () => getLoadouts(),
//   })
// }

// function Loadouts() {
//   const { data, isLoading } = useGetLoadouts()
//   console.log(data)

//   if (isLoading) {
//     return <div>Loading...</div>
//   }

//   return (
//     <div>
//       {data?.map((loadout: Loadout) => {
//         return (
//           <div key={loadout.id}>
//             <p>{loadout.title}</p>
//           </div>
//         )
//       })}
//     </div>
//   )
// }

// // "source": "YouTube",
// // "sourceUrl": "www.youtube.com",
// // "weaponCategory": "TestWeaponCategory",
// // "weaponName": "TestWeaponName",

// const categories = new Map<string, string>([
//   ['SubmachineGun', 'Submachine Gun'],
// ])

// const formSchema = z.object({
//   title: z.string().min(1).max(250),
//   source: z.string().min(1).max(250).optional(),
//   sourceUrl: z.string().min(1).max(250).optional(),
//   weaponCategory: z.string().min(1).max(250), // TODO: must be present in categories
//   weaponName: z.string().min(1).max(250),
// })

// function App() {
//   const form = useForm<z.infer<typeof formSchema>>({
//     resolver: zodResolver(formSchema),
//     defaultValues: {
//       title: '',
//     },
//   })

//   // 2. Define a submit handler.
//   function onSubmit(values: z.infer<typeof formSchema>) {
//     console.log(values)
//   }

//   import React, { StrictMode } from 'react'
//   import ReactDOM from 'react-dom/client'
//   import { RouterProvider, createRouter } from '@tanstack/react-router'

//   // Import the generated route tree
//   import { routeTree } from './routeTree.gen'

//   // Create a new router instance
//   const router = createRouter({ routeTree })

//   // Register the router instance for type safety

//   return (
//     <QueryClientProvider client={queryClient}>
//       <main className="p-12">
//         <Loadouts />
//         <Form {...form}>
//           <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
//             <FormField
//               control={form.control}
//               name="title"
//               render={({ field }) => (
//                 <FormItem>
//                   <FormLabel>Username</FormLabel>
//                   <FormControl>
//                     <Input placeholder="Enter a title..." {...field} />
//                   </FormControl>
//                   <FormMessage />
//                 </FormItem>
//               )}
//             />
//             <FormField
//               control={form.control}
//               name="source"
//               render={({ field }) => (
//                 <FormItem>
//                   <FormLabel>Source</FormLabel>
//                   <FormControl>
//                     <Input placeholder="Enter a source..." {...field} />
//                   </FormControl>
//                   <FormMessage />
//                 </FormItem>
//               )}
//             />
//             <FormField
//               control={form.control}
//               name="sourceUrl"
//               render={({ field }) => (
//                 <FormItem>
//                   <FormLabel>Source Url</FormLabel>
//                   <FormControl>
//                     <Input placeholder="Enter a source url..." {...field} />
//                   </FormControl>
//                   <FormMessage />
//                 </FormItem>
//               )}
//             />
//             <Button type="submit">Add</Button>
//           </form>
//         </Form>
//       </main>
//     </QueryClientProvider>
//   )
// }

// export default App
