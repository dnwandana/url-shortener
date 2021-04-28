import { useState } from "react"
import Link from "next/link"

const Navbar = (): JSX.Element => {
  const [isMenuOpen, setIsMenuOpen] = useState<boolean>(false)

  const toggleMenu = () => {
    setIsMenuOpen(!isMenuOpen)
  }

  return (
    <nav className="bg-gray-100">
      <div className="container max-w-6xl mx-auto px-4 pt-5 pb-6 md:flex md:justify-between md:items-center">
        <div className="flex items-center justify-between">
          <Link href="/">
            <a className="text-gradient bg-gradient-to-r from-indigo-600 to-gray-900 font-bold text-2xl md:text-3xl hover:text-gray-700">
              URL Shortener
            </a>
          </Link>

          {/* <!-- Mobile menu button --> */}
          <div className="flex md:hidden">
            <button
              type="button"
              className="text-gray-800 hover:text-gray-900"
              aria-label="toggle menu"
              onClick={toggleMenu}>
              {isMenuOpen ? (
                <svg
                  className="w-6 h-6 fill-current"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 24 24"
                  stroke="currentColor">
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M6 18L18 6M6 6l12 12"
                  />
                </svg>
              ) : (
                <svg
                  className="w-6 h-6 fill-current"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 24 24"
                  stroke="currentColor">
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M4 6h16M4 12h16M4 18h16"
                  />
                </svg>
              )}
            </button>
          </div>
        </div>

        {/* <!-- Mobile Menu open: "block", Menu closed: "hidden" --> */}
        <div
          className={`items-center md:flex ${isMenuOpen ? "block" : "hidden"}`}>
          <div className="flex flex-col md:items-center mt-3 md:mt-0 md:flex-row md:mx-6 space-y-3 space-x-0 md:space-y-0 md:space-x-5">
            <Link href="/sign-in">
              <a className="font-medium text-base md:text-lg lg:text-xl text-gray-800 hover:text-indigo-500">
                Sign In
              </a>
            </Link>
            <Link href="/sign-up">
              <a className="font-medium text-base md:text-lg lg:text-xl text-indigo-500 hover:text-gray-800">
                Sign Up
              </a>
            </Link>
          </div>
        </div>
      </div>
    </nav>
  )
}

export default Navbar
