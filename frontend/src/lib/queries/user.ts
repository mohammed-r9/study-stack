import { queryOptions, } from "@tanstack/react-query";
import { queryKeys } from "./keys";
import { httpClient } from "../api";


export const createUserQueryOptions = () => {
	return queryOptions({
		queryKey: queryKeys.currentUser,
		queryFn: httpClient.getCurrentUser,
		staleTime: Infinity
	})
}


