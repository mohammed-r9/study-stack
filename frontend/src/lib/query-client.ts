import { QueryClient } from '@tanstack/react-query';

export const queryClient = new QueryClient({
	defaultOptions: {
		queries: {
			retry: (failureCount, error: any) => {
				if (error.response?.status === 401 && failureCount < 2) {
					return true;
				}
				return false;
			},
			retryDelay: 0
		},
	},
});
