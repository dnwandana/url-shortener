import { config as dotenv } from "dotenv"
import express from "express"
import cors from "cors"
import helmet from "helmet"
import database from "./config/db"
import apiRoute from "./url/url.route"

const app = express()
const PORT = 5000 || process.env.PORT

dotenv()
app.use(cors())
app.use(helmet())
app.use(express.json())
app.use("/go", apiRoute)

const application = () => {
  try {
    database()
    app.listen(PORT, () =>
      console.log(`Server up and running on port: ${PORT}.`)
    )
  } catch (error) {
    console.log(error)
  }
}

export default application
