import { Route, Routes } from "react-router-dom"

import Login from "./pages/Login"
i

export default function App() {
  return (
    <>
      <TopBar />

      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/contato" element={<Contato />} />
        <Route path="/login" element={<Login />} />
        <Route path="*" element={<NotFound />} />
      </Routes>
    </>
  )
}
