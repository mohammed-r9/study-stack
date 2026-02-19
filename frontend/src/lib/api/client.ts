import type { AxiosInstance } from "axios";
import axios from "axios";
import { API_URL } from "../const";
import type { Collection, CreateCollectionReq, createLectureReq, CreateMaterialReq, GetAllLecturesParams, GetAllLecturesRes, LoginReq, Material, RefreshRes, RegisterReq, UpdateCollectionReq, User, UserLibrary } from "./types";
import { buildLectureFormData, getCSRFCookie } from "./utils";
import { useAuthStore } from "../store/auth";
import { toast } from "sonner";

class HttpClient {
	private api!: AxiosInstance

	constructor() {
		this.api = axios.create({
			baseURL: API_URL,
			withCredentials: true,
			timeout: 10000,
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

	public async getLibrary() {
		return this.api.get<UserLibrary>("/users/me/library")
	}

	public async getMaterialsByCollection(collectionID: string, archived: boolean) {
		return this.api.get<Material[]>(`/materials/?collection_id=${collectionID}&archived=${archived}`)
	}
	public async getAllCollections(getArchived: boolean) {
		return this.api.get<Collection[]>(`/collections?archived=${getArchived}`)
	}
	public async newMaterial(body: CreateMaterialReq) {
		return this.api.post("/materials", body)
	}
	public async newCollection(body: CreateCollectionReq) {
		return this.api.post("/collections", body)
	}

	public async updateCollection(body: UpdateCollectionReq, id: string) {
		if (!body.new_description && !body.new_title) {
			toast.error("Please enter some values first")
			return
		}
		const finalReq = {
			new_title: body.new_title || null,
			new_description: body.new_description || null
		}
		return this.api.patch(`/collections/${id}`, finalReq)
	}

	public async createLecture(body: createLectureReq) {
		const formData = buildLectureFormData(body)
		return this.api.post("/lectures", formData, {
			headers: {
				"Content-Type": "multipart/form-data",
			}
		})
	}

	public async getAllLectures(params: GetAllLecturesParams) {
		return this.api.get<GetAllLecturesRes>(`/lectures?m_id=${params.material_id}&last_seen=${params.last_seen_id}`)
	}


}

export const httpClient = new HttpClient()
