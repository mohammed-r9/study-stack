import z from "zod";

export const insertMaterialSchema = z.object({
	title: z.string().min(5)
})

export const insertCollectionSchema = z.object({
	title: z.string().min(5),
	description: z.string().min(12)
})

export const lectureFormSchema = z.object({
	lecture_title: z.string().min(1, "Required"),
	lecture_file: z
		.instanceof(File)
		.nullable()
		.refine(file => file !== null, { message: "Lecture file is required" })
		.refine(file => !file || file.type === "application/pdf", { message: "Only PDF allowed" })
})

export type insertMaterialInput = z.infer<typeof insertMaterialSchema>
export type InsetCollectionInput = z.infer<typeof insertMaterialSchema>
export type insertLectureInput = z.infer<typeof lectureFormSchema>
