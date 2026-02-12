import type { createLectureReq } from "./types";

export function getCSRFCookie(name = 'CSRF_token') {
	const value = `; ${document.cookie}`;
	const parts = value.split(`; ${name}=`);
	if (parts.length === 2) return parts.pop()?.split(';').shift();
	return null;
}

export function buildLectureFormData(values: createLectureReq) {
	const formData = new FormData()

	formData.append('material_id', values.material_id)
	formData.append('lecture_title', values.lecture_title)
	formData.append('lecture_file', values.lecture_file)

	return formData
}
