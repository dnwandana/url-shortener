import { MouseEventHandler } from "react"

type ButtonProps = {
  ButtonType: "button" | "submit"
  ButtonClick?: MouseEventHandler<HTMLButtonElement>
  ButtonClass: string
  ButtonText: string
}

const Button = ({
  ButtonType,
  ButtonClick,
  ButtonClass,
  ButtonText,
}: ButtonProps): JSX.Element => {
  return (
    <>
      <button type={ButtonType} onClick={ButtonClick} className={ButtonClass}>
        {ButtonText}
      </button>
    </>
  )
}

export default Button
