import { useParams } from "react-router"

export default function CatPage() {
  const { catName } = useParams()

  return (
    <>
      <h1>{catName} 🐱</h1>
    </>
  )
}