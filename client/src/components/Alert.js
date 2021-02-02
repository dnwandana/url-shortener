import PropTypes from "prop-types"

const Alert = ({ errorMessage }) => {
  return (
    <div className="mt-4 block w-full py-4 bg-red-300 rounded-md text-red-800 text-center">
      {errorMessage}
    </div>
  )
}

Alert.propTypes = {
  errorMessage: PropTypes.string.isRequired
}

export default Alert
