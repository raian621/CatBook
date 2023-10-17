import { Link, Outlet } from "react-router-dom";
import './Layout.css'

export default function Layout() {
  return (
    <>
      <nav>
        <Link to="/">
          <h1>🐱 CatBook</h1>
        </Link>
        <div className="links">
          <Link to="/register">Register</Link>
          <Link to="/signin">Sign In</Link>
          <Link to="/cat/cat">Cat</Link>
        </div>
      </nav>
      <Outlet/>
    </>
  )
}