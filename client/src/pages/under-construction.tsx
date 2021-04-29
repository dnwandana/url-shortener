import Metadata from "../components/Metadata"
import Link from "next/link"

const underConstruction = () => {
  return (
    <>
      <Metadata title="Under Construction" />
      <div className="min-h-screen px-4 flex flex-col justify-center items-center text-center bg-gray-100 space-y-5">
        <img
          className="w-1/2 md:w-1/4 lg:w-1/6"
          src="/assets/under-construction.svg"
          alt="Under construction"
        />
        <h1 className="text-gray-900 font-bold text-2xl md:text-3xl lg:text-4xl">
          Under construction
        </h1>
        <p className="text-gray-800 font-normal text-base md:text-lg lg:text-xl">
          Sorry, this website is not finished yet, but you can still shorten the
          URL.
        </p>
        <Link href="/">
          <a className="py-2 px-4 rounded-md font-medium text-base text-center transition-colors duration-300 ease-in-out bg-indigo-700 text-indigo-100 hover:bg-indigo-600 hover:text-indigo-50 focus:outline-none focus:ring focus:ring-indigo-300 tracking">
            Back to Homepage
          </a>
        </Link>
      </div>
    </>
  )
}

export default underConstruction
