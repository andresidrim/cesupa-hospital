'use client'

import { useState } from 'react'
import './style.css'
import { Input } from '@/components/Input'
import { Button } from '@/components/Button'
import { PacienteCard } from '@/components/PacienteCard'

interface Paciente {
  id: string
  nome: string
  cpf: string
  nascimento: string
  telefone: string
  sexo: string
  status?: string
}

const MOCK_PACIENTES: Paciente[] = [
  {
    id: '1',
    nome: 'João Silva',
    cpf: '123.456.789-00',
    nascimento: '1990-01-01',
    telefone: '(91) 99999-1234',
    sexo: 'Masculino',
    status: 'Ativo',
  },
  {
    id: '2',
    nome: 'Maria Oliveira',
    cpf: '987.654.321-00',
    nascimento: '1985-06-15',
    telefone: '(91) 99999-5678',
    sexo: 'Feminino',
    status: 'Inativo',
  },
]

export const PacienteConsulta = () => {
  const [query, setQuery] = useState('')
  const [resultados, setResultados] = useState<Paciente[]>([])
  const [erro, setErro] = useState<string | null>(null)

  const buscar = () => {
    if (!query.trim()) {
      setErro('Digite um nome ou CPF')
      return
    }

    const filtrados = MOCK_PACIENTES.filter((p) =>
      p.nome.toLowerCase().includes(query.toLowerCase()) ||
      p.cpf.includes(query) ||
      p.id === query
    )

    if (filtrados.length === 0) {
      setErro('Nenhum paciente encontrado.')
      setResultados([])
    } else {
      setErro(null)
      setResultados(filtrados)
    }
  }

  return (
    <div className="consulta-container">
      <div className="consulta-barra">
        <Input
          name="busca"
          placeholder="Digite nome, CPF ou ID"
          value={query}
          onChange={(e) => setQuery(e.target.value)}
        />
        <Button onClick={buscar}>Buscar</Button>
      </div>

      {erro && <span className="consulta-erro">{erro}</span>}

      <div className="consulta-lista">
        {resultados.map((p) => (
          <PacienteCard key={p.id} paciente={p} />
        ))}
      </div>
    </div>
  )
}
