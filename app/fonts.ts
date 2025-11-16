import localFont from "next/font/local";
import { Noto_Sans_JP } from "next/font/google";

export const noto = Noto_Sans_JP({
  subsets: ["latin"],
  weight: ["400", "700"],
  variable: "--font-noto",
})

export const genei = localFont({
  src: [
    {
      path: "./fonts/GenEiWebHonmon-R.woff2",
      weight: "400",
      style: "normal",
    },
    {
      path: "./fonts/GenEiWebHonmon-B.woff2",
      weight: "700",
      style: "normal",
    },
  ],
  variable: "--font-genei",
})
