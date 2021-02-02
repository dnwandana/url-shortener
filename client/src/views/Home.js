import Form from "../components/Form"
import headerLogo from "../img/header.svg"
import easyToUseImg from "../img/1-easy.svg"
import shortenImg from "../img/2-shorten.svg"
import shareImg from "../img/3-share.svg"

const Home = () => {
  return (
    <>
      <div className="bg-gray-100">
        <div className="container max-w-6xl mx-auto px-4 py-8 flex flex-col-reverse md:flex-row md:justify-between md:items-center bg-gray-100">
          <div className="w-full md:w-3/5 mt-6 md:mt-0 text-left space-y-2 md:space-y-0">
            <h1 className="text-gradient bg-gradient-to-r from-indigo-600 to-gray-900 font-bold text-xl md:text-3xl lg:text-5xl">
              Shorten and Share.
            </h1>
            <p className="text-gray-900 font-medium text-base md:text-xl lg:text-3xl">
              Shorten a long URL and share as a powerful.
            </p>
          </div>
          <img
            className="w-1/2 md:w-2/5 mx-auto"
            src={headerLogo}
            alt="Social media engagement"
          />
        </div>
      </div>
      <Form />
      <div className="bg-gray-100">
        <div className="container max-w-6xl mx-auto px-4 py-8 bg-gray-100">
          <h1 className="mb-6 text-center text-gray-900 font-bold text-xl md:text-3xl">
            Free URL Shortening Services
          </h1>
          <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
            <div className="p-4 bg-gray-50 border rounded space-y-3 shadow-sm hover:shadow-lg">
              <img
                className="w-1/3 mx-auto"
                src={easyToUseImg}
                alt="Easy to use"
              />
              <h3 className="font-semibold text-xl text-gray-900 text-center">
                Easy
              </h3>
              <p className="text-base text-gray-900 text-center">
                Easy and fast, just enter the long link to get your shortened
                link.
              </p>
            </div>
            <div className="p-4 bg-gray-50 border rounded space-y-3 shadow-sm hover:shadow-lg">
              <img
                className="w-1/3 mx-auto"
                src={shortenImg}
                alt="Shorten URL"
              />
              <h3 className="font-semibold text-xl text-gray-900 text-center">
                Shortened
              </h3>
              <p className="text-base text-gray-900 text-center">
                Any link, no matter what size, always shortens your link.
              </p>
            </div>
            <div className="p-4 bg-gray-50 border rounded space-y-3 shadow-sm hover:shadow-lg">
              <img className="w-1/3 mx-auto" src={shareImg} alt="Share URL" />
              <h3 className="font-semibold text-xl text-gray-900 text-center">
                Share
              </h3>
              <p className="text-base text-gray-900 text-center">
                Share your short links across social media or anywhere.
              </p>
            </div>
          </div>
        </div>
      </div>
      <div className="bg-gray-900">
        <div className="container max-w-6xl mx-auto py-6 px-4 bg-gray-900 inline-block md:flex md:justify-between md:items-center space-y-3 md:space-y-0">
          <p className="text-sm font-medium text-gray-300">
            Some icons made by{" "}
            <a
              className="text-indigo-500 hover:text-indigo-400"
              href="https://www.flaticon.com/authors/prosymbols"
              title="Prosymbols"
              target="_blank"
              rel="noreferrer"
            >
              Prosymbols
            </a>{" "}
            and{" "}
            <a
              className="text-indigo-500 hover:text-indigo-400"
              href="https://www.freepik.com"
              title="Freepik"
              target="_blank"
              rel="noreferrer"
            >
              Freepik
            </a>{" "}
            from{" "}
            <a
              className="text-indigo-500 hover:text-indigo-400"
              href="https://www.flaticon.com/"
              title="Flaticon"
              target="_blank"
              rel="noreferrer"
            >
              www.flaticon.com
            </a>
          </p>
          <p className="text-sm font-medium text-gray-300">
            By{" "}
            <a
              className="text-indigo-500 hover:text-indigo-400"
              href="https://github.com/dnwandana/url-shortener"
              title="Wandana"
              target="_blank"
              rel="noreferrer"
            >
              Wandana
            </a>
          </p>
        </div>
      </div>
    </>
  )
}

export default Home
