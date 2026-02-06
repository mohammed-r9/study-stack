import { queryOptions, useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { queryKeys } from "./keys";
import { httpClient } from "../api";
import type { UpdateCollectionReq } from "../api/types";


export const createUserQueryOptions = () => {
	return queryOptions({
		queryKey: queryKeys.currentUser,
		queryFn: httpClient.getCurrentUser,
		staleTime: Infinity
	})
}

export const useLibrary = () => useQuery({
	queryFn: () => httpClient.getLibrary(),
	queryKey: queryKeys.library,
	staleTime: Infinity
})

export const useMutateCollection = () => {
	const queryClient = useQueryClient()

	return useMutation({
		mutationFn: ({ id, body }: { id: string; body: UpdateCollectionReq }) =>
			httpClient.updateCollection(body, id),
		mutationKey: queryKeys.library,
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: queryKeys.library })
		},
	})
}
