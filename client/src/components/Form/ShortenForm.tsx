import type { ResponseData, ResponseMessage } from "../../@types/APIResponse"
import * as yup from "yup"
import axios, { AxiosError } from "axios"
import store from "../../store"
import { GlobalState } from "../../@types/GlobalState"
import { useForm } from "react-hook-form"
import { useState } from "react"
import { yupResolver } from "@hookform/resolvers/yup"

const ShortenForm = (): JSX.Element => {
  type FormInput = {
    url: string
    id: string
    ttl: string
  }

  const formValidationSchema = yup.object().shape({
    url: yup
      .string()
      .url("Please, provide a valid url")
      .required("Please, provide a valid url"),
    id: yup
      .string()
      .test(
        "min-check",
        "At least 3 characters",
        (id) => id.length > 2 || id.length == 0
      ),
  })

  const { register, handleSubmit, errors } = useForm<FormInput>({
    resolver: yupResolver(formValidationSchema),
  })

  const [isLoading, setIsLoading] = useState(false)
  const setAlertMessage = store((state: GlobalState) => state.setAlertMessage)
  const setSuccess = store((state: GlobalState) => state.setSuccess)
  const setLongURL = store((state: GlobalState) => state.setLongURL)
  const setShortURL = store((state: GlobalState) => state.setShortURL)

  const shortenURL = async (formData: FormInput) => {
    try {
      setIsLoading(true)
      const domain = process.env.NEXT_PUBLIC_API_ENDPOINT
      const response = await axios.post(domain, formData)
      const data: ResponseData = response.data
      setSuccess(true)
      setAlertMessage("")
      setLongURL(data.data.long_url)
      setShortURL(data.data.short_url)
      setIsLoading(false)
    } catch (err) {
      const { response }: AxiosError<ResponseMessage> = err
      setSuccess(false)
      setAlertMessage(response.data.message)
      setLongURL("")
      setShortURL("")
      setIsLoading(false)
    }
  }

  return (
    <>
      <form
        className="mt-4 w-full bg-white shadow"
        onSubmit={handleSubmit(shortenURL)}>
        <div className="grid gap-4 grid-cols-1 p-4 md:grid-cols-4">
          {/* URL */}
          <div className="md:col-span-4">
            <div className="flex justify-between">
              <label
                htmlFor="url"
                className="block text-gray-700 text-sm font-medium">
                Long URL
              </label>
              {errors.url?.message != "" && (
                <span className="text-red-500 text-sm font-medium">
                  {errors.url?.message}
                </span>
              )}
            </div>
            <input
              ref={register}
              type="text"
              name="url"
              id="url"
              placeholder="https://en.wikipedia.org/wiki/Lorem_ipsum#cite_note-Cibois-1"
              className="block mt-1 w-full text-sm border-gray-300 focus:border-indigo-500 rounded-md shadow-sm focus:ring-indigo-500"
            />
          </div>
          {/* ID */}
          <div className="md:col-span-2">
            <div className="flex justify-between">
              <label
                htmlFor="id"
                className="block text-gray-700 text-sm font-medium">
                Custom Back-Half
              </label>
              {errors.id?.message != "" && (
                <span className="text-red-500 text-sm font-medium">
                  {errors.id?.message}
                </span>
              )}
            </div>
            <div className="flex mt-1 rounded-md shadow-sm">
              <span className="hidden items-center px-3 text-gray-500 text-sm bg-gray-50 border border-r-0 border-gray-300 rounded-l-md md:inline-flex">
                https://dhanz.xyz/go/
              </span>
              <input
                ref={register}
                type="text"
                name="id"
                id="id"
                placeholder="my-portofolio"
                className="block flex-1 w-full text-sm border-gray-300 focus:border-indigo-500 rounded-md focus:ring-indigo-500 md:rounded-none md:rounded-r-md"
              />
            </div>
          </div>
          {/* TTL */}
          <div className="md:col-span-2">
            <label
              htmlFor="ttl"
              className="block text-gray-700 text-sm font-medium">
              Time To Life
            </label>
            <select
              defaultValue={"year"}
              ref={register}
              id="ttl"
              name="ttl"
              className="block mt-1 px-3 py-2 w-full text-sm bg-white border border-gray-300 focus:border-indigo-500 rounded-md focus:outline-none shadow-sm focus:ring-indigo-500">
              <option value="hour">1 Hour</option>
              <option value="week">1 Week</option>
              <option value="month">1 Month</option>
              <option value="year">1 Year (Default)</option>
            </select>
          </div>
        </div>
        {/* SUBMIT BUTTON */}
        <div className="p-4 bg-gray-100">
          <button
            type="submit"
            className="relative px-4 py-2 w-full text-white text-sm font-medium tracking-wide bg-indigo-600 hover:bg-indigo-700 border border-transparent rounded-md focus:outline-none shadow-sm focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2">
            <svg
              className={`absolute right-2 top-2 w-5 h-5 text-white animate-spin ${
                isLoading ? "" : "hidden"
              }`}
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24">
              <circle
                className="opacity-25"
                cx="12"
                cy="12"
                r="10"
                stroke="currentColor"
                strokeWidth="4"></circle>
              <path
                className="opacity-75"
                fill="currentColor"
                d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            SHORTEN
          </button>
        </div>
      </form>
    </>
  )
}

export default ShortenForm
