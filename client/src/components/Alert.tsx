type AlertProps = {
  AlertClass: string
  Message: string
}

const Alert = ({ AlertClass, Message }: AlertProps): JSX.Element => {
  return (
    <div className={AlertClass}>
      {Message}
    </div>
  )
}

export default Alert
