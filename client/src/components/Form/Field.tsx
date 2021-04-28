import { LegacyRef } from "react"

type FieldProps = {
  LabelClass?: string
  Label?: string
  InputType: "email" | "password" | "text"
  InputId?: string
  InputName: string
  InputPlaceholder: string
  InputClass: string
  Ref?: LegacyRef<HTMLInputElement>
  ErrorMessage?: string
}

const Field = ({
  LabelClass,
  Label,
  InputType,
  InputId,
  InputName,
  InputPlaceholder,
  InputClass,
  Ref,
  ErrorMessage,
}: FieldProps): JSX.Element => {
  return (
    <>
      {LabelClass === "" && (
        <label htmlFor={InputId} className={LabelClass}>
          {Label}
        </label>
      )}
      <input
        type={InputType}
        id={InputId}
        name={InputName}
        placeholder={InputPlaceholder}
        className={InputClass}
        ref={Ref}
      />
      {ErrorMessage !== "undefined" && (
        <div className="mt-1">
          <span className="text-red-600">{ErrorMessage}</span>
        </div>
      )}
    </>
  )
}

export default Field
