import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query"
import { httpClient } from "../api"
import { queryKeys } from "./keys"
import type { Collection, CreateCollectionReq, CreateMaterialReq, UpdateCollectionReq } from "../api/types"

type UseMaterialsOptions = {
	archived?: boolean
	enabled?: boolean
}
export const useMaterials = (
	collectionID: string,
	options?: UseMaterialsOptions
) => {
	const { archived = false, enabled = true } = options ?? {}
	return useQuery({
		queryFn: () => httpClient.getMaterialsByCollection(collectionID, archived),
		queryKey: queryKeys.library.materials(collectionID),
		staleTime: Infinity,
		enabled: enabled,
		retry: false
	})
}

export const useCollections = () => {
	return useQuery({
		queryFn: async () => {
			const res = await httpClient.getAllCollections(false)
			return res.data
		},
		queryKey: queryKeys.library.collections(),
		staleTime: Infinity
	})
}

export const useCreateCollection = () => {
	const queryClient = useQueryClient()

	return useMutation({
		mutationFn: async ({ body }: { body: CreateCollectionReq }) => {
			const res = await httpClient.newCollection(body)
			return res.data
		},
		onSuccess: (newCollection) => {
			console.table(newCollection)
			queryClient.setQueryData<Collection[]>(queryKeys.library.collections(), (old) => {
				return old ? [...old, newCollection] : [newCollection]
			})
		},
	})
}

export const useMutateCollection = (id: string) => {
	const queryClient = useQueryClient()

	return useMutation({
		mutationFn: async ({ body }: { body: UpdateCollectionReq }) => {
			const res = await httpClient.updateCollection(body, id)
			return res?.data
		},
		mutationKey: queryKeys.library.update(id),

		onSuccess: (updatedCollection: Collection) => {
			queryClient.setQueryData(queryKeys.library.collections(), (old: Collection[]) => {
				return old.map((c) => c.id === updatedCollection.id ? updatedCollection : c)
			})
		},
	})
}
// queryClient.invalidateQueries({ queryKey: queryKeys.library.materials(collectionID) })
export const useCreateMaterial = (collectionID: string) => {
	const queryClient = useQueryClient()

	return useMutation({
		mutationFn: ({ body }: { body: CreateMaterialReq }) => httpClient.newMaterial(body),
		onSuccess: (newMaterial) => {
			queryClient.setQueryData(queryKeys.library.materials(collectionID), (old: Collection[]) => {
				if (!old) return [newMaterial.data]
				return [...old, newMaterial.data]
			})
		}
	})
}

