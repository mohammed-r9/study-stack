export const queryKeys = {
	currentUser: ["current-user"],
	refreshToken: ["refresh-token"],
	library: {
		all: ["library"] as const,
		collections: () => [...queryKeys.library.all, "collections"] as const,
		collection: (collectionId: string) =>
			[...queryKeys.library.all, "collection", collectionId] as const,
		materials: (collectionId: string) =>
			[...queryKeys.library.all, "collection", collectionId, "materials"] as const,

		update: (collectionId: string) =>
			[...queryKeys.library.all, "collection", collectionId, "update"] as const,
		lectures: (materialId: string) =>
			[...queryKeys.library.all, "material", materialId, "lectures"] as const,
	},
}
