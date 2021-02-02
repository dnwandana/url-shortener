import { useState } from "react"
import loadable from "@loadable/component"
import { useFormik } from "formik"
import { object, string } from "yup"

const FormResult = loadable(() => import("./FormResult"))
const Alert = loadable(() => import("./Alert"))

const Form = () => {
  const [isSuccess, setIsSuccess] = useState(false)
  const [longUrl, setLongUrl] = useState("")
  const [shortUrl, setShortUrl] = useState("")

  const formik = useFormik({
    initialValues: {
      url: ""
    },
    validationSchema: object({
      url: string()
        .matches(
          /((https?):\/\/)?(www.)?[a-z0-9-]+(\.[a-z]{2,}){1,3}(#?\/?[a-zA-Z0-9#-]+)*\/?(\?[a-zA-Z0-9-_]+=[a-zA-Z0-9-%]+&?)?$/,
          "Please, provide a valid url."
        )
        .required("Please, provide a valid url.")
    }),
    onSubmit: async (values, { resetForm }) => {
      const res = await fetch(process.env.REACT_APP_API_ENDPOINT, {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify(values)
      })
      const { data } = await res.json()

      if (res.status === 201) {
        setIsSuccess(true)
        setShortUrl(data.shortUrl)
        setLongUrl(data.longUrl)
      }
      resetForm()
    }
  })

  return (
    <div className="bg-gray-900">
      <div className="container max-w-6xl mx-auto px-4 py-12 bg-gray-900">
        <form className="block md:flex relative" onSubmit={formik.handleSubmit}>
          <input
            id="url"
            name="url"
            className="block md:flex w-full md:w-3/4 rounded-md text-base text-indigo-600 border-indigo-300 shadow-sm focus:border-indigo-400 focus:ring focus:ring-indigo-400 focus:ring-opacity-50"
            type="text"
            placeholder="Shorten your link"
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            value={formik.values.url}
          />
          <button
            className="mt-3 ml-0 md:mt-0 md:ml-4 w-full md:w-1/4 px-2 py-3 rounded-md text-base tracking-wide transition-colors duration-300 ease-in-out bg-indigo-700 text-indigo-200 hover:bg-indigo-500 hover:text-indigo-100 focus:outline-none focus:ring focus:ring-indigo-300"
            type="submit"
          >
            Shorten
          </button>
        </form>
        {formik.touched.url && formik.errors.url ? (
          <Alert errorMessage={formik.errors.url} />
        ) : null}
        {isSuccess && <FormResult longUrl={longUrl} shortUrl={shortUrl} />}
      </div>
    </div>
  )
}

export default Form
