import eslint from "@eslint/js";
import tseslint from "typescript-eslint";
import vuePlugin from "eslint-plugin-vue";
import globals from "globals";
import { globalIgnores } from "eslint/config";

export default [
  eslint.configs.recommended,
  ...tseslint.configs.recommended,
  ...vuePlugin.configs["flat/essential"],
  {
    languageOptions: {
      ecmaVersion: 2022,
      sourceType: "module",
      parser: vuePlugin.parser,
      parserOptions: {
        parser: tseslint.parser,
        extraFileExtensions: [".vue"],
        ecmaFeatures: {
          jsx: true,
        },
      },
      globals: {
        ...globals.node,
      },
    },
    rules: {
      "no-console": process.env.NODE_ENV === "production" ? "warn" : "off",
      "no-debugger": process.env.NODE_ENV === "production" ? "warn" : "off",
      "vue/no-deprecated-slot-attribute": "off",
      "@typescript-eslint/no-explicit-any": "off",
    },
    files: ["**/*.{js,ts,vue}"],
  },
  globalIgnores([
    ".DS_Store",
    "node_modules",
    "/coverage",
    "**/ios/",
    "**/android/",
    "**/dist/",
    "**/gen/",
    //# local env files
    ".env.local",
    ".env.*.local",
    //# Log files
    "npm-debug.log*",
    "yarn-debug.log*",
    "yarn-error.log*",
    "pnpm-debug.log*",
    //# Editor directories and files
    ".idea",
    ".vscode",
    "*.suo",
    "*.ntvs*",
    "*.njsproj",
    "*.sln",
    "*.sw?",
  ]),
];
