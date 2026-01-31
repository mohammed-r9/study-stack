import type { MyRouterContext } from "@/routes/__root";
import { useAuthStore } from "./store/auth";
import { httpClient } from "./api";
import { queryKeys } from "./queries/keys";

const refreshToken = async (context: MyRouterContext) => {
	try {
		const newToken = await context.queryClient.fetchQuery({
			queryKey: queryKeys.refreshToken,
			queryFn: () => httpClient.refreshToken(),
			retry: false,
			staleTime: 1000 * 60 * 15,
		});
		useAuthStore.getState().setAccessToken(newToken.data.access_token ?? null);
	} catch (err) {
		console.error(err)
		useAuthStore.getState().setAccessToken(null);
	}
};

export const authLoader = async (context: MyRouterContext) => {
	await refreshToken(context);
	setInterval((context: MyRouterContext) => refreshToken(context), 1000 * 60 * 14.5);
};

