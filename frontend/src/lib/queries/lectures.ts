import { useMutation } from "@tanstack/react-query"
import { httpClient } from "../api"
import type { createLectureReq } from "../api/types"
import { toast } from "sonner"

export function useCreateLecture() {
	return useMutation({
		mutationFn: async (body: createLectureReq) => {
			const res = await httpClient.createLecture(body)
			return res.data
		},
		onSuccess: () => {
			toast.success('Lecture uploaded successfully')
		},
		onError: (err: any) => {
			toast.error(err?.response?.data?.message || err?.message || 'Failed to upload lecture')
		},
	})
}
