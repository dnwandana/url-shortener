import { useRef, useState } from "react"
import dynamic from "next/dynamic"
import { useForm } from "react-hook-form"
import * as yup from "yup"
import { yupResolver } from "@hookform/resolvers/yup"
import axios from "axios"
import Link from "next/link"
import Metadata from "../components/Metadata"
import Navbar from "../components/Navbar"
import Form from "../components/Form"
import Field from "../components/Form/Field"
import Button from "../components/Button"

const FormAlert = dynamic(() => import("../components/Alert"))
const FormResult = dynamic(() => import("../components/Form/Result"))

type FormInput = {
  url: string
}

type ApiResponse = {
  statusCode: number
  url: {
    id: string
    title: string
    url: string
    createdAt: string
    updatedAt: string
  }
}

const formValidationSchema = yup.object().shape({
  url: yup
    .string()
    .url("Please, provide a valid url.")
    .required("Please, provide a valid url."),
})

const Homepage = () => {
  const [isSuccess, setIsSuccess] = useState<boolean>(false)
  const [longUrl, setLongUrl] = useState<string>("")
  const [shortUrl, setShortUrl] = useState<string>("")
  const { register, handleSubmit, errors } = useForm<FormInput>({
    resolver: yupResolver(formValidationSchema),
  })

  const onSubmit = async (url: FormInput) => {
    const domain = process.env.NEXT_PUBLIC_API_ENDPOINT
    const res = await axios.post(domain, url)
    const data: ApiResponse = res.data
    if (res.status === 201) {
      setIsSuccess(true)
      setShortUrl(`${domain}/${data.url.id}`)
      setLongUrl(data.url.url)
    }
  }

  const featureRef = useRef<null | HTMLDivElement>(null)
  const scrollToFeatureSection = () => featureRef.current.scrollIntoView()

  return (
    <>
      <Metadata title="URL Shortener" />
      <Navbar />
      {/* Hero */}
      <div className="bg-gray-100">
        <div className="bg-gray-100 container max-w-6xl mx-auto px-4 py-8 flex flex-col-reverse md:flex-row md:justify-between md:items-center">
          <div className="w-full md:w-3/5 mt-6 md:mt-0 text-left">
            <h1 className="text-indigo-700 font-bold text-2xl md:text-3xl lg:text-4xl">
              Shorten and Share
            </h1>
            <h2 className="mt-1 text-gray-800 font-medium text-xl md:text-2xl lg:text-3xl">
              Shorten a long URL and share as a powerful
            </h2>
            <div className="mt-6 md:mt-8 flex space-x-2 md:space-x-4">
              <Link href="/sign-up">
                <a className="flex justify-center items-center w-1/2 md:w-1/4 px-5 py-3 rounded-md font-medium text-base transition-colors duration-300 ease-in-out bg-indigo-700 text-indigo-100 hover:bg-indigo-600 hover:text-indigo-50 focus:outline-none focus:ring focus:ring-indigo-300">
                  Sign Up
                </a>
              </Link>
              <Button
                ButtonType="button"
                ButtonText="Learn More"
                ButtonClick={scrollToFeatureSection}
                ButtonClass="w-1/2 md:w-1/4 px-5 py-3 rounded-md font-medium text-base transition-colors duration-300 ease-in-out bg-indigo-200 text-indigo-700 hover:bg-indigo-300 hover:text-indigo-900 focus:outline-none focus:ring   focus:ring-indigo-300"
              />
            </div>
          </div>
          <img
            src="/assets/header.svg"
            alt="Social media engagement"
            className="w-1/2 md:w-2/5 mx-auto"
          />
        </div>
      </div>
      {/* URL Shortener Form */}
      <div className="bg-gray-900">
        <div className="bg-gray-900 container max-w-6xl mx-auto px-4 py-12">
          <Form
            FormClass="block md:flex relative"
            onSubmit={handleSubmit(onSubmit)}>
            <Field
              InputType="text"
              InputName="url"
              InputPlaceholder="Shorten your link here"
              InputClass="block md:flex w-full md:w-3/4 rounded-md font-medium text-base text-indigo-600 placeholder-indigo-500 border-indigo-300 shadow-sm focus:border-indigo-400 focus:ring focus:ring-indigo-400 focus:ring-opacity-50"
              Ref={register}
            />
            <Button
              ButtonType="submit"
              ButtonText="Shorten"
              ButtonClass="mt-3 ml-0 md:mt-0 md:ml-4 w-full md:w-1/4 px-2 py-3 rounded-md font-medium text-base tracking-wide transition-colors duration-300 ease-in-out bg-indigo-700 text-indigo-100 hover:bg-indigo-600 hover:text-indigo-50 focus:outline-none focus:ring focus:ring-indigo-300"
            />
          </Form>
          {/* <Alert Message /> */}
          {errors.url?.message && (
            <FormAlert
              AlertClass="block w-full py-4 rounded-md font-medium text-base text-center mt-4 bg-red-300 text-red-800"
              Message={errors.url?.message}
            />
          )}
          {/* <URL Result /> */}
          {isSuccess && <FormResult longURL={longUrl} shortURL={shortUrl} />}
        </div>
      </div>
      {/* Feature */}
      <div className="bg-gray-100" ref={featureRef}>
        <div className="bg-gray-100 container max-w-6xl mx-auto px-4 py-12">
          <h2 className="text-gray-800 text-center font-bold text-xl md:text-2xl lg:text-3xl">
            More Than URL Shortener
          </h2>
          <div className="mt-12">
            {/* Feature #1 */}
            <div className="flex sm:flex-row flex-col w-full md:w-4/5 mx-auto items-center border-b pb-12 mb-12 border-gray-300 space-x-0 space-y-8 md:space-x-8 md:space-y-0">
              <div className="space-y-2">
                <h2 className="text-gray-800 text-left font-medium text-lg md:text-xl lg:text-2xl">
                  Any link, any size
                </h2>
                <p className="text-gray-800 font-normal text-base md:text-lg lg:text-xl">
                  Any link, no matter what size, always shorten your links.
                </p>
              </div>
              <img
                src="/assets/any-link-any-size.gif"
                alt="Any link, any size"
                className="w-full md:w-1/2 rounded-xl shadow-md"
              />
            </div>
            {/* Feature #2 */}
            <div className="flex sm:flex-row flex-col-reverse w-full md:w-4/5 mx-auto items-center border-b pb-12 mb-12 border-gray-300 space-x-0 space-y-8 space-y-reverse md:space-x-8 md:space-y-0">
              <img
                src="/assets/recognizable-links.gif"
                alt="Recognizable links"
                className="w-full md:w-1/2 rounded-xl shadow-md"
              />
              <div className="space-y-2">
                <h2 className="text-gray-800 text-left font-medium text-lg md:text-xl lg:text-2xl">
                  Create recognizable links
                </h2>
                <p className="text-gray-800 font-normal text-base md:text-lg lg:text-xl">
                  Instead of a unique id, you can edit the back-halves of the
                  link into something meaningful.
                </p>
              </div>
            </div>
            <div className="flex flex-col items-center mb-2">
              <h2 className="text-gray-800 text-center font-bold text-xl md:text-2xl lg:text-3xl mb-4">
                Free URL Shortener
              </h2>
              <Link href="/sign-up">
                <a className="px-5 py-3 rounded-md font-medium text-base text-center transition-colors duration-300 ease-in-out bg-indigo-700 text-indigo-100 hover:bg-indigo-600 hover:text-indigo-50 focus:outline-none focus:ring focus:ring-indigo-300 tracking">
                  Sign Up Now
                </a>
              </Link>
            </div>
          </div>
        </div>
      </div>
      {/* Footer */}
      <div className="bg-gray-900">
        <div className="container max-w-6xl mx-auto py-6 px-4 bg-gray-900 inline-block md:flex md:justify-between md:items-center space-y-3 md:space-y-0">
          <p className="text-sm font-medium text-gray-300">
            SVG icons made by{" "}
            <a
              className="text-indigo-500 hover:text-indigo-400"
              href="https://heroicons.com/"
              title="Heroicons"
              target="_blank"
              rel="noreferrer">
              Heroicons
            </a>{" "}
            and{" "}
            <a
              className="text-indigo-500 hover:text-indigo-400"
              href="https://www.flaticon.com/authors/freepik"
              title="Freepik"
              target="_blank"
              rel="noreferrer">
              Freepik
            </a>
          </p>
          <p className="text-sm font-medium text-gray-300">
            By{" "}
            <a
              className="text-indigo-500 hover:text-indigo-400"
              href="https://github.com/dnwandana/url-shortener"
              title="Wandana"
              target="_blank"
              rel="noreferrer">
              Wandana
            </a>
          </p>
        </div>
      </div>
    </>
  )
}

export default Homepage
