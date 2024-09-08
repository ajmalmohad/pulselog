import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import { resolve } from "node:path";

const root = resolve(__dirname, "src");

export default defineConfig(() => {
	return {
		build: {
			outDir: "build",
		},
		plugins: [react()],
		resolve: {
			alias: {
				"@app": resolve(root, "."),
				"@components": resolve(root, "components"),
				"@pages": resolve(root, "pages"),
			},
		},
	};
});