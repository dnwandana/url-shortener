import Joi from "joi"

const urlValidation = (data) => {
  const pattern = /((https?):\/\/)?(www.)?[a-z0-9-]+(\.[a-z]{2,}){1,3}(#?\/?[a-zA-Z0-9#-]+)*\/?(\?[a-zA-Z0-9-_]+=[a-zA-Z0-9-%]+&?)?$/
  const urlSchema = Joi.object({
    url: Joi.string().regex(pattern).required().messages({
      "string.pattern.base": "Please, provide a valid url.",
      "string.empty": "Please, provide a valid url.",
      "any.required": 'Missing "url" value.'
    })
  })

  return urlSchema.validate(data)
}

export default urlValidation
