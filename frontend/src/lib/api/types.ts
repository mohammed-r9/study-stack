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
	updated_at: string
	created_at: string
	verified_at: string | null
}

export type Collection = {
	id: string
	title: string
	description: string
}

export type Material = {
	id: string
	collection_id: string
	title: string
}
