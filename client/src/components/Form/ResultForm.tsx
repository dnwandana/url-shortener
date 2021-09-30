import CopyToClipboard from "react-copy-to-clipboard"
import { useState } from "react"

type ResultProps = {
  longURL: string
  shortURL: string
}

const ResultForm = ({ longURL, shortURL }: ResultProps): JSX.Element => {
  const [copyStatus, setCopyStatus] = useState<boolean>(false)

  const copyToClipboard = () => {
    setCopyStatus(true)
    setTimeout(() => {
      setCopyStatus(false)
    }, 1000)
  }

  return (
    <div className="block p-4 bg-white shadow space-y-2 md:flex md:items-center md:justify-between md:space-x-5 md:space-y-0">
      {/* LONG-URL */}
      <div className="w-full truncate md:w-2/5">
        <span className="text-gray-500 text-sm font-medium">{longURL}</span>
      </div>
      <div className="block space-y-4 md:flex md:items-center md:justify-end md:w-2/4 md:text-right md:space-x-5 md:space-y-0">
        {/* SHORT-URL */}
        <div className="w-full truncate">
          <a
            target="_blank"
            rel="noopener noreferrer"
            href={shortURL}
            title={`Shortened URL for ${longURL}`}
            className="text-indigo-600 text-sm font-medium">
            {shortURL}
          </a>
        </div>
        {/* COPY BUTTON */}
        <CopyToClipboard text={shortURL} onCopy={copyToClipboard}>
          <button
            className={`py-2 px-4 w-full rounded-md font-medium text-sm transition-colors duration-300 ease-in-out focus:outline-none focus:ring-2 border border-transparent shadow-sm focus:ring-offset-2 md:w-28 ${
              copyStatus
                ? "bg-green-500 text-green-100 focus:ring-green-500"
                : "bg-indigo-200 text-indigo-700 hover:bg-indigo-300 hover:text-indigo-900 focus:ring-indigo-300"
            }`}>
            {copyStatus ? "Copied!" : "Copy"}
          </button>
        </CopyToClipboard>
      </div>
    </div>
  )
}

export default ResultForm
