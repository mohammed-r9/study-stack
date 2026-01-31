import type { AxiosInstance } from "axios";
import axios from "axios";
import { API_URL } from "../const";
import type { LoginReq, RefreshRes, RegisterReq, User } from "./types";
import { getCSRFCookie } from "./utils";
import { useAuthStore } from "../store/auth";

class HttpClient {
	private api!: AxiosInstance

	constructor() {
		this.api = axios.create({
			baseURL: API_URL,
			withCredentials: true,
			timeout: 10000,
			headers: { "Content-Type": "application/json" },
		})

		// interceptors
		// attach the token to auhtenticate requests
		this.api.interceptors.request.use(
			(config) => {
				const token = useAuthStore.getState().accessToken
				if (token) {
					config.headers.Authorization = `Bearer ${token}`;
				}
				return config;
			},
			(error) => Promise.reject(error)
		);

	}
	public async login(body: LoginReq) {
		return this.api.post<RefreshRes>("/users/login", body)
	}
	public async register(body: RegisterReq) {
		return this.api.post("/users/register", body)
	}

	public async refreshToken() {
		return this.api.post<RefreshRes>("/users/refresh", {}, {
			headers: {
				'X-CSRF-Token': getCSRFCookie(),
			}
		})
	}

	public async logout() {
		return this.api.post("/users/logout")
	}

	public async getCurrentUser() {
		return this.api.get<User>("/users/me")
	}


}

export const httpClient = new HttpClient()
