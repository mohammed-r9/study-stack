import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query"
import { httpClient } from "../api"
import { queryKeys } from "./keys"
import type { Collection, CreateCollectionReq, CreateMaterialReq, Material, UpdateCollectionReq } from "../api/types"
import type { AxiosError } from "axios"
import { toast } from "sonner"

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
		queryFn: async () => {
			const res = await httpClient.getMaterialsByCollection(collectionID, archived)
			return res?.data
		},
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
			return res?.data
		},
		onSuccess: (newCollection) => {
			console.table(newCollection)
			queryClient.setQueryData<Collection[]>(queryKeys.library.collections(), (old) => {
				return old ? [...old, newCollection] : [newCollection]
			})
		},
		onError: (e) => {
			const err = e as AxiosError
			if (err.status === 403) toast.error("You are only allowed to have a maximum of 20 collections")
		}
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
export const useCreateMaterial = (collectionID: string) => {
	const queryClient = useQueryClient()

	return useMutation({
		mutationFn: async ({ body }: { body: CreateMaterialReq }) => {
			const res = await httpClient.newMaterial(body)
			return res?.data
		},
		onSuccess: (newMaterial: Material) => {
			queryClient.setQueryData(queryKeys.library.materials(collectionID), (old: Material[]) => {
				if (!old) return [newMaterial]
				return [...old, newMaterial]
			})
		}
	})
}

