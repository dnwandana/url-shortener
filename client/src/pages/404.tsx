import Link from "next/link"
import Metadata from "../components/Metadata"

const ErrorPage = () => {
  return (
    <>
      <Metadata title="URL Shortener - Page Not Found" />
      <div className="flex flex-col justify-center mx-auto min-h-screen bg-gray-200">
        <div className="mx-auto max-w-screen-md">
          <div className="flex flex-col p-4 md:flex-row md:items-center">
            <div className="flex flex-row px-2 text-center space-x-1 md:flex-col md:px-4 md:space-x-0">
              <h1 className="text-indigo-600 text-2xl font-bold md:text-3xl">
                404
              </h1>
              <h1 className="text-indigo-600 text-2xl font-bold md:text-3xl">
                ERROR
              </h1>
            </div>
            <div className="flex flex-col mt-4 px-2 border-gray-400 md:mt-0 md:px-4 md:border-l">
              <h1 className="text-gray-900 text-2xl font-bold md:text-3xl">
                Can't find what you looking for
              </h1>
              <p className="mt-1 text-gray-800 text-sm md:text-base">
                It looks like you've clicked on a bad link or entered an invalid
                URL
              </p>
            </div>
          </div>
          <div className="mt-6 px-6 md:flex md:items-center md:justify-center md:mt-4">
            <Link href="/">
              <a className="px-4 py-2 text-white text-sm font-medium tracking-wide bg-indigo-600 hover:bg-indigo-700 border border-transparent rounded-md focus:outline-none shadow-sm focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2">
                Shorten URL
              </a>
            </Link>
          </div>
        </div>
      </div>
    </>
  )
}

export default ErrorPage
