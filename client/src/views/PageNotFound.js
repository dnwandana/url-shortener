import { useHistory } from "react-router-dom"
import sorryImg from "../img/sorry.svg"

const PageNotFound = () => {
  let history = useHistory()

  const backToHomePage = () => history.push("/")

  return (
    <div className="min-h-screen px-4 flex flex-col justify-center items-center text-center bg-gray-100 space-y-5">
      <img className="w-1/2 md:w-1/4 lg:w-1/6" src={sorryImg} alt="Sorry" />
      <h1 className="text-gray-900 text-2xl md:text-3xl font-medium md:font-semibold">
        Can't find what you looking for
      </h1>
      <p className="text-gray-900 text-base md:text-lg font-normal">
        I think you've clicked on a bad link or entered an invalid URL
      </p>
      <button
        className="py-2 px-4 rounded-md text-base font-medium bg-indigo-700 text-indigo-200 hover:bg-indigo-500 hover:text-indigo-100 transition-colors duration-300 ease-in-out"
        onClick={backToHomePage}
      >
        Back to Homepage
      </button>
    </div>
  )
}

export default PageNotFound
