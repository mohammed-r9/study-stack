export const queryKeys = {
	currentUser: ["current-user"] as const,
	refreshToken: ["refresh-token"] as const,
	library: {
		all: ["library"] as const,
		collections: () =>
			[...queryKeys.library.all, "collections"] as const,
		collection: (collectionId: string) =>
			[...queryKeys.library.all, "collection", collectionId] as const,
		materials: (collectionId: string) =>
			[...queryKeys.library.all, "collection", collectionId, "materials"] as const,
		lectures: (materialId: string) =>
			[...queryKeys.library.all, "material", materialId, "lectures"] as const,
		update: (collectionId: string) =>
			[...queryKeys.library.all, "collection", collectionId, "update"] as const,
		signedURL: (lectureId: string) =>
			[...queryKeys.library.all, "lecture", lectureId, "signed-url"] as const,
	},
}
