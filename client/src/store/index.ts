import create from "zustand"

const store = create((set) => ({
  success: true,
  setSuccess: (status: boolean) => {
    set(() => ({ success: status }))
  },
  alertMessage: "",
  setAlertMessage: (message: string) => {
    set(() => ({ alertMessage: message }))
  },
  shortURL: "",
  setShortURL: (url: string) => {
    set(() => ({ shortURL: url }))
  },
  longURL: "",
  setLongURL: (url: string) => {
    set(() => ({ longURL: url }))
  },
}))

export default store
