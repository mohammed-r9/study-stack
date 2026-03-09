import { useInfiniteQuery, useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { queryKeys } from "./keys";
import { httpClient } from "../api";
import { toast } from "sonner";
import type { UpdateFlashcardParams } from "../api/types";

export const useFlashcard = (cursor: number) =>
	useQuery({
		queryKey: [queryKeys.flashcards.study(), cursor],
		queryFn: () => httpClient.getOneFlashcard(),
		staleTime: 0,
	})


export const useInfiniteFlashcards = () => {
	return useInfiniteQuery({
		queryKey: queryKeys.flashcards.list(),

		queryFn: async ({ pageParam }) => {
			const response = await httpClient.getAllFlashcards({
				last_seen_flashcard_id: pageParam,
			})
			return response.data
		},
		initialPageParam: "",
		getNextPageParam: (lastPage) => {
			if (!lastPage.has_next_page || lastPage.flashcards.length === 0) {
				return undefined
			}

			return lastPage.flashcards[lastPage.flashcards.length - 1].id
		},

		staleTime: Infinity
	})
}

export const useMutateFlashcard = () => {
	const queryClient = useQueryClient()

	return useMutation({
		mutationFn: async ({ body: params }: { body: UpdateFlashcardParams }) => {
			const res = await httpClient.updateFlashcard(params)
			return res.data
		},
		mutationKey: queryKeys.flashcards.list(),

		onSuccess: (newFlashcard) => {
			queryClient.setQueryData(
				queryKeys.flashcards.list(),
				(old: any) => {
					return {
						...old,
						pages: old.pages.map((page: any) => ({
							...page,
							flashcards: page.flashcards?.map((f: any) =>
								f.id === newFlashcard.id ? newFlashcard : f
							) || []
						}))
					}
				}
			)

			toast.success('Flashcard updated successfully');
		},
	})
}


export const useDeleteFlashcard = () => {
	const queryClient = useQueryClient()

	return useMutation({
		mutationFn: async ({ id }: { id: string }) => {
			const res = await httpClient.deleteFlashcard(id)
			return res?.data
		},
		mutationKey: queryKeys.flashcards.list(),

		onSuccess: (deletedId) => {
			queryClient.setQueryData(
				queryKeys.flashcards.list(),
				(old: any) => {
					if (!old) return old;

					return {
						...old,
						pages: old.pages.map((page: any) => ({
							...page,
							flashcards: page.flashcards.filter((f: any) => f.id !== deletedId)
						}))
					};
				}
			);
			toast.success('Flashcard deleted successfully');
		},
	})
}
