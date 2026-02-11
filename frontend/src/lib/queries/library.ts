import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query"
import { httpClient } from "../api"
import { queryKeys } from "./keys"
import type { CreateCollectionReq, CreateMaterialReq, UpdateCollectionReq } from "../api/types"

export const useMaterials = (collectionID: string, getArchived = false as boolean) => {
	return useQuery({
		queryFn: () => httpClient.getMaterialsByCollection(collectionID, getArchived),
		queryKey: queryKeys.library.materials(collectionID),
		staleTime: Infinity
	})
}

export const useCollections = (getArchived = false as boolean) => {
	return useQuery({
		queryFn: () => httpClient.getAllCollections(getArchived),
		queryKey: queryKeys.library.collections(),
		staleTime: Infinity
	})
}

export const useMutateCollection = (id: string) => {
	const queryClient = useQueryClient()

	return useMutation({
		mutationFn: ({ body }: { body: UpdateCollectionReq }) =>
			httpClient.updateCollection(body, id),
		mutationKey: queryKeys.library.update(id),

		onSuccess: () => {
			queryClient.invalidateQueries({
				queryKey: queryKeys.library.collections(),
			})
			queryClient.invalidateQueries({
				queryKey: queryKeys.library.collection(id),
			})
		},
	})
}

export const useCreateMaterial = (collectionID: string) => {
	const queryClient = useQueryClient()

	return useMutation({
		mutationFn: ({ body }: { body: CreateMaterialReq }) => httpClient.newMaterial(body),
		onSuccess: () => queryClient.invalidateQueries({ queryKey: queryKeys.library.materials(collectionID) })
	})
}

export const useCreateCollection = () => {
	const queryClient = useQueryClient()

	return useMutation({
		mutationFn: ({ body }: { body: CreateCollectionReq }) => httpClient.newCollection(body),
		onSuccess: () => queryClient.invalidateQueries({ queryKey: queryKeys.library.collections() })
	})
}
