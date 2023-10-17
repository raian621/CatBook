import { useRef } from "react"

export default function Register() {
  const username = useRef<string>("")
  const password = useRef<string>("")

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()

    console.log({
      username: username.current,
      password: password.current
    })
  }

  return (
    <>
      <h1>Register</h1>
      <form onSubmit={handleSubmit}>
        <label htmlFor="username">Username</label>
        <input name="username"  onChange={
          (e: React.ChangeEvent<HTMLInputElement>) => {
            username.current = e.currentTarget.value
          }
        }/>
        <label htmlFor="password">Password</label>
        <input name="password" type="password" onChange={
          (e: React.ChangeEvent<HTMLInputElement>) => {
            password.current = e.currentTarget.value
          }
        }/>
        <input type="submit" value="Register"/>
      </form>
    </>
  )
}