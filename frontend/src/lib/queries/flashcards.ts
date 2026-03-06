import { useQuery } from "@tanstack/react-query";
import { queryKeys } from "./keys";
import { httpClient } from "../api";

export const useFlashcard = (cursor: number) =>
	useQuery({
		queryKey: [...queryKeys.flashcard, cursor],
		queryFn: () => httpClient.getOneFlashcard(),
		staleTime: 0,
	})


