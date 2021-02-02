import mongoose from "mongoose"

const db = () => {
  return mongoose.connect(
    process.env.DB_URI,
    {
      useNewUrlParser: true,
      useUnifiedTopology: true
    },
    () => console.log("Connected to database.")
  )
}

export default db
