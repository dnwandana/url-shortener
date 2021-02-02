import { useState } from "react"
import { CopyToClipboard } from "react-copy-to-clipboard"
import PropTypes from "prop-types"

const FormResult = ({ longUrl, shortUrl }) => {
  const [copyStatus, setCopyStatus] = useState(false)

  const copyToClipboard = () => {
    setCopyStatus(true)
    setTimeout(() => {
      setCopyStatus(false)
    }, 1000)
  }

  return (
    <ul className="mt-4 py-4 px-3 md:px-4 bg-gray-100 rounded-md block md:flex md:justify-between md:items-center space-y-2 md:space-y-0 transition duration-1000 ease-in-out">
      <li className="w-full md:w-1/2 truncate">
        <span className="text-base text-gray-900">{longUrl}</span>
      </li>
      <div className="block md:flex md:items-center space-y-3 md:space-y-0 md:space-x-5">
        <li>
          <a
            className="text-base text-indigo-600"
            href={shortUrl}
            target="_blank"
            rel="noreferrer"
          >
            {shortUrl}
          </a>
        </li>
        <li>
          <CopyToClipboard text={shortUrl} onCopy={copyToClipboard}>
            <button
              className={`w-full py-2 px-4 md:px-6 rounded-md text-base transition-colors duration-500 ease-in-out focus:outline-none focus:ring ${
                copyStatus
                  ? "bg-green-500 text-green-100 focus:ring-green-500"
                  : "bg-indigo-100 text-indigo-600 hover:bg-indigo-200 hover:text-indigo-700 focus:ring-indigo-200"
              }`}
            >
              {copyStatus ? "Copied!" : "Copy"}
            </button>
          </CopyToClipboard>
        </li>
      </div>
    </ul>
  )
}

FormResult.propTypes = {
  longUrl: PropTypes.string.isRequired,
  shortUrl: PropTypes.string.isRequired
}

export default FormResult
