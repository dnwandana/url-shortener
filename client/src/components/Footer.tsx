const Footer = (): JSX.Element => {
  return (
    <div className="flex justify-between mt-2">
      <span className="text-gray-900 text-sm font-medium">
        <a
          href="https://dhanz.xyz/api-spec"
          target="_blank"
          rel="noopener noreferrer"
          className="hover:text-indigo-600">
          OpenAPI Specification
        </a>
      </span>
      <span className="text-gray-900 text-sm font-medium">
        By{" "}
        <a
          href="https://github.com/dnwandana"
          target="_blank"
          rel="noopener noreferrer"
          className="hover:text-indigo-600">
          Wandana
        </a>
      </span>
    </div>
  )
}

export default Footer
