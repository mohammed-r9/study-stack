import z from "zod";

export const insertMaterialSchema = z.object({
	title: z.string().min(5)
})
export const insertCollectionSchema = z.object({
	title: z.string().min(5),
	description: z.string().min(12)
})

export type insertMaterialInput = z.infer<typeof insertMaterialSchema>
export type InsetCollectionInput = z.infer<typeof insertMaterialSchema>
