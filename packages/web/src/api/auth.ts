import {Result} from "../type/result.ts";
import {req} from "./req.ts";

type loginForm = {
  email: string,
  password: string,
}
const login = (form: loginForm): Promise<Result<string>> => {
  return req.post("/user/login", form)
}

export {login}