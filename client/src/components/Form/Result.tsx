import { useState } from "react"
import CopyToClipboard from "react-copy-to-clipboard"

type ResultProps = {
  longURL: string
  shortURL: string
}

const Result = ({ longURL, shortURL }: ResultProps): JSX.Element => {
  const [copyStatus, setCopyStatus] = useState<boolean>(false)

  const copyToClipboard = () => {
    setCopyStatus(true)
    setTimeout(() => {
      setCopyStatus(false)
    }, 1000)
  }

  return (
    <div className="mt-4 py-4 px-3 md:px-4 bg-gray-100 rounded-md block md:flex md:justify-between md:items-center space-y-2 md:space-y-0 transition duration-1000 ease-in-out">
      <div className="w-full md:w-1/2 truncate">
        <span className="font-medium text-base text-gray-700">{longURL}</span>
      </div>
      <div className="block md:flex md:items-center space-y-3 md:space-y-0 md:space-x-5">
        <div>
          <a
            className="font-medium text-base text-indigo-600"
            href={shortURL}
            target="_blank"
            rel="noreferrer"
          >
            {shortURL}
          </a>
        </div>
        <div className="md:w-24">
          <CopyToClipboard text={shortURL} onCopy={copyToClipboard}>
            <button
              className={`w-full py-2 px-4 md:px-6 rounded-md font-medium text-base transition-colors duration-300 ease-in-out focus:outline-none focus:ring ${
                copyStatus
                  ? "bg-green-500 text-green-100 focus:ring-green-500"
                  : "bg-indigo-200 text-indigo-700 hover:bg-indigo-300 hover:text-indigo-900 focus:ring-indigo-300"
              }`}
            >
              {copyStatus ? "Copied!" : "Copy"}
            </button>
          </CopyToClipboard>
        </div>
      </div>
    </div>
  )
}

export default Result
