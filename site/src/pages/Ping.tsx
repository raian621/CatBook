import { useEffect, useState } from "react"
import { apiUrl } from "src/config"

export function Ping() {
  const [serverStatus, setServerStatus] = useState('Pending...')

  useEffect(() => {
    console.log('API_URL=', apiUrl)
    fetch(`${apiUrl}/ping`, {
      method: 'GET',
      mode: 'cors'
    }).then(response => {
      if (response.ok) {
        response.json().then(data => {
          console.log(data)
          setServerStatus('Online')
        }).catch((err) => {
          console.error(err)
          setServerStatus('Error')
        })
      }
    }).catch(err => {
      console.error(err)
      setServerStatus('Error')
    })
  }, [])

  return (
    <p><strong>Server status:</strong> {serverStatus}</p>
  )
}