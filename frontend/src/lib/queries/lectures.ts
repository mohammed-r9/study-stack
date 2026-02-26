import { useInfiniteQuery, useMutation, useQuery, useQueryClient } from "@tanstack/react-query"
import { httpClient } from "../api"
import type { createLectureReq } from "../api/types"
import { toast } from "sonner"
import { queryKeys } from "./keys"
import { useRouter } from "@tanstack/react-router"

export function useCreateLecture(materialID: string, title: string) {
	const router = useRouter()
	const query = useQueryClient()
	return useMutation({
		mutationFn: async (body: createLectureReq) => {
			const res = await httpClient.createLecture(body)
			return res.data
		},
		onSuccess: () => {
			toast.success('Lecture uploaded successfully')
			query.invalidateQueries({
				queryKey: queryKeys.library.lectures(materialID),
				refetchType: "all"
			})
			router.navigate({ to: "/materials/$id", params: { id: materialID }, search: { title: title } })
		},
		onError: (err: any) => {
			toast.error(err?.response?.data?.message || err?.message || 'Failed to upload lecture')
		},
	})
}


export const useInfiniteLectures = (materialId: string) => {
	return useInfiniteQuery({
		queryKey: queryKeys.library.lectures(materialId),

		queryFn: async ({ pageParam }) => {
			const response = await httpClient.getAllLectures({
				material_id: materialId,
				last_seen_id: pageParam,
			});
			return response.data;
		},

		initialPageParam: "",

		getNextPageParam: (lastPage) => {
			if (!lastPage.has_next_page || lastPage.lectures.length === 0) {
				return undefined;
			}
			return lastPage.lectures[lastPage.lectures.length - 1].id;
		},

		staleTime: Infinity
	});
};

export function useSignedURL(lectureId: string) {
	return useQuery({
		queryKey: queryKeys.library.signedURL(lectureId),
		queryFn: () => httpClient.getSignedURL(lectureId),
		staleTime: 5 * 1000 * 60
	})
}
