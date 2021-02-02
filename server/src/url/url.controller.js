import { config as dotenv } from "dotenv"
import Data from "./url.model"
import validate from "./url.validation"

dotenv()
const { DOMAIN } = process.env

const shortenUrlController = async (req, res) => {
  try {
    const { error } = validate(req.body)

    if (error) {
      return res.status(400).json({
        error: {
          message: error.details[0].message
        }
      })
    }

    const { url } = req.body
    const data = new Data({
      url
    })

    await data.save()
    return res.status(201).json({
      data: {
        longUrl: url,
        shortUrl: `${DOMAIN}/go/${data._id}`
      }
    })
  } catch (error) {
    console.error(error)
    return res.status(400).json({
      error
    })
  }
}

const getUrlController = async (req, res) => {
  try {
    const { id } = req.params
    const data = await Data.findById(id)

    if (!data) {
      return res.redirect(`${DOMAIN}/404`)
    }

    return res.redirect(data.url)
  } catch (error) {
    console.error(error)
    return res.status(400).json({
      error
    })
  }
}

const redirectToHomepage = (_, res) => {
  return res.redirect(DOMAIN)
}

export { shortenUrlController, getUrlController, redirectToHomepage }
