import mongoose from "mongoose"
import { customAlphabet } from "nanoid"

const alphabet =
  "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const nanoid = customAlphabet(alphabet, 5)

const modelSchema = new mongoose.Schema(
  {
    _id: {
      type: String,
      default: nanoid
    },
    url: {
      type: String,
      required: true,
      max: 255
    },
    createdAt: {
      type: Date,
      default: Date.now
    }
  },
  {
    versionKey: false
  }
)

export default mongoose.model("Data", modelSchema)
