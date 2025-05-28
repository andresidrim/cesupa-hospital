'use client'

import { PacienteForm } from '@/components/PacienteForm'
import { PageContainer } from '@/components/PageContainer'
import { Header } from '@/components/Header'
import { useRouter } from 'next/navigation'

export default function NovoPacientePage() {
  const router = useRouter()

  const handlePacienteSubmit = async (data: Record<string, string>) => {
    try {
      const response = await fetch('http://localhost:3333/pacients', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data),
      })
      if (!response.ok) throw new Error('Erro ao cadastrar paciente')
      router.push('/pacients')
    } catch (err) {
      console.error(err)
    }
  }

  return (
    <PageContainer>
      <Header title="Cadastrar Paciente" actionLabel="Ver Pacientes" onActionClick={() => router.push('/pacientes')} />
      <PacienteForm onSubmit={handlePacienteSubmit} />
    </PageContainer>
  )
}
