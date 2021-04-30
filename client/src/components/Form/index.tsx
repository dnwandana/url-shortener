import { FormEventHandler, ReactNode } from "react"

type FormProps = {
  onSubmit?: FormEventHandler<HTMLFormElement>
  FormClass?: string
  children?: ReactNode
}

const Form = ({ onSubmit, FormClass, children }: FormProps) => {
  return (
    <>
      <form onSubmit={onSubmit} className={FormClass}>
        {children}
      </form>
    </>
  )
}

export default Form
