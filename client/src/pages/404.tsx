import Link from "next/link"
import Metadata from "../components/Metadata"

const Custom404Error = () => {
  return (
    <>
      <Metadata title="404 | Page Not Found" />
      <div className="min-h-screen px-4 flex flex-col justify-center items-center text-center bg-gray-100 space-y-5">
        <img
          className="w-1/2 md:w-1/4 lg:w-1/6"
          src="/assets/svg/sorry.svg"
          alt="Sorry"
        />
        <h1 className="text-gray-900 text-2xl md:text-3xl font-medium md:font-semibold">
          Can't find what you looking for
        </h1>
        <p className="text-gray-900 text-base md:text-lg font-normal">
          I think you've clicked on a bad link or entered an invalid URL
        </p>
        <Link href="/">
          <a className="py-2 px-4 rounded-md text-base font-medium bg-indigo-600 text-indigo-100 hover:bg-indigo-700 hover:text-indigo-50 focus:outline-none focus:ring-2 focus:ring-indigo-600 focus:ring-opacity-50 transition-colors duration-300 ease-in-out">
            Back to Homepage
          </a>
        </Link>
      </div>
    </>
  )
}

export default Custom404Error
