export type LoginReq = {
	email: string
	password: string
}
export type RegisterReq = {
	name: string
	email: string
	password: string
	confirm_password: string
}

export type RefreshRes = {
	access_token: string
}

export type User = {
	id: string
	name: string
	email: string
	updated_at: Date
	created_at: Date
	verified_at: string | null
}

export type Collection = {
	id: string
	title: string
	description: string
	updated_at: Date
	created_at: Date
}

export type Material = {
	id: string
	collection_id: string
	title: string
	created_at: Date
	updated_at: Date
}

export type UserLibrary = {
	id: string
	title: string
	description: string
	materials: {
		id: string
		title: string
		description: string
	}[]
}[]

export type UpdateCollectionReq = {
	new_title: string
	new_description: string
}


export type CreateMaterialReq = {
	collection_id: string
	title: string
}
export type CreateCollectionReq = {
	title: string
	description: string
}

export type createLectureReq = {
	lecture_title: string
	lecture_file: File
	material_id: string
}

export type GetAllLecturesParams = {
	material_id: string
	last_seen_id: string
}

export type Lecture = {
	id: string
	material_id: string
	title: string
	file_size: string
	created_at: Date
	updated_at: Date
	archived_at: Date
}

export type GetAllLecturesRes = {
	lectures: Lecture[]
	has_next_page: boolean
}

export type SignedURL = {
	url: string
}
