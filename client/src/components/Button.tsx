import { MouseEventHandler, ReactNode } from "react"

type ButtonProps = {
  ButtonType: "button" | "submit"
  ButtonClick?: MouseEventHandler<HTMLButtonElement>
  ButtonClass: string
  ButtonText: string
  children?: ReactNode
}

const Button = ({
  ButtonType,
  ButtonClick,
  ButtonClass,
  ButtonText,
  children,
}: ButtonProps): JSX.Element => {
  return (
    <>
      <button type={ButtonType} onClick={ButtonClick} className={ButtonClass}>
        {children} {ButtonText}
      </button>
    </>
  )
}

export default Button
