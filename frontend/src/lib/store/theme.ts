import { create } from "zustand"

type ThemeStore = {
	theme: "dark" | "light"
	toggleTheme: () => void
}
const getInitialTheme = (): "dark" | "light" => {
	const value = localStorage.getItem("vite-ui-theme")

	if (value === "dark" || value === "light") {
		return value
	}

	return "dark"
}

export const useThemeStore = create<ThemeStore>((set) => ({
	theme: getInitialTheme(),
	toggleTheme: () =>
		set((state) => {
			const newTheme = state.theme === "dark" ? "light" : "dark"
			localStorage.setItem("vite-ui-theme", newTheme)
			return ({
				theme: newTheme,
			})
		}),
}))


