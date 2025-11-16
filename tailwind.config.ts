import type { Config } from "tailwindcss";

export default {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        background: "var(--background)",
        foreground: "var(--foreground)",
      },
    },
    fontFamily: {
      genei: ['var(--font-genei)', 'var(--font-noto)', "Hiragino Sans", "Hiragino Kaku Gothic ProN", "Meiryo", "sans-serif"]
    }
  },
  plugins: [],
} satisfies Config;
