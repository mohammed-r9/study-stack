import type { AxiosInstance } from "axios";
import axios from "axios";
import { API_URL } from "../const";
import type { Collection, CreateCollectionReq, CreateFlashCardBody, createLectureReq, CreateMaterialReq, Flashcard, GetAllFlashcardParams, GetAllFlashcardsRes, GetAllLecturesParams, GetAllLecturesRes, LoginReq, Material, RefreshRes, RegisterReq, SignedURL, UpdateCollectionReq, UpdateFlashcardParams, User, UserLibrary } from "./types";
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
		// attach the token to authenticate requests
		this.api.interceptors.request.use(
			async (config) => {
				if (config.url?.endsWith("/refresh")) return config;

				const token = useAuthStore.getState().accessToken;

				if (token) {
					const payload = JSON.parse(atob(token.split('.')[1]));
					const expiresIn = payload.exp * 1000 - Date.now();

					if (expiresIn < (60000 / 2)) {
						const tokenRes = await this.refreshToken();
						const newToken = tokenRes.data.access_token;
						useAuthStore.getState().setAccessToken(newToken);
						config.headers.Authorization = `Bearer ${newToken}`;
					} else {
						config.headers.Authorization = `Bearer ${token}`;
					}
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

	public async getSignedURL(id: string) {
		return this.api.get<SignedURL>(`/lectures/${id}/download-link`)
	}

	public async createFlashcard(body: CreateFlashCardBody) {
		return this.api.post("/flashcards", body)
	}

	public async getOneFlashcard() {
		return this.api.get<Flashcard>("/flashcards?mode=study")
	}

	public async getAllFlashcards(params: GetAllFlashcardParams) {
		return this.api.get<GetAllFlashcardsRes>(`/flashcards?mode=list&last_seen_flashcard_id=${params.last_seen_flashcard_id}`)
	}

	public async updateFlashcard(params: UpdateFlashcardParams) {
		return this.api.patch<Flashcard>(`/flashcards/${params.id}`, { back: params.back, front: params.front })
	}

	public async deleteFlashcard(id: string) {
		return this.api.delete<string>(`/flashcards/${id}`)
	}

}

export const httpClient = new HttpClient()
