'use client'

import { LoginForm } from '@/components/LoginForm'

export default function LoginPage() {
  const handleLogin = async (cpf: string, password: string) => {
    try {
      const response = await fetch('http://localhost:3333/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ cpf, password }),
      })

      if (!response.ok) {
        throw new Error('Falha no login')
      }

      const data = await response.json()
      console.log('Token:', data.token)
    } catch (err) {
      console.error(err)
    }
  }

  return (
    <div style={{ padding: '40px' }}>
      <LoginForm onSubmit={handleLogin} />
    </div>
  )
}
