import * as yup from "yup"
import { AxiosError, AxiosResponse } from "axios"
import { useRef, useState } from "react"
import { useForm } from "react-hook-form"
import { useRouter } from "next/router"
import { yupResolver } from "@hookform/resolvers/yup"
import Button from "../components/Button"
import Field from "../components/Form/Field"
import Form from "../components/Form"
import Link from "next/link"
import Metadata from "../components/Metadata"
import Navbar from "../layout/public/Navbar"
import axios from "../axiosInstance"
import dynamic from "next/dynamic"

const Alert = dynamic(() => import("../components/Alert"))

type FormInput = {
  fullname: string
  email: string
  password: string
  confirmationPassword: string
}

type ApiResponse = {
  statusCode: number
  user: {
    id: string
    fullname: string
  }
  error: string
}

const signUp = () => {
  const formValidationSchema = yup.object().shape({
    fullname: yup.string().required(),
    email: yup.string().email().required(),
    password: yup.string().min(8).required(),
    confirmationPassword: yup
      .string()
      .min(8, "password must be at least 8 characters")
      .oneOf([yup.ref("password")], "password does not match")
      .required("password does not match"),
  })

  const [isSuccess, setIsSuccess] = useState<boolean>(false)
  const [alertMessage, setAlertMessage] = useState<string>("")
  const alertRef = useRef<null | HTMLDivElement>(null)
  const scrollToAlert = () => alertRef.current.scrollIntoView()
  const router = useRouter()

  const { register, handleSubmit, errors } = useForm<FormInput>({
    resolver: yupResolver(formValidationSchema),
  })

  const submitForm = async (user: FormInput) => {
    setIsSuccess(true)
    try {
      await axios.post("/sign-up", user)
      setTimeout(() => {
        router.push("/thank-you")
      }, 1500)
    } catch (err) {
      const error = err as AxiosError
      const { data }: AxiosResponse<ApiResponse> = error.response
      setIsSuccess(false)
      setAlertMessage(`Error: ${data.error}`)
      scrollToAlert()
    }
  }

  return (
    <>
      <Metadata title="Sign Up – URL Shortener" />
      <Navbar />
      <div className="flex min-h-screen items-center justify-center bg-gray-100">
        <div className="p-4 w-full max-w-sm">
          {!isSuccess && alertMessage && (
            <div className="py-4" ref={alertRef}>
              <Alert
                Message={alertMessage}
                AlertClass="block w-full py-4 rounded-md font-medium text-sm md:text-base text-center bg-red-300 text-red-800"
              />
            </div>
          )}
          <div className="px-6 pt-8 pb-6 bg-white rounded-t-md shadow-md">
            <h1 className="text-center text-gray-800 font-bold text-xl md:text-2xl">
              Sign up and start shortening
            </h1>
            <Form
              FormClass="mt-8 grid grid-cols-1 gap-6"
              onSubmit={handleSubmit(submitForm)}>
              <div>
                <Field
                  LabelClass="block text-gray-700 text-sm font-medium"
                  Label="Fullname"
                  InputType="text"
                  InputId="fullname"
                  InputName="fullname"
                  InputPlaceholder="John Doe"
                  InputClass="mt-1 block w-full text-sm md:text-base rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-300 focus:ring-opacity-50"
                  Ref={register}
                  ErrorMessage={errors.fullname?.message}
                />
              </div>
              <div>
                <Field
                  LabelClass="block text-gray-700 text-sm font-medium"
                  Label="Email address"
                  InputType="email"
                  InputId="email"
                  InputName="email"
                  InputPlaceholder="john.doe@domain.name"
                  InputClass="mt-1 block w-full text-sm md:text-base rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-300 focus:ring-opacity-50"
                  Ref={register}
                  ErrorMessage={errors.email?.message}
                />
              </div>
              <div>
                <Field
                  LabelClass="block text-gray-700 text-sm font-medium"
                  Label="Password"
                  InputType="password"
                  InputId="password"
                  InputName="password"
                  InputPlaceholder="Password"
                  InputClass="mt-1 block w-full text-sm md:text-base rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-300 focus:ring-opacity-50"
                  Ref={register}
                  ErrorMessage={errors.password?.message}
                />
              </div>
              <div>
                <Field
                  LabelClass="block text-gray-700 text-sm font-medium"
                  Label="Confirmation password"
                  InputType="password"
                  InputId="confirmationPassword"
                  InputName="confirmationPassword"
                  InputPlaceholder="Confirmation password"
                  InputClass="mt-1 block w-full text-sm md:text-base rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-300 focus:ring-opacity-50"
                  Ref={register}
                  ErrorMessage={errors.confirmationPassword?.message}
                />
              </div>
              <Button
                ButtonType="submit"
                ButtonText="Sign Up"
                ButtonClass="w-full flex justify-center items-center px-4 py-2 rounded-md font-medium text-sm md:text-base tracking-wide transition-colors duration-300 ease-in-out bg-indigo-700 text-indigo-100 hover:bg-indigo-600 hover:text-indigo-50 focus:outline-none focus:ring focus:ring-indigo-300">
                <svg
                  className={`animate-spin mr-3 h-5 w-5 text-purple-300 ${
                    isSuccess ? "" : "hidden"
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
              </Button>
            </Form>
          </div>
          <div className="flex items-center justify-center py-4 text-center bg-gray-50 rounded-b-md shadow-md">
            <span className="text-sm text-gray-600">
              Already have an account?{" "}
            </span>{" "}
            <Link href="/sign-in">
              <a className="mx-2 text-sm font-bold text-indigo-600 hover:text-indigo-500">
                Sign in
              </a>
            </Link>
          </div>
        </div>
      </div>
    </>
  )
}

export default signUp
