import { useEffect, useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

function App() {
  const [mensaje, setMensaje] = useState("")

  useEffect(() => {
    fetch("http://localhost:3000/api")
    .then((res) => res.json())
    .then((data) => setMensaje(data.message))
  }, [])

  return (
    <>
      <div className="flex flex-col items-center justify-center min-h-screen bg-slate-100">
        <h1 className="text-3xl font-bold text-blue-600">{ mensaje }</h1>
      </div>
    </>
  )
}

export default App
