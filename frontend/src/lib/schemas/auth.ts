import { z } from 'zod';

export const loginSchema = z.object({
	email: z
		.email('Invalid email address'),
	password: z
		.string()
		.min(8, "Password must be at least 8 characters")
		.max(46, "Password is too long")
		.regex(/[A-Z]/, "Must contain at least one uppercase letter")
		.regex(/[a-z]/, "Must contain at least one lowercase letter")
		.regex(/[0-9]/, "Must contain at least one number"),
});

export const registerSchema = z
	.object({
		email: z.email("Invalid email address"),
		name: z.string().max(32, "Name is too long"),
		password: z
			.string()
			.min(8, "Password must be at least 8 characters")
			.max(46, "Password is too long")
			.regex(/[A-Z]/, "Must contain at least one uppercase letter")
			.regex(/[a-z]/, "Must contain at least one lowercase letter")
			.regex(/[0-9]/, "Must contain at least one number"),
		confirm_password: z.string(),
	})
	.refine((data) => data.password === data.confirm_password, {
		message: "Passwords don't match",
		path: ["confirmPassword"],
	});

export type LoginInputType = z.infer<typeof loginSchema>;
export type RegisterInputType = z.infer<typeof registerSchema>;
