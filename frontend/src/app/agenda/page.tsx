'use client'

import { PageContainer } from '@/components/PageContainer'
import { Header } from '@/components/Header'
import './style.css'

const mockConsultas = [
  {
    id: 'c1',
    paciente: 'Ana Clara',
    medico: 'Dra. Amanda Leal',
    data: '2025-06-10',
    horario: '14:30',
    motivo: 'Retorno',
  },
  {
    id: 'c2',
    paciente: 'Carlos Andrade',
    medico: 'Dr. Luiz Campos',
    data: '2025-06-11',
    horario: '10:00',
    motivo: 'Encaminhamento',
  },
  {
    id: 'c3',
    paciente: 'João Silva',
    medico: 'Dr. Rafael Borges',
    data: '2025-06-12',
    horario: '08:00',
    motivo: 'Avaliação inicial',
  },
]

export default function AgendaPage() {
  return (
    <PageContainer>
      <Header title="Agenda de Consultas" />

      <div className="agenda-container">
        {mockConsultas.map((consulta) => (
          <div key={consulta.id} className="agenda-card">
            <p><strong>Paciente:</strong> {consulta.paciente}</p>
            <p><strong>Médico:</strong> {consulta.medico}</p>
            <p><strong>Data:</strong> {consulta.data}</p>
            <p><strong>Horário:</strong> {consulta.horario}</p>
            <p><strong>Motivo:</strong> {consulta.motivo}</p>
          </div>
        ))}
      </div>
    </PageContainer>
  )
}
