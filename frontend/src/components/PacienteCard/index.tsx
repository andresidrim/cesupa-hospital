'use client'

import './style.css'
import { useRouter } from 'next/navigation'
import { useState } from 'react'

interface Paciente {
  id: string
  nome: string
  cpf: string
  sexo: string
  nascimento: string
  telefone: string
  status?: string
}

interface PacienteCardProps {
  paciente: Paciente
}

export const PacienteCard: React.FC<PacienteCardProps> = ({ paciente }) => {
  const router = useRouter()
  const [showActions, setShowActions] = useState(false)

  const togglePopover = () => setShowActions(!showActions)

  return (
    <div className="paciente-card">
      <div className="paciente-topo">
        <strong>{paciente.nome}</strong>
        <div className="actions-container">
          <span className={`status ${paciente.status?.toLowerCase()}`}>{paciente.status || 'Ativo'}</span>
          <button className="actions-button" onClick={togglePopover}>⋯</button>
          {showActions && (
            <div className="popover-menu">
              <button onClick={() => router.push(`/pacientes/${paciente.id}/editar`)}>✏️ Editar</button>
              <button onClick={() => router.push(`/pacientes/${paciente.id}/inativar`)}>🚫 Inativar</button>
              <button onClick={() => router.push(`/pacientes/${paciente.id}/agendar`)}>📅 Agendar</button>
            </div>
          )}
        </div>
      </div>

      <div className="paciente-info">
        <span><b>CPF:</b> {paciente.cpf}</span>
        <span><b>Nascimento:</b> {paciente.nascimento}</span>
        <span><b>Sexo:</b> {paciente.sexo}</span>
        <span><b>Telefone:</b> {paciente.telefone}</span>
      </div>
    </div>
  )
}
