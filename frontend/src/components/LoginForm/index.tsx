'use client'

import React, { useState } from 'react'
import './style.css'
import { Input } from '@/components/Input'
import { Button } from '@/components/Button'

interface LoginFormProps {
  onSubmit: (cpf: string, password: string) => void
}

export const LoginForm: React.FC<LoginFormProps> = ({ onSubmit }) => {
  const [cpf, setCpf] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState<string | null>(null)

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    if (!cpf || !password) {
      setError('Preencha todos os campos')
      return
    }
    setError(null)
    onSubmit(cpf, password)
  }

  return (
    <form className="login-form" onSubmit={handleSubmit}>
      <h2>Login</h2>
      <Input
        label="CPF"
        name="cpf"
        value={cpf}
        onChange={(e) => setCpf(e.target.value)}
      />
      <Input
        label="Senha"
        name="password"
        type="password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />
      {error && <span className="login-error">{error}</span>}
      <Button type="submit">Entrar</Button>
    </form>
  )
}
