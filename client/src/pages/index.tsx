import type { GlobalState } from "../@types/GlobalState"
import Metadata from "../components/Metadata"
import dynamic from "next/dynamic"
import store from "../store"

const Alert = dynamic(() => import("../components/Alert"))
const ShortenForm = dynamic(() => import("../components/Form/ShortenForm"))
const ResultForm = dynamic(() => import("../components/Form/ResultForm"))
const Footer = dynamic(() => import("../components/Footer"))

const index = () => {
  const alertMessage = store((state: GlobalState) => state.alertMessage)
  const shortURL = store((state: GlobalState) => state.shortURL)
  const longURL = store((state: GlobalState) => state.longURL)

  return (
    <>
      <Metadata title="URL Shortener - Short URLs &amp; Custom Free Link Shortener" />
      <div className="flex flex-col justify-center mx-auto min-h-screen bg-gray-200">
        <div className="p-4 max-w-screen-md md:mx-auto">
          {/* ALERT */}
          {alertMessage !== "" && (
            <Alert
              AlertClass="block w-full p-4 font-medium text-base text-center bg-red-300 text-red-800 shadow"
              Message={alertMessage}
            />
          )}
          {/* FORM */}
          <ShortenForm />
          {/* RESULT FORM */}
          {shortURL !== "" && longURL !== "" && (
            <ResultForm longURL={longURL} shortURL={shortURL} />
          )}
          {/* FOOTER */}
          <Footer />
        </div>
      </div>
    </>
  )
}

export default index
