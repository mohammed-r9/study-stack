import { useInfiniteQuery, useMutation, useQuery, useQueryClient } from "@tanstack/react-query"
import { httpClient } from "../api"
import type { createLectureReq, Flashcard, UpdateFlashcardParams } from "../api/types"
import { toast } from "sonner"
import { queryKeys } from "./keys"
import { useRouter } from "@tanstack/react-router"

export function useCreateLecture(materialID: string, title: string) {
	const router = useRouter()
	const query = useQueryClient()
	return useMutation({
		mutationFn: async (body: createLectureReq) => {
			const res = await httpClient.createLecture(body)
			return res.data.lecture
		},
		onSuccess: (newLecture) => {
			query.setQueryData(
				queryKeys.library.lectures(materialID),
				(old: any) => {
					const newPage = { lectures: [newLecture], has_next_page: false };
					console.log(old)

					if (!old) {
						return {
							pages: [newPage],
							pageParams: [""],
						};
					}

					return {
						...old,
						pages: old.pages.map((page: any, index: number) =>
							index === 0
								? { ...page, lectures: [newLecture, ...(page.lectures || [])] }
								: page
						),
					};
				}
			);

			router.navigate({
				to: "/materials/$id",
				params: { id: materialID },
				search: { title: title }
			});
			toast.success('Lecture uploaded successfully');
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


