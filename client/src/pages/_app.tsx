import type { AppProps } from "next/app"
import "../styles/globals.css"

function urlShortener({ Component, pageProps }: AppProps) {
  return <Component {...pageProps} />
}

export default urlShortener
