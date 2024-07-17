import { queryOptions, useQuery } from '@tanstack/react-query'
import { api } from '../lib/api-client'
import { QueryConfig } from '../lib/react-query'

interface Loadout {
  id: string
  title: string
}

export const getComments = ({
  discussionId,
}: {
  discussionId?: string
}): Promise<Loadout[]> => {
  return api.get(`/me/loadouts`, {
    params: {
      discussionId,
    },
  })
}

export const getCommentsQueryOptions = (discussionId?: string) => {
  return queryOptions({
    queryKey: ['comments', discussionId],
    queryFn: () => getComments({ discussionId }),
  })
}

type UseCommentsOptions = {
  discussionId: string
  queryConfig?: QueryConfig<typeof getComments>
}

// TODO: use useInfiniteQuery
export const useComments = ({
  discussionId,
  queryConfig,
}: UseCommentsOptions) => {
  return useQuery({
    ...getCommentsQueryOptions(discussionId),
    ...queryConfig,
  })
}
