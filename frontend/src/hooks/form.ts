import AppField from "@/components/form/app-field";
import { createFormHook } from "@tanstack/react-form";
import { fieldContext, formContext } from "./form-context";

export const { useAppForm, withForm, withFieldGroup } = createFormHook({
	fieldContext: fieldContext,
	formContext: formContext,
	fieldComponents: {
		AppField: AppField,
	},
	formComponents: {},
});
