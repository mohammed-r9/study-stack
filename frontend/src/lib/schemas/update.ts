import { z } from 'zod';


export const updateCollectionSchema = z.object({
	new_title: z
		.string()
		.refine((val) => val === "" || val.length >= 6, {
			message: "Title should be at least 6 characters",
		}),
	new_description: z
		.string()
		.refine((val) => val === "" || val.length >= 24, {
			message: "Description should be at least 24 characters",
		}),
})


export const updateMaterialSchema = z.object({
	new_title: z
		.string()
		.refine((val) => val === "" || val.length >= 6, {
			message: "Title should be at least 6 characters",
		}),
	new_description: z
		.string()
		.refine((val) => val === "" || val.length >= 24, {
			message: "Description should be at least 24 characters",
		}),
})

export type updateCollectionInput = z.infer<typeof updateCollectionSchema>
export type updateMaterialType = z.infer<typeof updateMaterialSchema>
