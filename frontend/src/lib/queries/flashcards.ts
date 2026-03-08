import { useInfiniteQuery, useQuery } from "@tanstack/react-query";
import { queryKeys } from "./keys";
import { httpClient } from "../api";

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
