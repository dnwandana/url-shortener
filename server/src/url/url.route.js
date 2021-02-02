import { Router } from "express"
import {
  shortenUrlController,
  getUrlController,
  redirectToHomepage
} from "./url.controller"

const router = Router()

router
  .route("/")
  .get(redirectToHomepage)
  .post(shortenUrlController)

router
  .route("/:id")
  .get(getUrlController)

export default router
