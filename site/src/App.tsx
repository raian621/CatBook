import { BrowserRouter, Routes, Route } from 'react-router-dom'
import './App.css'
import Layout from './Layout'
import CatPage from './pages/CatPage'
import Home from './pages/Home'
import Register from './pages/Register'
import SignIn from './pages/SignIn'

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout/>}>
          <Route index element={<Home/>}/>
          <Route path="register" element={<Register/>}/>
          <Route path="signin" element={<SignIn/>}/>
          <Route path="cat/:catName" element={<CatPage/>}/>
        </Route>
      </Routes>
    </BrowserRouter>
  )
}

export default App
