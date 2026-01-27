import { useAuthStore } from "../store/auth";

export type AuthContext = {
    isAuthenticated: boolean;
    invalidate: () => Promise<void> | void;
}

export const authContext: AuthContext = {
    get isAuthenticated() {
        return !!useAuthStore.getState().accessToken;
    },
    invalidate: () => { }
}
